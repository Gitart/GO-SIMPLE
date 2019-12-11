# Если карта не является ссылочной переменной, что это?

В моем [предыдущем посте](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go) я показал, что карты Go не являются ссылочными переменными и не передаются по ссылке. Это оставляет вопрос, если карты не являются ссылочными переменными, что они?

Для нетерпеливых ответ таков:

> Значение карты \- это указатель на `runtime.hmap`  структуру.

Если вы не удовлетворены этим объяснением, читайте дальше.

# Какой тип значения карты?

Когда вы пишете заявление

m: = make (map \[int\] int)

Компилятор заменяет его вызовом [`runtime.makemap`]( https://golang.org/src/runtime/hashmap.go#L222) , который имеет подпись

// makemap реализует создание карты Go (make \[k\] v, подсказка) // Если компилятор определил, что карта или первый сегмент // может быть создан в стеке, h и / или bucket могут отличаться от nil. // Если h! = Nil, карту можно создать непосредственно в h. // Если bucket! = Nil, то bucket можно использовать в качестве первого bucket. func makemap (t \* maptype, hint int64, h \* hmap, bucket unsafe.Pointer) \* hmap

Как видите, тип возвращаемого значения `runtime.makemap` \- это указатель на [`runtime.hmap`](https://golang.org/src/runtime/hashmap.go#L106) структуру. Мы не можем видеть это из нормального кода Go, но мы можем подтвердить, что значение карты имеет тот же размер, что и одно `uintptr` машинное слово.

основной пакет
 импорт ( "FMT" «Небезопасный» )
 func main () { var m map \[int\] int var p uintptr fmt.Println (unsafe.Sizeof (m), unsafe.Sizeof (p)) // 8 8 (linux / amd64) }

# Если карты являются указателями, не должны ли они иметь значение \* map \[key\]?

Хороший вопрос, что если карты являются значениями указателя, почему выражение `make(map[int]int)` возвращает значение с типом `map[int]int` . Разве это не должно вернуть `*map[int]int` ? Ян Тейлор [ответил на этот вопрос в последнее время в golang гайки](https://groups.google.com/d/msg/golang-nuts/SjuhSYDITm4/jnrp7rRxDQAJ) резьбы 1 .

> В первые дни то, что мы называем картами, теперь было написано как указатели, поэтому вы написали \* map \[int\] int. Мы отошли от этого, когда поняли, что никто никогда не писал \`map\` без написания\` \* map\`.

Возможно переименование типа от \* `map[int]int` до `map[int]int` , а толку , так как тип не *выглядеть* как указатель, был менее запутанным , чем профилированного значение указателя , которое не может быть разыменовано.

# Вывод

Карты, как и каналы, но в *отличие от* слайсов, являются просто указателями на `runtime` типы. Как вы видели выше, карта \- это просто указатель на `runtime.hmap` структуру.

Карты имеют ту же семантику указателя, что и любое другое значение указателя в программе Go. Нет ничего волшебного, кроме переписывания синтаксиса карты компилятором в вызовах функций в `runtime/hmap.go` .