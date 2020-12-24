# Logging Go Programs

The standard library package log provides a basic infrastructure for log management in GO language that can be used for logging our GO programs. The main purpose of logging is to get a trace of what's happening in the program, where it's happening, and when it's happening. Logs can be providing code tracing, profiling, and analytics. Logging(eyes and ears of a programmer) is a way to find those bugs and learn more about how the program is functioning.

#### To work with package log, we must add it to the list of imports:

import (
	"log"
)

In its simplest usage, it formats messages and sends them to Standard Error check below example:

// Program in GO language to demonstrates how to use base log package.
package main
import (
	"log"
)
func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
// Println writes to the standard logger.
	log.Println("main started")

// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}

After executing this code, the output would look something like this:

C:\\golang>go run example38.go LOG: 2017/06/25 14:49:41.989813 C:/golang/example38.go:11: init started LOG: 2017/06/25 14:49:41.990813 C:/golang/example38.go:15: main started LOG: 2017/06/25 14:49:41.990813 C:/golang/example38.go:18: fatal message exit status 1

Sending messages to Standard Error is useful for simple tools. When we're building servers, applications, or system services, we need a better place to send your log messages. Here all error messages are all sent to Standard Error, regardless of whether the message is an actual error or an informational message.

The standard log entry contains below things:
\- a prefix (log.SetPrefix("LOG: "))
\- a datetime stamp (log.Ldate)
\- full path to the source code file writing to the log (log.Llongfile)
\- the line of code performing the write and finally the message.

This pieces of information are automatically generated for us, information about when the event happened and information about where it happened.

Println is the standard way to write log messages.
Fatalln or any of the other "fatal" calls, the library prints the error message and then calls os.Exit(1), forcing the program to quit.
Panicln is used to write a log message and then issue a panic which may unless recovered or will cause the program to terminate.

## Program in GO language with real world example of logging.

Now i am taking a real world example and implementing above log package in my program. For example i am testing an SMTP connection is working fine or not. For test case I am going to connect to a SMTP server "smtp.smail.com" which is not exist, hence program will terminate with a log message.

// Program in GO language with real world example of logging.
package main

import (
"log"
"net/smtp"
)
func init(){
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
// Connect to the remote SMTP server.
client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
client.Data()
}

C:\\golang>go run example39.go TRACE: 2017/06/25 14:54:42.662011 C:/golang/example39.go:9: init started TRACE: 2017/06/25 14:55:03.685213 C:/golang/example39.go:15: dial tcp 23.27.98.252:25: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond. exit status 1

The above program is throwing fatal exception from log.Fatalln(err).
