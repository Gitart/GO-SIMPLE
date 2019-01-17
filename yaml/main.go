package main

import (
	"fmt"
	"flag"
	"os"
	"runtime"
//"github.com/fatih/color"
  "gopkg.in/yaml.v2"
  "log"
  "io/ioutil"
  "strings"
)

var (
  // Command line flags.
	httpMethod      string
)

func init() {
	  // flag.StringVar(&httpMet, "X", "GET", "HTTP method to use")
	  // flag.Usage = usage
	  Port := flag.String("p", "1968", "Input Port")   
	  // flag.StringVar(&Port, "p", "1968", "Input Port")   
	  flag.Parse()
	  fmt.Println("Port:", *Port)
}

func usage() {
	   fmt.Fprintf(os.Stderr,  "Usage: %s [OPTIONS] URL\n\n", os.Args[0])
	   fmt.Fprintln(os.Stderr, "OPTIONS:")
}

// An example showing how to unmarshal embedded structs from YAML.
type BB struct {
	 E   string `yaml:"e"`
	 A   string `yaml:"a"`
}

type JX struct {
	 F   string `yaml:"f"`
	 J   string `yaml:"j"`
}

type LM struct {
	 Code   string `yaml:"code"`
	 Lode   string `yaml:"lode"`
}

type IOS struct {
	 Keypath   string `yaml:"key_path"`
	 Keytype   string `yaml:"key_type"`
	 Enb       bool   `yaml:"enabale"`
	 Password  string `yaml:"password"`
}

type API struct {
	 Pushuri       string `yaml:"push_uri"`
	 Stat_go_uri   string `yaml:"stat_go_uri"`
   Metric_uri    string `yaml:"metric_uri"`
}

type T struct {
	B           BB       
	X           JX
	Nams        LM
	Ios         IOS
	Api         API
	AA          string `yaml:"a"`
	SS          string `yaml:"s"`
	LL          string `yaml:"l"`
	DD          string `yaml:"d"`
}

type T1 struct {
        A string
        L string
        S string
        B struct {
                RenamedC string   `yaml:"e"`
                Re       string   `yaml:"a"`
        }
}

//*****************************************
// Main
//*****************************************
func main() {
     YamlReadFile()
     // YamlTest()
}

//*****************************************
// Main Procedure
//*****************************************
func  YamlTests() {
	    // Read YAML file
	    YamlReadFile()

	    // Read enviroument
      fmt.Println("VIEW:", os.Getenv("VIEW"))

      // Loop enviroument
	    for _, e := range os.Environ() {
	       pair := strings.Split(e, "=")
	       fmt.Println(pair[0]," = ",pair[1])
	    }
}

//*****************************************
// Read Yaml file for structure
//*****************************************
func YamlReadFile() {
    bs:= T{}
    
    content, err := ioutil.ReadFile("./ng.yaml")
	  
    if err != nil {
	     log.Fatal(err)
	  }

    yaml.Unmarshal([]byte(content), &bs)

    // Preview
    fmt.Println("B.E:",               bs.B.E)
	  fmt.Println("B.A:",               bs.B.A)
    fmt.Println("DD:",                bs.DD)
    fmt.Println("LL:",                bs.LL)
    fmt.Println("FJ:",                bs.X.F)
    fmt.Println("JR:",                bs.X.J)
    fmt.Println("Nam Code:",          bs.Nams.Code)
    fmt.Println("Nam Lode:",          bs.Nams.Lode)
    fmt.Println("Ios Keypath:",       bs.Ios.Keypath)
    fmt.Println("Ios Keytype:",       bs.Ios.Keytype)
    fmt.Println("Ios Enable:",        bs.Ios.Enb)
    fmt.Println("Ios Password:",      bs.Ios.Password)
    fmt.Println("Api.Pushuri:",       bs.Api.Pushuri)
    fmt.Println("Api.Stat_go_uri",    bs.Api.Stat_go_uri)
    fmt.Println("Api.Metric_uri",     bs.Api.Metric_uri)
}

// *****************************************
// Read Yaml string for structure
// *****************************************
func YamlTest(){

// bs:= StructB{}
bs:= T{}

data:=`
a:  Articles
s:  Stopping
l:  Location
d:  Dinamycs
b:
  e: enabale httpd server
  a: ip address to bind
x:
  f: Settings path
  j: Normal settings
nam:
  code: Codes             # Code for site and settings
  lode: Load              # Code for site and settings
ios:
  enabled:    false
  key_path:   "key.pem"
  key_base64: ""          # load iOS key from base64 input
  key_type:   "pem"       # could be pem, p12 or p8 type
  password:   "eee"       # certificate password, default as empty string.
  production: false
  max_retry:  0           # resend fail notification, default value zero is disabled
  key_id:     ""          # KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
  team_id:    ""          # TeamID from developer account (View Account -> Membership)  
`
	err := yaml.Unmarshal([]byte(data), &bs)

	if err != nil {
	   log.Fatalf("cannot unmarshal data: %v", err.Error())
	}

  // Preview
	fmt.Println("B.E:",           bs.B.E)
	fmt.Println("B.A:",           bs.B.A)
  fmt.Println("DD:",            bs.DD)
  fmt.Println("LL:",            bs.LL)
  fmt.Println("FJ:",            bs.X.F)
  fmt.Println("JR:",            bs.X.J)
  fmt.Println("Nam Code:",      bs.Nams.Code)
  fmt.Println("Nam Lode:",      bs.Nams.Lode)
  fmt.Println("Ios Keypath:",   bs.Ios.Keypath)
  fmt.Println("Ios Keytype:",   bs.Ios.Keytype)
  fmt.Println("Ios Enable:",    bs.Ios.Enb)
  fmt.Println("Ios Password:",  bs.Ios.Password)

  flag.Parse()
  version:="version 123"
	
	fmt.Printf("%s %s (runtime: %s)\n", os.Args[0], version, runtime.Version())
	os.Exit(0)
}
