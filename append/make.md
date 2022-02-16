
# :m: MAKE 

[Slice description](https://go.dev/blog/slices)


```go
 make([]string, 10, 12)
```
## Краткое описание

* Первый парметр - тип массива
* Второй параметр всегда должен указан - это длина маасива с проинициализированными элементами
* Третий параметр **cup()** - указывает на количество зарезервированных элементов для вставки и если его не укзать он всегда будет в два раза больше первого парметра. 
  Но до тех пор пока количество добавляемых элементов не превысит указанное количество элементов в первом параметре. Этот парметр всегда можно измерять **len()**. 
  По умолчанию это параметр равен первому параметру.
  При добавлении в масси  командой **append** элемнтов - второй параметр не меняется до тех пор пока количество добавляемых элементов не больше третьего параметра.
  


**ёмкость(cap)** - это *выделенная память* под элементы, при превышении размер *автоматически* увеличивается в **два** раза.   
**длина(len)** - это *инициализированная* память элементов, для превышения(добавления) нужно вручную использовать append.      

По умолчанию cap = len
Всегда будет cap >= len
Грубо говоря, **cap** выделяет память, а **len** инициализирует её всю или только часть .


Если кратко, то ёмкость слайса - длина массива, который хранит элементы слайса. Это Байты оперативки, выделенные под хранение ваших данных. 
Оно всегда больше или равно длины слайса. 
Сам слайс это указатель на какой-либо из элементов в этом масиве (не оьязательно первый) + длина. 
Это искуственное ограничение, фактически это отрезок массива, в котором что-то записано. 
При добавлении в слайс нового элемента слайс увеличит свою длину на 1, пооказывая что кол-во данных в массиве увеличилось на 1, 
а емкость самого массива не изменится, если её хватало для хранения нового элемента. Если же весь массив был заолнен, 
то слайс просто выделит новый массив в памяти большей ёмкости (обычно в два раза большей предыдущей ёмкости) 
и перенесёт все данные туда вместе с новым добавленным элементом. При этом, очевидно, сам слайс не изменит своей длины, т.к. 
Кол-во данных не изменилось, а только поменяет указатель на начальный элемент, т.к. Массив с данными выделен новый по новому адресу в памяти.



Массивы имеют фиксированный размер. 
Срез (slice) - это гибкое отображение элементов массива с возможностью динамического изменения размера. На практике срезы более распространены, чем массивы.

### Срезы это как указатели на массивы

Срез не хранит никаких данных, он всего лишь обозначает секцию нижележащего массива.

Изменения элементов среза приводят к модификации соответствующих элементов его нижележащего массива.

Другие срезы, имеющие общий нижележащий массив, также увидят эти изменения.

```
package main

import "fmt"

func main() {
  names := [4]string{
    "Евгений",
    "Иван",
    "Георгий",
    "Сергей",
  }
  fmt.Println(names)

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)
}

```

Вывод:

```
[Евгений Иван Георгий Сергей]
[Евгений Иван] [Иван Георгий]
[Евгений XXX] [XXX Георгий]
[Евгений XXX Георгий Сергей]
```

## Sample code
[Samples](https://go.dev/play/p/WCgYEaoyRWe)

```go
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
```


Output
```
LEN:10 CAP: 10 
1 zero
2 
3 
4 
5 5zero
6 
7 
8 
9 
10 
11 Церера
12 Плутно
13 Хаумеа
14 Макемаке
15 Эрида
16 ddd
17 cлонце
18 Венера
19 Церера1
20 Плутно1
21 Хаумеа1
22 Макемаке1
23 Эрида1
24 ssss1
25 Церера2
26 Плутно2
27 Хаумеа2
28 Макемаке2
29 Эрид2
30 ssss2
LEN:30 CAP: 40 
```


