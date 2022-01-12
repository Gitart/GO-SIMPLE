// 
// https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
// 
package main

import (
 "io"
  "io/ioutil"
  "log"
  "os"
)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func Init(traceHandle,infoHandle,warningHandle,errorHandle io.Writer) {
    Trace   = log.New(traceHandle,   "TRACE: ",   log.Ldate|log.Ltime|log.Lshortfile)
    Info    = log.New(infoHandle,    "INFO: ",    log.Ldate|log.Ltime|log.Lshortfile)
    Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error   = log.New(errorHandle,   "ERROR: ",   log.Ldate|log.Ltime|log.Lshortfile)
}

func Logsd() {
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    Trace.Println("I have something standard to say")
    Info.Println("Special Information")
    Warning.Println("There is something you need to know about")
    Error.Println("Something has failed")
}
