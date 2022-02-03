package main

import "fmt"

func main() {
	dwarfs := make([]string, 10)
	dwarfs[0] = "zero"
	dwarfs[4] = "5zero"

	fmt.Printf("LEN:%v CAP: %v \n", len(dwarfs), cap(dwarfs))
	dwarfs = append(dwarfs, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида", "ddd", "cлонце", "Венера")

	dwarfs = append(dwarfs, "Церера1", "Плутно1", "Хаумеа1", "Макемаке1", "Эрида1", "ssss1")
	dwarfs = append(dwarfs, "Церера2", "Плутно2", "Хаумеа2", "Макемаке2", "Эрид2", "ssss2")

	//fmt.Printf("LEN:%v CAP: %v \n", len(dwarfs), cap(dwarfs))
	//	fmt.Printf("%s \n", dwarfs)

	//	dwarfs[8] = "ehith"

	cnt := 0
	for _, l := range dwarfs {
		cnt++
		fmt.Println(cnt, l)
	}

	fmt.Printf("LEN:%v CAP: %v \n", len(dwarfs), cap(dwarfs))
}
