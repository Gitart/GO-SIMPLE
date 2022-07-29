package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// define client secrets
	// Upvest Client ID
	clientID := "d5193a20-b1a1-4984-a678-8742b704a177"
	// Signing key
	keyID := "b2a53e9f-8e1b-426c-89f3-fe78818359b8"
	_, privateKey, _ := ed25519.GenerateKey(rand.Reader)

	// create a generic HTTP request
	requestBody := []byte(`{"key": "value"}`)
	request, _ := http.NewRequest(http.MethodPost, "https://server", bytes.NewReader(requestBody))

	// set mandatory request headers
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-length", strconv.Itoa(len(requestBody)))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("date", time.Now().Format(time.RFC1123))
	request.Header.Set("upvest-client-id", clientID)

	// set request digest
	requestBodyHash := sha256.Sum256(requestBody)
	digestValue := "SHA-256=" + base64.StdEncoding.EncodeToString(requestBodyHash[:])
	request.Header.Set("digest", digestValue)

	// set signature input
	signatureInput, err := prepareSignatureInput(request, keyID)
	if err != nil {
		log.Println("failed to prepare signature input:", err)
		os.Exit(1)
	}
	request.Header.Add("signature-input", "sig1="+signatureInput.signatureInput())

	// set signature
	signaturePayload := signatureInput.signaturePayload()
	signatureValue := ed25519.Sign(privateKey, signaturePayload)
	signature := fmt.Sprintf("sig1=:%s:", base64.StdEncoding.EncodeToString(signatureValue))
	request.Header.Add("signature", signature)

	// dump signed request
	rawRequest, _ := httputil.DumpRequest(request, true)
	log.Println(string(rawRequest))
}

// signatureComponent define simple key-value pair for the signature components
type signatureComponent struct {
	key   string
	value interface{}
}

// signatureInput aggregation of all signature components
// required for signature-input header creation and filling of payload message for signature calculation
type signatureInput struct {
	components []signatureComponent
	meta       []signatureComponent
}

// prepareSignatureInput fulfills signatureInput using values from request header
func prepareSignatureInput(r *http.Request, keyID string) (*signatureInput, error) {
	signatureInput := signatureInput{}

	// add common HTTP properties: method, path and optional query
	signatureInput.addComponent("@method", r.Method)
	signatureInput.addComponent("@path", r.URL.Path)
	if r.URL.RawQuery != "" {
		signatureInput.addComponent("@query", "?"+r.URL.RawQuery)
	}

	// include supported request headers
	signatureInput.addComponent("accept", r.Header.Get("accept"))
	signatureInput.addComponent("authorization", r.Header.Get("authorization"))
	signatureInput.addComponent("content-length", r.Header.Get("content-length"))
	signatureInput.addComponent("content-type", r.Header.Get("content-type"))
	signatureInput.addComponent("digest", r.Header.Get("digest"))
	signatureInput.addComponent("idempotency-key", r.Header.Get("idempotency-key"))
	signatureInput.addComponent("upvest-client-id", r.Header.Get("upvest-client-id"))

	// add metadata fields
	signatureInput.addMeta("keyid", keyID)

	if r.Header.Get("date") != "" {
		created, err := time.Parse(time.RFC1123, r.Header.Get("date"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse 'date' header: %v", err)
		}
		signatureInput.addMeta("created", created.Unix())
	}
	if r.Header.Get("expires") != "" {
		expires, err := time.Parse(time.RFC1123, r.Header.Get("expires"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse 'expires' header: %v", err)
		}
		signatureInput.addMeta("expires", expires.Unix())
	}

	nonce, err := generateNonce()
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce value: %v", err)
	}
	signatureInput.addMeta("nonce", nonce)

	return &signatureInput, nil
}

// addComponent adds signature-input component
func (s *signatureInput) addComponent(key string, value string) {
	if value == "" {
		return
	}
	s.components = append(s.components, signatureComponent{key, value})
}

// addComponent adds signature-input meta information
func (s *signatureInput) addMeta(key string, value interface{}) {
	if value == "" {
		return
	}
	s.meta = append(s.meta, signatureComponent{key, value})
}

// signatureInput returns signature metadata for signature-input header
func (s *signatureInput) signatureInput() string {
	var signatureParams strings.Builder
	signatureParams.WriteByte('(')
	for i, component := range s.components {
		signatureParams.WriteByte('"')
		signatureParams.WriteString(component.key)
		signatureParams.WriteByte('"')
		if i != len(s.components)-1 {
			signatureParams.WriteByte(' ')
		}
	}
	signatureParams.WriteByte(')')
	for _, component := range s.meta {
		signatureParams.WriteByte(';')
		signatureParams.WriteString(component.key)
		signatureParams.WriteByte('=')
		switch value := component.value.(type) {
		case string:
			signatureParams.WriteByte('"')
			signatureParams.WriteString(value)
			signatureParams.WriteByte('"')
		case []byte:
			signatureParams.WriteByte('"')
			signatureParams.Write(value)
			signatureParams.WriteByte('"')
		case int64:
			signatureParams.WriteString(strconv.FormatInt(value, 10))
		}
	}
	return signatureParams.String()
}

// signaturePayload returns signature payload
func (s *signatureInput) signaturePayload() []byte {
	var signaturePayload bytes.Buffer
	for _, component := range s.components {
		signaturePayload.WriteString(component.key)
		signaturePayload.WriteString(": ")
		signaturePayload.WriteString(component.value.(string))
		signaturePayload.WriteByte('\n')
	}
	signaturePayload.WriteString("@signature-params: ")
	signaturePayload.WriteString(s.signatureInput())
	return signaturePayload.Bytes()
}

// generateNonce returns random nonce value that consists of characters from defined alphabet
func generateNonce() ([]byte, error) {
	var nonceAlphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	const nonceLength = 16

	runes := make([]byte, nonceLength)
	for i := 0; i < len(runes); i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(nonceAlphabet))))
		if err != nil {
			return nil, err
		}
		runes[i] = nonceAlphabet[r.Int64()]
	}
	return runes, nil
}
