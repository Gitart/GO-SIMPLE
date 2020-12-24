# Files and directories with examples

The most important package that allows us to manipulate files and directories as entities is the `os` package.
The `io` package has the `io.Reader` interface to reads and transfers data from a source into a stream of bytes. The `io.Writer` interface reads data from a provided stream of bytes and writes it as output to a target resource.

---

## Create an empty file

package main

import (
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("empty.txt")
	if err !\= nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

C:\\golang\\working\-with\-files>go fmt example1.go C:\\golang\\working\-with\-files>golint example1.go C:\\golang\\working\-with\-files>go run example1.go 2018/08/11 15:46:04 &{0xc042060780} C:\\golang\\working\-with\-files>

---

## Go program to Create directory or folder if not exist

package main

import (
	"log"
	"os"
)

func main() {
	\_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir !\= nil {
			log.Fatal(err)
		}

	}
}

---

## Rename a file in Golang

package main

import (
	"log"
	"os"
)

func main() {
	oldName := "test.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err !\= nil {
		log.Fatal(err)
	}
}

C:\\golang\\working\-with\-files>go fmt example.go C:\\golang\\working\-with\-files>golint example.go C:\\golang\\working\-with\-files>go run example.go C:\\golang\\working\-with\-files>

---

## Move a file from one location to another in Golang

os.Rename() can also move file from one location to another at same time renaming file name.

package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err !\= nil {
		log.Fatal(err)
	}
}

---

## Golang Create Copy of a file at another location

package main

import (
	"io"
	"log"
	"os"
)

func main() {

	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err !\= nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err !\= nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err !\= nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

\[root@server src\]# clear \[root@server src\]# go run example6.go 2018/08/15 03:43:39 Copied 100 bytes. \[root@server src\]#

---

## Get file information in Golang

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileStat, err := os.Stat("test.txt")

	if err !\= nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}

C:\\golang\\working\-with\-files>go fmt example.go example.go C:\\golang\\working\-with\-files>golint example.go C:\\golang\\working\-with\-files>go run example.go File Name: test.txt Size: 100 Permissions: \-rw\-rw\-rw\- Last Modified: 2018\-08\-11 20:19:14.2671925 +0530 IST Is Directory: false C:\\golang\\working\-with\-files>

---

## Golang program to delete a specific file

package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("/var/www/html/test.txt")
	if err !\= nil {
		log.Fatal(err)
	}
}

---

## Go program to read a text file character by character

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err !\= nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
}

---

## Reduce file size

os.Truncate() function will reduce the file content upto N bytes passed in second parameter. In below example if size of test.txt file is more that 1Kb(100 byte) then it will truncate the remaining content.

package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("test.txt", 100)

	if err !\= nil {
		log.Fatal(err)
	}
}

C:\\golang\\working\-with\-files>go fmt example10.go C:\\golang\\working\-with\-files>golint example10.go C:\\golang\\working\-with\-files>go run example10.go C:\\golang\\working\-with\-files>

---

## Go program to add or append content at the end of text file

package main

import (
	"fmt"
	"os"
)

func main() {
	message := "Add this content at end"
	filename := "test.txt"

	f, err := os.OpenFile(filename, os.O\_RDWR|os.O\_APPEND|os.O\_CREATE, 0660)

	if err !\= nil {
		fmt.Println(err)
		os.Exit(\-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\\n", message)
}

---

## Golang Changing permissions, ownership, and timestamps

package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// Test File existence.
	\_, err := os.Stat("test.txt")
	if err !\= nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// Change permissions Linux.
	err = os.Chmod("test.txt", 0777)
	if err !\= nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err !\= nil {
		log.Println(err)
	}

	// Change file timestamps.
	addOneDayFromNow := time.Now().Add(24 \* time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err !\= nil {
		log.Println(err)
	}
}

---

## Go program to compress list of files into Zip file

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func appendFiles(filename string, zipw \*zip.Writer) error {
	file, err := os.Open(filename)
	if err !\= nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err !\= nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if \_, err := io.Copy(wr, file); err !\= nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func main() {
	flags := os.O\_WRONLY | os.O\_CREATE | os.O\_TRUNC
	file, err := os.OpenFile("test.zip", flags, 0644)
	if err !\= nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	var files = \[\]string{"test1.txt", "test2.txt", "test3.txt"}

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for \_, filename := range files {
		if err := appendFiles(filename, zipw); err !\= nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
}

---

## Go program to read list of files inside Zip file

package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func listFiles(file \*zip.File) error {
	fileread, err := file.Open()
	if err !\= nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fmt.Fprintf(os.Stdout, "%s:", file.Name)

	if err !\= nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	fmt.Println()

	return nil
}

func main() {
	read, err := zip.OpenReader("test.zip")
	if err !\= nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for \_, file := range read.File {
		if err := listFiles(file); err !\= nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}

---

## Go program to extracting or unzip a Zip format file

package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	zipReader, \_ := zip.OpenReader("test.zip")
	for \_, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err !\= nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O\_WRONLY|os.O\_CREATE|os.O\_TRUNC,
				file.Mode(),
			)
			if err !\= nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			\_, err = io.Copy(outputFile, zippedFile)
			if err !\= nil {
				log.Fatal(err)
			}
		}
	}
}
