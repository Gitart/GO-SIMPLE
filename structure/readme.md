## Использование тегов структур в Go
![изображение](https://user-images.githubusercontent.com/3950155/166150492-7ff3904a-76dd-4b8b-9e16-cee6effc52b0.png)

Введение

Структуры используются для сбора различных элементов информации внутри одной единицы. Эти наборы информации используются для описания концепций более высокого уровня. Так, адрес состоит из области, города, улицы, почтового индекса и т. д. Когда вы считываете эту информацию из баз данных, API или других подобных систем, вы можете использовать теги структур для контроля присвоения этой информации в поля структуры. Структурные теги — это небольшие элементы метаданных, прикрепленные к полям структуры. Они содержат инструкции для другого кода Go, который работает с этой структурой.
Как выглядит структурный тег?

Структурные теги в Go представляют собой аннотации, которые отображаются после типа в декларации структуры Go. Каждый тег состоит из коротких строк, которым назначены определенные значения.

Структурный тег выделяется символами апострофа ````` и выглядит следующим образом:

Как выглядит структурный тег?

Структурные теги в Go представляют собой аннотации, которые отображаются после типа в декларации структуры Go. Каждый тег состоит из коротких строк, которым назначены определенные значения.

Структурный тег выделяется символами апострофа ````` и выглядит следующим образом:
```go
type User struct {
	Name string `example:"name"`
}
```

Другой код Go может оценивать структуры и извлекать значения, назначенные определенным ключам, которые он запрашивает. Структурные теги не влияют на работу кода без кода, который их использует.

С помощью этого примера вы увидите, как выглядят структурные теги, и как они не действуют без кода из другого пакета.

```go
package main

import "fmt"

type User struct {
	Name string `example:"name"`
}

func (u *User) String() string {
	return fmt.Sprintf("Hi! My name is %s", u.Name)
}

func main() {
	u := &User{
		Name: "Sammy",
	}

	fmt.Println(u)
}
```

Результат будет выглядеть так:

Output
Hi! My name is Sammy

В этом примере определяется тип User с полем Name. Для поля Name назначен структурный тег example:"name". Мы ссылаемся на этот тег как на “структурный тег example”, поскольку в качестве ключа в нем используется слово example. Структурный тег example имеет значение "name" для поля Name. Для типа User мы также определим метод String(), который требуется для интерфейса fmt.Stringer. Он вызывается автоматически при передаче типа в fmt.Println и позволяет нам вывести хорошо отформатированную версию нашей структуры.

В теле main мы создадим новый экземпляр типа User и передадим его в fmt.Println. Хотя в структуре имеется структурный тег, он не влияет на выполнение этого кода Go. Он выполняется точно так же, как если бы структурного тега не было.

Чтобы использовать структурные теги для каких-либо целей, необходимо написать другой код Go, который будет запрашивать их во время исполнения. В стандартной библиотеке имеются пакеты, которые используют структурные теги в своей работе. Наиболее популярный из них — пакет encoding/json.
Кодировка JSON

JavaScript Object Notation (JSON) представляет собой текстовый формат кодирования наборов данных, организованных по различным ключам строк. Он обычно используется для обмена данными между разными программами, поскольку имеет достаточно простой формат для расшифровки библиотеками многих разных языков. Ниже приведен пример JSON:

```go
{
  "language": "Go",
  "mascot": "Gopher"
}
```

Этот объект JSON содержит два ключа, language и mascot. За этими ключами идут связанные с ними значения. Ключ language имеет значение Go, а ключу mascot присвоено значение Gopher.

Кодировщик JSON в стандартной библиотеке использует структурные теги как аннотацию, указывая кодировщику, какие имена вы хотите присвоить полям в выводимых JSON результатах. Эти механизмы кодировки и декодировки JSON содержатся в пакете encoding/json.

В этом примере показана кодировка JSON без структурных тегов:
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string
	Password      string
	PreferredFish []string
	CreatedAt     time.Time
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
```

Этот код распечатывает следующее:

```
Output
{
  "Name": "Sammy the Shark",
  "Password": "fisharegreat",
  "CreatedAt": "2019-09-23T15:50:01.203059-04:00"
}
```


Мы определили структурный тег, описывающий пользователя с помощью полей, включая имя, пароль и время создания пользователя. В функции main мы создаем экземпляр этого пользователя, предоставляя значения для всех полей, кроме PreferredFish (Sammy нравятся все рыбы). Затем мы передаем экземпляр User в функцию json.MarshalIndent. Это позволяет нам просматривать результаты выполнения JSON в удобном виде без внешнего инструмента форматирования. Этот вызов можно заменить на json.Marshal(u) для получения JSON без дополнительных пробелов. Два дополнительных аргумента json.MarshalIndent определяют префикс результатов (который мы пропустили при выводе пустой строки) и символы отступа, в данном случае — два символа пробела. Любые ошибки json.MarshalIndent регистрируются, и программа завершает работу с помощью os.Exit(1). Наконец, мы кастуем []byte, возвращаемый json.MarshalIndent, в string, а затем передаем эту строку в функцию fmt.Println для печати на терминале.

Поля структуры соответствуют присвоенным им именам. Это не обычный стиль JSON, где используются названия полей, где первая буква каждого слова, кроме первого, — заглавная («верблюжий стиль»). Вы можете изменить имена полей в соответствии с «верблюжьим стилем», как показано в следующем примере. Как видно при выполнении этого образца, это не сработает, поскольку желаемые имена полей противоречат правилам Go в отношении имен экспортируемых полей.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	name          string
	password      string
	preferredFish []string
	createdAt     time.Time
}

func main() {
	u := &User{
		name:      "Sammy the Shark",
		password:  "fisharegreat",
		createdAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
```

В результате выводится следующее:
```
Output
{}
```

В этой версии мы изменили имена полей в соответствии с «верблюжьим стилем». Теперь Name соответствует name, Password соответствует password, а CreatedAt соответствует createdAt. В теле main мы изменили инициациализацию структуры для использования новых имен. Затем мы передаем структуру в функцию json.MarshalIndent, как и ранее. Сейчас в результате выводится пустой объект JSON, {}.

Для правильного отображения полей в «верблюжьем стиле» требуется, чтобы первый символ был в нижнем регистре. Хотя JSON не важны имена полей, для Go они имеют значение, поскольку от этого зависит видимость полей вне пакета. Поскольку пакет encoding/json является отдельным пакетом от используемого нами пакета main, первый символ его имени должен быть в верхнем регистре, чтобы он был видимым для encoding/json. Похоже мы в безвыходном положении, и нам нужен способ передать кодировщику JSON желаемое имя этого поля.
Использование структурных тегов для управления кодировкой

Вы можете изменить предыдущий пример так, чтобы экспортируемые поля правильно кодировались с именами полей в «верблюжьем стиле». Для этого мы аннотируем каждое поле структурным тегом. Структурный тег, распознаваемый encoding/json, имеет ключ json и значение, определяющее выводимый результат. Если мы поместим имена полей в «верблюжьем стиле» в качестве значения ключа json, кодировщик будет использовать эти имена. В данном примере решена проблема, наблюдавшаяся в предыдущих двух попытках:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
```

Результат будет выглядеть так:
```
Output
{
  "name": "Sammy the Shark",
  "password": "fisharegreat",
  "preferredFish": null,
  "createdAt": "2019-09-23T18:16:17.57739-04:00"
}
```

Мы снова изменили имена полей так, чтобы сделать их видимыми для других пакетов. Для этого мы сделали заглавными первые буквы имен этих полей. Однако в этот раз мы добавили структурные теги в форме json:"name", где "name" — имя, которое json.MarshalIndent должен использовать при печати нашей структуры в формате JSON.

Мы успешно и правильно отформатировали код JSON. Однако следует отметить, что поля для некоторых значений были распечатаны, хотя мы и не задавали эти значения. Если вы захотите, кодировщик JSON может ликвидировать эти поля.
Удаление пустых полей JSON

Чаще всего мы не хотим выводить поля, которые не заданы в JSON. Поскольку все типы Go имеют заданное по умолчанию «нулевое значение», пакету encoding/json требуется дополнительная информация, чтобы он считал поле не заданным, если оно имеет это нулевое значение. В части значения любого структурного тега json вы можете задать суффикс желаемого имени поля с опцией ,omitempty, чтобы кодировщик JSON не выводил это поле, если для него задано нулевое значение. В следующем примере мы устранили заметную в предыдущих примерах проблему вывода пустых полей:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
```

Результат выполнения будет выглядеть так:

Output
```
{
  "name": "Sammy the Shark",
  "password": "fisharegreat",
  "createdAt": "2019-09-23T18:21:53.863846-04:00"
}
```

Мы изменили предыдущие примеры так, что теперь поле PreferredFish имеет структурный тег json:"preferredFish,omitempty". Благодаря опции ,omitempty кодировщик JSON пропускает это поле, поскольку мы не задаем его. В предыдущих примерах результатом было значение null.

Теперь результаты выглядят намного лучше, однако мы по прежнему распечатываем пароль пользователя. Пакет encoding/json дает нам еще один способ полностью игнорировать конфиденциальные поля.
Игнорирование конфиденциальных полей

Некоторые поля необходимо экспортировать из структур, чтобы другие пакеты могли правильно взаимодействовать с типом. Однако эти поля могут носить конфиденциальный характер, и в данном случае мы хотим, чтобы кодировщик JSON полностью игнорировал поле—даже если оно задано. Для этого используется специальное значение - в качестве аргумента для структурного тега json:.

В этом примере мы исправили проблему раскрытия пароля пользователя.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
```

При запуске этого примера вы увидите следующие результаты:

Output
```
{
  "name": "Sammy the Shark",
  "createdAt": "2019-09-23T16:08:21.124481-04:00"
}
```

Единственное изменение в этом примере по сравнению с предыдущими заключается в том, что в поле пароля теперь используется специальное значение "-" для структурного тега json:. В результатах в этом примере мы видим, что поле password больше не отображается.

Эти возможности пакета encoding/json, ,omitempty и "-" не являются стандартными. Действия пакета со значениями структурного тега зависят от реализации. Поскольку пакет encoding/json является частью стандартной библиотеки, в других пакетах эти возможности также реализованы стандартным образом. Однако важно ознакомиться с документацией по любым сторонним пакетам, использующим теги структуры, чтобы узнать, что поддерживается, а что нет.
Заключение

Структурные теги дают мощные возможности улучшения функциональности кода, работающего с вашими структурами. Многие стандартные библиотеки и сторонние пакеты поддерживают индивидуальные настройки с помощью структурных тегов. Их эффективное использование в коде поддерживает персонализацию поведения, а также в них кратко документируется использование этих полей будущими разработчиками.
