## Activate web camera and broadcast out base64 encoded images



For this tutorial, we will learn how to activate web camera with Golang + OpenCV and stream out base64 encoded images via 
HTTP server. This tutorial is derived from the lesson that I use to teach my son on the subject of computer vision and Golang.
It is based on the Go-OpenCV tutorial on activating web camera (https://github.com/lazywei/go-opencv/blob/master/samples/webcam.go).

The program will activate the web camera when you point your web browser to localhost:8080 and begin streaming out base64 encoded images. 
You will have to scroll down the see the new images. LOL! Good enough to get a 3 years old kid excited about computer vision.  


Before you start.
```
go get github.com/lazywei/go-opencv
```
Here you go!


```golang
 package main

 import (
 	"bytes"
 	"encoding/base64"
 	"fmt"
 	"github.com/lazywei/go-opencv/opencv"
 	"image/png"
 	"net/http"
 )

 func broadcast(w http.ResponseWriter, r *http.Request) {
 	webCamera := opencv.NewCameraCapture(0)

 	if webCamera == nil {
 		panic("Unable to open camera")
 	}

 	defer webCamera.Release()

 	for {
 		if webCamera.GrabFrame() {
 			imgFrame := webCamera.RetrieveFrame(1)
 			if imgFrame != nil {
 				//fmt.Println(imgFrame.ImageSize())
 				//fmt.Println(imgFrame.ToImage())

 				// convert IplImage(Intel Image Processing Library)
 				// to image.Image
 				goImgFrame := imgFrame.ToImage()

 				// and then convert to []byte
 				// with the help of png.Encode() function

 				frameBuffer := new(bytes.Buffer)
 				//frameBuffer := make([]byte, imgFrame.ImageSize())
 				err := png.Encode(frameBuffer, goImgFrame)

 				if err != nil {
 					panic(err)
 				}

 				// convert the buffer bytes to base64 string - use buf.Bytes() for new image
 				imgBase64Str := base64.StdEncoding.EncodeToString(frameBuffer.Bytes())

 				// Embed into an html without PNG file
 				img2html := "<html><body><img src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"

 				w.Write([]byte(fmt.Sprintf(img2html)))

 				// TODO :
 				// encode frames to stream via WebRTC

 				fmt.Println("Streaming....")

 			}
 		}
 	}

 }

 func main() {
 	fmt.Println("Broadcasting...")
 	mux := http.NewServeMux()
 	mux.HandleFunc("/", broadcast)
 	http.ListenAndServe(":8080", mux)
 }
 ```
 
streaming out base64 encoded images to browser
