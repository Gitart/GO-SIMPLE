package magicimage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var SetValue http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := New(r, 32<<20)

	switch r.Header.Get("set-func") {
	case "allowext":
		magic.SetAllowExt([]string{"gif", "png"})
		fmt.Fprint(rw, strings.Join(magic.AllowExt, ","))
		return
	case "maxfilesize":
		magic.SetMaxFileSize(1)
		fmt.Fprint(rw, magic.MaxFileSize)
		return
	case "required":
		magic.SetRequired(false)
		fmt.Fprint(rw, magic.Required)
		return
	case "minfileinslize":
		magic.SetMinFileInSlice(2)
		fmt.Fprint(rw, magic.MinFileInSlice)
		return
	case "maxfileinslice":
		magic.SetMaxFileInSlice(10)
		fmt.Fprint(rw, magic.MaxFileInSlice)
		return
	}
}

var SingleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := New(r, 32<<20)

	switch r.Header.Get("set-func") {
	case "required":
		magic.SetRequired(false)
	}

	if err := magic.ValidateSingleImage("file"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

	fmt.Fprint(rw, "success")
}

var MultipleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := New(r, 32<<20)
	magic.SetMinFileInSlice(2)
	magic.SetMaxFileInSlice(3)

	switch r.Header.Get("set-func") {
	case "required":
		magic.SetRequired(false)
	}

	if err := magic.ValidateMultipleImage("files"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

	fmt.Fprint(rw, "success")
}

var SaveImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := New(r, 32<<20)
	if err := magic.ValidateMultipleImage("files"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

	magic.SaveImages(200, 200, "out/this-is-slug", true)

	fmt.Fprint(rw, "success")
}

func TestSetValue(t *testing.T) {
	tests := [...]struct {
		name     string
		header   string
		expected string
	}{
		{
			name:     "SetAllowExt",
			header:   "allowext",
			expected: "gif,png",
		},
		{
			name:     "SetMaxFileSize",
			header:   "maxfilesize",
			expected: "1",
		},
		{
			name:     "SetRequired",
			header:   "required",
			expected: "false",
		},
		{
			name:     "SetMinFileInSlice",
			header:   "minfileinslize",
			expected: "2",
		},
		{
			name:     "SetMaxFileInSlice",
			header:   "maxfileinslice",
			expected: "10",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request.Header.Add("Set-Func", test.header)
			recorder := httptest.NewRecorder()

			SetValue(recorder, request)

			response := recorder.Result()
			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, test.expected, string(body))
		})
	}
}

func TestValidationSingleImage(t *testing.T) {
	tests := [...]struct {
		name     string
		expected string
		image    string
	}{
		{
			name:     "required false",
			expected: "success",
		},
		{
			name:     "required true",
			expected: "Image is required.",
		},
		{
			name:     "invalid extension",
			expected: "Image must be between jpeg, png.",
			image:    "testdata/test.txt",
		},
		{
			name:     "invalid size file",
			expected: "An image cannot greater than 4 Mb.",
			image:    "testdata/size.jpeg",
		},
		{
			name:     "success upload",
			expected: "success",
			image:    "testdata/valid1.jpg",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			writer := multipart.NewWriter(b)

			switch test.name {
			case "invalid extension", "invalid size file", "success upload":
				file, _ := writer.CreateFormFile("file", strings.Split(test.image, "/")[1])
				testdata, _ := ioutil.ReadFile(test.image)
				file.Write(testdata)
			}
			writer.Close()

			request := httptest.NewRequest(http.MethodPost, "/", b)
			if test.name == "required false" {
				request.Header.Add("Set-Func", "required")
			}
			request.Header.Set("Content-Type", writer.FormDataContentType())
			recorder := httptest.NewRecorder()

			SingleImage(recorder, request)

			response := recorder.Result()
			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, test.expected, string(body))
		})
	}
}

func TestValidationMultipleImage(t *testing.T) {
	tests := [...]struct {
		name     string
		expected string
		images   []string
	}{
		{
			name:     "required false",
			expected: "success",
		},
		{
			name:     "required true",
			expected: "Image is required.",
		},
		{
			name:     "invalid extension",
			expected: "The image at index 2 must be between jpeg, png.",
			images:   []string{"testdata/valid1.jpg", "testdata/test.gif"},
		},
		{
			name:     "invalid size",
			expected: "An image at index 2 cannot greater than 4 Mb.",
			images:   []string{"testdata/valid1.jpg", "testdata/size.jpeg"},
		},
		{
			name:     "image must be unique",
			expected: "Each image must be unique.",
			images:   []string{"testdata/valid1.jpg", "testdata/valid1.jpg"},
		},
		{
			name:     "min image in slice",
			expected: "At least 2 image must be upload.",
			images:   []string{"testdata/valid1.jpg"},
		},
		{
			name:     "max image in slice",
			expected: "Maximal 3 images to be upload.",
			images:   []string{"testdata/valid1.jpg", "testdata/valid2.jpg", "testdata/valid3.png", "testdata/valid4.png"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			writer := multipart.NewWriter(b)

			switch test.name {
			case "invalid extension", "invalid size", "image must be unique", "min image in slice", "max image in slice":
				for _, val := range test.images {
					file, _ := writer.CreateFormFile("files", strings.Split(val, "/")[1])
					testdata, _ := ioutil.ReadFile(val)
					file.Write(testdata)
				}
			}
			writer.Close()

			request := httptest.NewRequest(http.MethodPost, "/", b)
			if test.name == "required false" {
				request.Header.Add("Set-Func", "required")
			}
			request.Header.Set("Content-Type", writer.FormDataContentType())
			recorder := httptest.NewRecorder()

			MultipleImage(recorder, request)

			response := recorder.Result()
			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, test.expected, string(body))
		})
	}
}

func TestSaveImage(t *testing.T) {
	b := new(bytes.Buffer)
	writer := multipart.NewWriter(b)

	images := []string{
		"testdata/valid1.jpg",
		"testdata/valid2.jpg",
		"testdata/valid3.png",
		"testdata/valid4.png",
		"testdata/valid5.png",
		"testdata/iphone.jpeg",
	}
	for _, val := range images {
		file, _ := writer.CreateFormFile("files", strings.Split(val, "/")[1])
		testdata, _ := ioutil.ReadFile(val)
		file.Write(testdata)
	}
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/", b)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	SaveImage(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, "success", string(body))
	// check len files
	time.Sleep(1 * time.Second)
	files, _ := ioutil.ReadDir("out/this-is-slug")
	assert.Equal(t, 6, len(files))
	os.RemoveAll("out")
}
