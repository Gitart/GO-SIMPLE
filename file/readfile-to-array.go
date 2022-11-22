package main

import (
     "bufio"
     "fmt"
     "io"
     "os"
     "regexp"
     "log"
)

func main(){
	fmt.Println("Start...")
	ReadEachline("r.txt")
}

// *********************************************
// Чтение файла построчно
// С добавлением в массив
// *********************************************
func ReadEachline(namefile string) {
	file, err := os.Open(namefile)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

	fmt.Println("eachline----", txtlines[2])
}

// *********************************************
// Чтение по словно - 
// *********************************************
func wordByWord1(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	
	r := bufio.NewReader(f)


for {

	line, err := r.ReadString('\n')
	
	if err == io.EOF {
	   break
    } else if err != nil {
      fmt.Printf("error reading file %s", err)
      return err

      r := regexp.MustCompile("[^\\s]+")
      words := r.FindAllString(line, -1)
   
      for i := 0; i < len(words); i++ {
       fmt.Println("word : ",words[i])
      }
  }
}

return nil
}
