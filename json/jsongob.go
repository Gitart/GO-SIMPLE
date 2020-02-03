package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func createMap(max int) map[int64]float64 {
	m := make(map[int64]float64)
	for i := 0; i < max; i++ {
		m[int64(i)] = float64(i)
	}
	return m
}

func createSliceMap(max int) []map[int64]float64 {
	list := make([]map[int64]float64, max)
	for i := 0; i < max; i++ {
		list[i] = createMap(max)
	}
	return list
}
func encodeGob(v interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func decodeGob(b []byte, result interface{}) {
	buf := bytes.NewBuffer(b)
	enc := gob.NewDecoder(buf)

	err := enc.Decode(result)
	if err != nil {
		panic(err)
	}
}

func encodeJSON(v interface{}) []byte {
	byt, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return byt
}

func decodeJSON(b []byte, result interface{}) {
	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}
}
