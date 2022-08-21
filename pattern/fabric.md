# **Абстрактная фабрика** на Go

**Абстрактная фабрика** — это порождающий паттерн проектирования, который решает проблему создания целых семейств связанных продуктов, без указания конкретных классов продуктов.

Абстрактная фабрика задаёт интерфейс создания всех доступных типов продуктов, а каждая конкретная реализация фабрики порождает продукты одной из вариаций. Клиентский код вызывает методы фабрики для получения продуктов, вместо самостоятельного создания с помощью оператора `new`. При этом фабрика сама следит за тем, чтобы создать продукт нужной вариации.

[Подробней об Абстрактной фабрике](https://refactoring.guru/ru/design-patterns/abstract-factory)

Навигация

 [Интро](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#)

 [Концептуальный пример](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0)

 [i­Sports­Factory](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iSportsFactory-go)

 [adidas](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidas-go)

 [nike](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nike-go)

 [i­Shoe](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iShoe-go)

 [adidas­Shoe](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidasShoe-go)

 [nike­Shoe](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nikeShoe-go)

 [i­Shirt](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iShirt-go)

 [adidas­Shirt](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidasShirt-go)

 [nike­Shirt](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nikeShirt-go)

 [main](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--main-go)

 [output](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--output-txt)

## Концептуальный пример

Представим, что вам нужно купить спортивную форму, состоящую из двух разных вещей: пара обуви и футболка. Вы хотите приобрести полный набор от одного бренда, чтобы вещи сочитались между собой.

Переводя вышесказаное в код, абстрактная фабрика поможет нам создавать наборы продуктов, которые всегда будут подходить друг к другу.

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iSportsFactory-go)**iSportsFactory.go:** Интерфейс абстрактной фабрики

package main

import "fmt"

type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
    if brand \== "adidas" {
        return &Adidas{}, nil
    }

    if brand \== "nike" {
        return &Nike{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidas-go)**adidas.go:** Конкретная фабрика

package main

type Adidas struct {
}

func (a \*Adidas) makeShoe() IShoe {
    return &AdidasShoe{
        Shoe: Shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a \*Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            logo: "adidas",
            size: 14,
        },
    }
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nike-go)**nike.go:** Конкретная фабрика

package main

type Nike struct {
}

func (n \*Nike) makeShoe() IShoe {
    return &NikeShoe{
        Shoe: Shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n \*Nike) makeShirt() IShirt {
    return &NikeShirt{
        Shirt: Shirt{
            logo: "nike",
            size: 14,
        },
    }
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iShoe-go)**iShoe.go:** Абстрактный продукт

package main

type IShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type Shoe struct {
    logo string
    size int
}

func (s \*Shoe) setLogo(logo string) {
    s.logo \= logo
}

func (s \*Shoe) getLogo() string {
    return s.logo
}

func (s \*Shoe) setSize(size int) {
    s.size \= size
}

func (s \*Shoe) getSize() int {
    return s.size
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidasShoe-go)**adidasShoe.go:** Конкретный продукт

package main

type AdidasShoe struct {
    Shoe
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nikeShoe-go)**nikeShoe.go:** Конкретный продукт

package main

type NikeShoe struct {
    Shoe
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--iShirt-go)**iShirt.go:** Абстрактный продукт

package main

type IShirt interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type Shirt struct {
    logo string
    size int
}

func (s \*Shirt) setLogo(logo string) {
    s.logo \= logo
}

func (s \*Shirt) getLogo() string {
    return s.logo
}

func (s \*Shirt) setSize(size int) {
    s.size \= size
}

func (s \*Shirt) getSize() int {
    return s.size
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--adidasShirt-go)**adidasShirt.go:** Конкретный продукт

package main

type AdidasShirt struct {
    Shirt
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--nikeShirt-go)**nikeShirt.go:** Конкретный продукт

package main

type NikeShirt struct {
    Shirt
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--main-go)**main.go:** Клиентский код

package main

import "fmt"

func main() {
    adidasFactory, \_ :\= GetSportsFactory("adidas")
    nikeFactory, \_ :\= GetSportsFactory("nike")

    nikeShoe :\= nikeFactory.makeShoe()
    nikeShirt :\= nikeFactory.makeShirt()

    adidasShoe :\= adidasFactory.makeShoe()
    adidasShirt :\= adidasFactory.makeShirt()

    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShirtDetails(s IShirt) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

#### [](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0--output-txt)**output.txt:** Результат выполнения

Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
