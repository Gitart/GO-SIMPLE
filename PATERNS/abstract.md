# ПАТТЕРНЫ ПРОЕКТИРОВАНИЯ на Go
https://refactoring.guru/ru/design-patterns/go

## Каталог **Go**\-примеров

#### Порождающие паттерны

![Абстрактная фабрика](https://refactoring.guru/images/patterns/cards/abstract-factory-mini.png?id=4c3927c446313a38ce77)

#### Абстрактная фабрика

Abstract Factory

Позволяет создавать семейства связанных объектов, не привязываясь к конкретным классам создаваемых объектов.

[Главный раздел](https://refactoring.guru/ru/design-patterns/abstract-factory)

[Пример кода](https://refactoring.guru/ru/design-patterns/abstract-factory/go/example#example-0)

![Строитель](https://refactoring.guru/images/patterns/cards/builder-mini.png?id=19b95fd05e6469679752)

#### Строитель

Builder

Позволяет создавать сложные объекты пошагово. Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

[Главный раздел](https://refactoring.guru/ru/design-patterns/builder)

[Пример кода](https://refactoring.guru/ru/design-patterns/builder/go/example#example-0)

![Фабричный метод](https://refactoring.guru/images/patterns/cards/factory-method-mini.png?id=72619e9527893374b98a)

#### Фабричный метод

Factory Method

Определяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

[Главный раздел](https://refactoring.guru/ru/design-patterns/factory-method)

[Пример кода](https://refactoring.guru/ru/design-patterns/factory-method/go/example#example-0)

![Прототип](https://refactoring.guru/images/patterns/cards/prototype-mini.png?id=bc3046bb39ff36574c08)

#### Прототип

Prototype

Позволяет копировать объекты, не вдаваясь в подробности их реализации.

[Главный раздел](https://refactoring.guru/ru/design-patterns/prototype)

[Пример кода](https://refactoring.guru/ru/design-patterns/prototype/go/example#example-0)

![Одиночка](https://refactoring.guru/images/patterns/cards/singleton-mini.png?id=914e1565dfdf15f240e7)

#### Одиночка

Singleton

Гарантирует, что у класса есть только один экземпляр, и предоставляет к нему глобальную точку доступа.

[Главный раздел](https://refactoring.guru/ru/design-patterns/singleton)

[Наивный Одиночка](https://refactoring.guru/ru/design-patterns/singleton/go/example#example-0)

[Многопоточный Одиночка](https://refactoring.guru/ru/design-patterns/singleton/go/example#example-1)

#### Структурные паттерны

![Адаптер](https://refactoring.guru/images/patterns/cards/adapter-mini.png?id=b2ee4f681fb589be5a06)

#### Адаптер

Adapter

Позволяет объектам с несовместимыми интерфейсами работать вместе.

[Главный раздел](https://refactoring.guru/ru/design-patterns/adapter)

[Пример кода](https://refactoring.guru/ru/design-patterns/adapter/go/example#example-0)

![Мост](https://refactoring.guru/images/patterns/cards/bridge-mini.png?id=b389101d8ee8e23ffa1b)

#### Мост

Bridge

Разделяет один или несколько классов на две отдельные иерархии — абстракцию и реализацию, позволяя изменять их независимо друг от друга.

[Главный раздел](https://refactoring.guru/ru/design-patterns/bridge)

[Пример кода](https://refactoring.guru/ru/design-patterns/bridge/go/example#example-0)

![Компоновщик](https://refactoring.guru/images/patterns/cards/composite-mini.png?id=a369d98d18b417f255d0)

#### Компоновщик

Composite

Позволяет сгруппировать объекты в древовидную структуру, а затем работать с ними так, как будто это единичный объект.

[Главный раздел](https://refactoring.guru/ru/design-patterns/composite)

[Пример кода](https://refactoring.guru/ru/design-patterns/composite/go/example#example-0)

![Декоратор](https://refactoring.guru/images/patterns/cards/decorator-mini.png?id=d30458908e315af195cb)

#### Декоратор

Decorator

Позволяет динамически добавлять объектам новую функциональность, оборачивая их в полезные «обёртки».

[Главный раздел](https://refactoring.guru/ru/design-patterns/decorator)

[Пример кода](https://refactoring.guru/ru/design-patterns/decorator/go/example#example-0)

![Фасад](https://refactoring.guru/images/patterns/cards/facade-mini.png?id=71ad6fa98b168c11cb3a)

#### Фасад

Facade

Предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

[Главный раздел](https://refactoring.guru/ru/design-patterns/facade)

[Пример кода](https://refactoring.guru/ru/design-patterns/facade/go/example#example-0)

![Легковес](https://refactoring.guru/images/patterns/cards/flyweight-mini.png?id=422ca8d2f90614dce810)

#### Легковес

Flyweight

Позволяет вместить бóльшее количество объектов в отведённую оперативную память. Легковес экономит память, разделяя общее состояние объектов между собой, вместо хранения одинаковых данных в каждом объекте.

[Главный раздел](https://refactoring.guru/ru/design-patterns/flyweight)

[Пример кода](https://refactoring.guru/ru/design-patterns/flyweight/go/example#example-0)

![Заместитель](https://refactoring.guru/images/patterns/cards/proxy-mini.png?id=25890b11e7dc5af29625)

#### Заместитель

Proxy

Позволяет подставлять вместо реальных объектов специальные объекты-заменители. Эти объекты перехватывают вызовы к оригинальному объекту, позволяя сделать что-то до или после передачи вызова оригиналу.

[Главный раздел](https://refactoring.guru/ru/design-patterns/proxy)

[Пример кода](https://refactoring.guru/ru/design-patterns/proxy/go/example#example-0)

#### Поведенческие паттерны

![Цепочка обязанностей](https://refactoring.guru/images/patterns/cards/chain-of-responsibility-mini.png?id=36d85eba8d14986f0531)

#### Цепочка обязанностей

Chain of Responsibility

Позволяет передавать запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

[Главный раздел](https://refactoring.guru/ru/design-patterns/chain-of-responsibility)

[Пример кода](https://refactoring.guru/ru/design-patterns/chain-of-responsibility/go/example#example-0)

![Команда](https://refactoring.guru/images/patterns/cards/command-mini.png?id=b149eda017c0583c1e92)

#### Команда

Command

Превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

[Главный раздел](https://refactoring.guru/ru/design-patterns/command)

[Пример кода](https://refactoring.guru/ru/design-patterns/command/go/example#example-0)

![Итератор](https://refactoring.guru/images/patterns/cards/iterator-mini.png?id=76c28bb48f997b369659)

#### Итератор

Iterator

Даёт возможность последовательно обходить элементы составных объектов, не раскрывая их внутреннего представления.

[Главный раздел](https://refactoring.guru/ru/design-patterns/iterator)

[Пример кода](https://refactoring.guru/ru/design-patterns/iterator/go/example#example-0)

![Посредник](https://refactoring.guru/images/patterns/cards/mediator-mini.png?id=a7e43ee8e17e4474737b)

#### Посредник

Mediator

Позволяет уменьшить связанность множества классов между собой, благодаря перемещению этих связей в один класс-посредник.

[Главный раздел](https://refactoring.guru/ru/design-patterns/mediator)

[Пример кода](https://refactoring.guru/ru/design-patterns/mediator/go/example#example-0)

![Снимок](https://refactoring.guru/images/patterns/cards/memento-mini.png?id=8b2ea4dc2c5d15775a65)

#### Снимок

Memento

Позволяет делать снимки состояния объектов, не раскрывая подробностей их реализации. Затем снимки можно использовать, чтобы восстановить прошлое состояние объектов.

[Главный раздел](https://refactoring.guru/ru/design-patterns/memento)

[Пример кода](https://refactoring.guru/ru/design-patterns/memento/go/example#example-0)

![Наблюдатель](https://refactoring.guru/images/patterns/cards/observer-mini.png?id=fd2081ab1cff29c60b49)

#### Наблюдатель

Observer

Создаёт механизм подписки, позволяющий одним объектам следить и реагировать на события, происходящие в других объектах.

[Главный раздел](https://refactoring.guru/ru/design-patterns/observer)

[Пример кода](https://refactoring.guru/ru/design-patterns/observer/go/example#example-0)

![Состояние](https://refactoring.guru/images/patterns/cards/state-mini.png?id=f4018837e0641d1dade7)

#### Состояние

State

Позволяет объектам менять поведение в зависимости от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

[Главный раздел](https://refactoring.guru/ru/design-patterns/state)

[Пример кода](https://refactoring.guru/ru/design-patterns/state/go/example#example-0)

![Стратегия](https://refactoring.guru/images/patterns/cards/strategy-mini.png?id=d38abee4fb6f2aed909d)

#### Стратегия

Strategy

Определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

[Главный раздел](https://refactoring.guru/ru/design-patterns/strategy)

[Пример кода](https://refactoring.guru/ru/design-patterns/strategy/go/example#example-0)

![Шаблонный метод](https://refactoring.guru/images/patterns/cards/template-method-mini.png?id=9f200248d88026d8e79d)

#### Шаблонный метод

Template Method

Определяет скелет алгоритма, перекладывая ответственность за некоторые его шаги на подклассы. Паттерн позволяет подклассам переопределять шаги алгоритма, не меняя его общей структуры.

[Главный раздел](https://refactoring.guru/ru/design-patterns/template-method)

[Пример кода](https://refactoring.guru/ru/design-patterns/template-method/go/example#example-0)

![Посетитель](https://refactoring.guru/images/patterns/cards/visitor-mini.png?id=854a35a62963bec1d75e)

#### Посетитель

Visitor

Позволяет создавать новые операции, не меняя классы объектов, над которыми эти операции могут выполняться.

[Главный раздел](https://refactoring.guru/ru/design-patterns/visitor)
