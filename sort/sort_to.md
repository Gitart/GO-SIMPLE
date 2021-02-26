# Алгоритмы сортировки на Go II

Posted on 2019, Feb 16 2 мин. чтения

## InsertionSort (Сортировка вставками)

У нас массив разделен на две области \- слева отсортированные значения, справа неотсортированные. Берем элемент из неотсортированной области, последовательно просматриваем отсортированные элементы справа налево, находим место для вставки и переставляем элемент.

```go
func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
```

[https://play.golang.org/p/J\_g\-oPeumg0](https://play.golang.org/p/J_g-oPeumg0)

Полезные ссылки:

*   [http://algolist.manual.ru/sort/insert\_sort.php](http://algolist.manual.ru/sort/insert_sort.php)
*   [https://en.wikipedia.org/wiki/Insertion\_sort](https://en.wikipedia.org/wiki/Insertion_sort)

## ShellSort (Сортировка Шелла)

Вариация InsertionSort, где мы меняем шаг сортировки. Усовершенстование в том, что за счет более крупного шага, мы переставляем элементы сразу на большие дистанции (если повезет).

```go
func ShellSort(ar []int) {
	for gap := len(ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ar); i++ {
			x := ar[i]
			j := i
			for ; j >= gap && ar[j-gap] > x; j -= gap {
				ar[j] = ar[j-gap]
			}
			ar[j] = x
		}
	}
}
```

[https://play.golang.org/p/LhZfuONgTHl](https://play.golang.org/p/LhZfuONgTHl)

Полезные ссылки:

*   [http://algolist.manual.ru/sort/shell\_sort.php](http://algolist.manual.ru/sort/shell_sort.php)
*   [https://en.wikipedia.org/wiki/Shellsort](https://en.wikipedia.org/wiki/Shellsort)
*   [Сортировка Шелла](https://ru.wikipedia.org/wiki/%D0%A1%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B0_%D0%A8%D0%B5%D0%BB%D0%BB%D0%B0)

## HeapSort (Пирамидальная сортировка)

Делим массив на две части \- не отсортированная и отсортированная. На не отсортированной части строим упорядоченное бинарное дерево (heap). (Строим \- имеется ввиду мапим ячейки массива на листья бинарного дерева.) Строим по правилу, что потомки не больше предка, корневую ноду располагаем по индексу 0, а индексы листьев вычисляем по формуле 2i+1, 2i+2.

В итоге, в корневой ноде (ar\[0\]) оказывается наибольшее значение. Дальше мы меняем местами значения в начале и в конце не упорядоченной зоны. Получается, что отсортированная часть выросла на 1 элемент, не отсортированная уменьшилась на 1. Опять перестраиваем дерево, и повторяем все шаги пока полностью не отсортируем.

![](https://xn----7sbaabdq1a3ayardngchc4a.xn--p1ai/%D0%B7%D0%B0%D0%BC%D0%B5%D1%82%D0%BA%D0%B8/%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D1%8B-%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B8-%D0%BD%D0%B0-go-ii/img/heap_sort.png)

```go
func HeapSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	heapIt(ar)

	ar[0], ar[len(ar)-1] = ar[len(ar)-1], ar[0]

	HeapSort(ar[:len(ar)-1])
}

func heapIt(ar []int) {

	// left = 2*i + 1
	// right = 2*i + 2
	// i = 0
	// мы каждый раз рассматриваем 3 ноды - рут и 2 листа
	// и по кажлому листу повторяем рекурсивно heapIt(ar[1:]), heapIt(ar[2:])
	if len(ar) < 2 {
		return
	}

	if len(ar) == 2 {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}

	if len(ar) > 3 {
		heapIt(ar[1:])
		heapIt(ar[2:])
	}

	if ar[1] > ar[2] {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
	} else {
		if ar[0] < ar[2] {
			ar[0], ar[2] = ar[2], ar[0]
		}
	}
}
```

[https://play.golang.org/p/cchsTOhNol5](https://play.golang.org/p/cchsTOhNol5)

Полезные ссылки:

*   [http://algolist.manual.ru/sort/pyramid\_sort.php](http://algolist.manual.ru/sort/pyramid_sort.php)
*   [https://en.wikipedia.org/wiki/Heapsort](https://en.wikipedia.org/wiki/Heapsort)
