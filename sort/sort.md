Алгоритмы сортировки на Go
 Posted on2019, Feb 06  3 мин. чтения
Стало вдруг интересно вспомнить некоторые алгоритмы сортировки и реализовать их на go.

Quicksort (Быстрая сортировка)
Основная идея - берем опорную точку (pivot), проходим массив, чтобы элементы, которые меньше опорной точки оказались слева от нее, а которые больше - справа. Дальше берем часть массива до опорной точки и вторую часть после опортной точки, повторяем на них сортировку. Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}
https://play.golang.org/p/4evoDVHJp6H

Полезные ссылки:

https://en.wikipedia.org/wiki/Quicksort
http://algolist.manual.ru/sort/quick_sort.php
BubbleSort (Сортировка пузырьком)
У нас два цикла - первым мы последовательно берем ячейку (текущая ячейка), где у нас будет (в итоге) наимениший из оставшихся элементов, а вторым циклом пробегаем от выбранного элемента до конца массива, сравниваем пары элементов, если есть меньшее значение чем в текущей ячейке, то меняем местами.

Или так - моделируем последовательность, как пузырьки, на каждом шаге наиболее легкий пузырек поднимается к верху.

Первый вариант у меня получился вот такой:

func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, i, j)
			}
		}
	}
}
https://play.golang.org/p/fsYoaAxFKGI

Потом я подглядел в вики и algolist, и получился такой вариант, где мы сравниваем соседние пары и гоним пузырек наверх:

func BubbleSort2(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				swap(ar, j-1, j)
			}
		}
	}
}
https://play.golang.org/p/QscHUwnoJoP

и вариант с тонущим пузырьком (sinking sort), когда “тяжелый” пузырек опускается вниз:

func BubbleSort3(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				swap(ar, j-1, j)
			}
		}
	}
}
https://play.golang.org/p/haO7NiJfYBF

Если третий вариант - “тонущий” пузырек, второй - “легкий”, то первый вариант можно назвать “газировкой” :) Т.к. пузырьки хаотически прыгают в начало списка и потом обратно на место более легких.

Полезные ссылки:

http://algolist.manual.ru/sort/bubble_sort.php
https://en.wikipedia.org/wiki/Bubble_sort
SelectSort (Сортировка выбором)
Основная идея состоит в том, что мы делим массив на две части - то, что уже отсортировано и остальное. На каждом шаге пробегаем не отсортированную последовательность, выбираем наименьший элемент и дополняем отсортированную часть.

Внешний цикл отщелкивает границу между отсортированным и не отсортированным, а внутренний цикл выбирает наименьшее из не отсортированной части.

func SelectSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}

		if min != i {
			swap(ar, i, min)
		}
	}
}
https://play.golang.org/p/YamBvzfrPSS

Полезные ссылки:

https://en.wikipedia.org/wiki/Selection_sort
http://algolist.manual.ru/sort/select_sort.php
Update
Идеоматический свап на го:

a, b = b, a
