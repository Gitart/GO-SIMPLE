// Compilation:
//    $ go build merge.go
//
// Build:
//    $ ./merge FILE

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func matched(pat *regexp.Regexp, b []byte) string {
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(b, -1) {
		res = pat.Expand(res, []byte("$file"), b, s)
	}
	return string(res)
}

func openTex(file string, pat *regexp.Regexp) {
	f, err := os.Open(file)
	checkFatal(err)
	defer f.Close()

	dir := filepath.Dir(file)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		checkFatal(err)

		if strings.HasPrefix(line, "%") {
			continue
		}
		if pat.Match([]byte(line)) {
			m := matched(pat, []byte(line))
			openTex(filepath.Join(dir, m+".tex"), pat)
		} else {
			fmt.Printf(line)
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 1 {
		pat, err := regexp.Compile(`\\(input|include){(?P<file>\w+)}`)
		checkFatal(err)
		openTex(flag.Arg(0), pat)
	} else {
		fmt.Printf("Usage %s FILE\n", os.Args[0])
	}
}
