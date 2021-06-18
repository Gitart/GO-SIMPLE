# **Абстрактная фабрика** на Go

**Абстрактная фабрика** — это порождающий паттерн проектирования, который решает проблему создания целых семейств связанных продуктов, без указания конкретных классов продуктов.

Абстрактная фабрика задаёт интерфейс создания всех доступных типов продуктов, а каждая конкретная реализация фабрики порождает продукты одной из вариаций. Клиентский код вызывает методы фабрики для получения продуктов, вместо самостоятельного создания с помощью оператора `new`. При этом фабрика сама следит за тем, чтобы создать продукт нужной вариации.

[Подробней об Абстрактной фабрике](https://refactoring.guru/ru/design-patterns/abstract-factory)

Навигация

 [Интро](#)

 [Концептуальный пример](#example-0)

 [i­Sports­Factory](#example-0--iSportsFactory-go)

 [adidas](#example-0--adidas-go)

 [nike](#example-0--nike-go)

 [i­Shoe](#example-0--iShoe-go)

 [adidas­Shoe](#example-0--adidasShoe-go)

 [nike­Shoe](#example-0--nikeShoe-go)

 [i­Shirt](#example-0--iShirt-go)

 [adidas­Shirt](#example-0--adidasShirt-go)

 [nike­Shirt](#example-0--nikeShirt-go)

 [main](#example-0--main-go)

 [output](#example-0--output-txt)

## Концептуальный пример

Представим, что вам нужно купить спортивную форму, состоящую из двух разных вещей: пара обуви и футболка. Вы хотите приобрести полный набор от одного бренда, чтобы вещи сочитались между собой.

Переводя вышесказаное в код, абстрактная фабрика поможет нам создавать наборы продуктов, которые всегда будут подходить друг к другу.

#### [](#example-0--iSportsFactory-go)**iSportsFactory.go:** Интерфейс абстрактной фабрики

package main

import "fmt"

type iSportsFactory interface {
    makeShoe() iShoe
    makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
    if brand \== "adidas" {
        return &adidas{}, nil
    }

    if brand \== "nike" {
        return &nike{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}

#### [](#example-0--adidas-go)**adidas.go:** Конкретная фабрика

package main

type adidas struct {
}

func (a \*adidas) makeShoe() iShoe {
    return &adidasShoe{
        shoe: shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a \*adidas) makeShirt() iShirt {
    return &adidasShirt{
        shirt: shirt{
            logo: "adidas",
            size: 14,
        },
    }
}

#### [](#example-0--nike-go)**nike.go:** Конкретная фабрика

package main

type nike struct {
}

func (n \*nike) makeShoe() iShoe {
    return &nikeShoe{
        shoe: shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n \*nike) makeShirt() iShirt {
    return &nikeShirt{
        shirt: shirt{
            logo: "nike",
            size: 14,
        },
    }
}

#### [](#example-0--iShoe-go)**iShoe.go:** Абстрактный продукт

package main

type iShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type shoe struct {
    logo string
    size int
}

func (s \*shoe) setLogo(logo string) {
    s.logo \= logo
}

func (s \*shoe) getLogo() string {
    return s.logo
}

func (s \*shoe) setSize(size int) {
    s.size \= size
}

func (s \*shoe) getSize() int {
    return s.size
}

#### [](#example-0--adidasShoe-go)**adidasShoe.go:** Конкретный продукт

package main

type adidasShoe struct {
    shoe
}

#### [](#example-0--nikeShoe-go)**nikeShoe.go:** Конкретный продукт

package main

type nikeShoe struct {
    shoe
}

#### [](#example-0--iShirt-go)**iShirt.go:** Абстрактный продукт

package main

type iShirt interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type shirt struct {
    logo string
    size int
}

func (s \*shirt) setLogo(logo string) {
    s.logo \= logo
}

func (s \*shirt) getLogo() string {
    return s.logo
}

func (s \*shirt) setSize(size int) {
    s.size \= size
}

func (s \*shirt) getSize() int {
    return s.size
}

#### [](#example-0--adidasShirt-go)**adidasShirt.go:** Конкретный продукт

package main

type adidasShirt struct {
    shirt
}

#### [](#example-0--nikeShirt-go)**nikeShirt.go:** Конкретный продукт

package main

type nikeShirt struct {
    shirt
}

#### [](#example-0--main-go)**main.go:** Клиентский код

package main

import "fmt"

func main() {
    adidasFactory, \_ :\= getSportsFactory("adidas")
    nikeFactory, \_ :\= getSportsFactory("nike")

    nikeShoe :\= nikeFactory.makeShoe()
    nikeShirt :\= nikeFactory.makeShirt()

    adidasShoe :\= adidasFactory.makeShoe()
    adidasShirt :\= adidasFactory.makeShirt()

    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShoeDetails(s iShoe) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShirtDetails(s iShirt) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

#### [](#example-0--output-txt)**output.txt:** Результат выполнения

Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
