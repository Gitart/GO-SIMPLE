// Copyright 2012 Tetsuo Kiso.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
// This program aims at grep-like filter program.
// Currently, this code only finds files with specified "suffix".
// For example, the following command will return files whose name ends with 
// suffix ".rc" such as mail.rc
//
//   $ ./filter  /etc .rc

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func find(path string) []os.FileInfo {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

func Filter(path, pattern string, filter func([]os.FileInfo, string) []os.FileInfo) []os.FileInfo {
	entries := find(path)
	filtered := filter(entries, pattern)
	return filtered
}

func hasSuffix(es []os.FileInfo, pattern string) []os.FileInfo {
	var buf []os.FileInfo
	for _, e := range es {
		if strings.HasSuffix(e.Name(), pattern) {
			buf = append(buf, e)
		}
	}
	return buf
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("usage: ./filter path pattern")
		os.Exit(1)
	}
	path := flag.Arg(0)
	suffix := flag.Arg(1)

	filtered := Filter(path, suffix, hasSuffix)
	for _, e := range filtered {
		fmt.Println(e.Name())
	}
}
