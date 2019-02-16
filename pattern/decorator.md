# The Decorator 
добавляет новые функциональные возможности к существующему объекту, не изменяя его структуру. 
Это структурный шаблон, так как этот шаблон действует как оболочка для существующего класса.

Например, структура декоратора, которая украшает (оборачивает) исходный объект и обеспечивает 
дополнительную функциональность, сохраняя сигнатуру его методов без изменений.

## Цель
Прикрепите дополнительные обязанности к объекту динамически.
Декораторы предоставляют гибкую альтернативу наследования для расширения функциональности.
Завернуть подарок, положить его в коробку и обернуть коробку.

Схема шаблона проектирования  

Шаблон Decorator имеет следующие объекты:

## Диаграмма классов декораторов
**Component** определяет интерфейс для объектов, к которым можно динамически добавлять обязанности.    
**ConcreteComponent** определяет объект, к которому могут быть прикреплены дополнительные обязанности.    
**Decorator**  поддерживает ссылку на объект Компонента и определяет интерфейс, который соответствует интерфейсу Компонента.   
**ConcreteDecorator**  добавляет обязанности к компоненту.

## Реализация
Мы рассмотрим использование шаблона декоратора на следующем примере, в котором мы расширим существующий объект, который извлекает данные из веб-службы. Мы украсим его, добавив возможности автоматического выключателя без изменения интерфейса структуры.

Позволяет иметь Fetcherинтерфейс, который определяет контракт на выборку некоторых данных из разных источников.

```golang
// Args of fetching function
type Args map[string]string

// Data returned by fetch
type Data map[string]string

// Fetcher fetches a data from remote endpoint
type Fetcher interface {
	// Fetch fetches the data
	Fetch(args Args) (Data, error)
}
```

Конкретной реализацией Fetcher интерфейса является Repository структура, которая предоставляет некоторые фиктивные данные, 
если предоставленные аргументы не пусты, в противном случае возвращает ошибку. 
Структура Repository является конкретным компонентом в контексте шаблона The Decorator.

```golang
// Repository of data
type Repository struct{}

// Fetch fetches data
func (r *Repository) Fetch(args Args) (Data, error) {
	if len(args) == 0 {
		return Data{}, fmt.Errorf("No arguments are provided")
	}

	data := Data{
		"user":     "root",
		"password": "swordfish",
	}
	fmt.Printf("Repository fetched data successfully: %v\n", data)
	return data, nil
}
```

### Retrier 
Структура является декоратором, который добавляет возможности автоматического выключателя к любому компоненту, 
который реализует Fetcherинтерфейс. 

Retrier Имеет несколько свойств , которые позволяют это. RetryCount Свойство определяет количество раз, 
что retrier должен попытаться извлечь , если есть ошибка. 

### WaitInterval 
Свойство определяет интервал между каждой повторной попыткой. 

### Fetcher
Свойство указывает на объект, который оформлен. 

В Retrier вызовах Fetch функции из декорированного объекта, пока не завершится успешно или превышать 
политику повторных попыток.

```golang
// Retrier retries multiple times
type Retrier struct {
	RetryCount   int
	WaitInterval time.Duration
	Fetcher      Fetcher
}

// Fetch fetches data
func (r *Retrier) Fetch(args Args) (Data, error) {
	for retry := 1; retry <= r.RetryCount; retry++ {
		fmt.Printf("Retrier retries to fetch for %d\n", retry)
		if data, err := r.Fetcher.Fetch(args); err == nil {
			fmt.Printf("Retrier fetched for %d\n", retry)
			return data, nil
		} else if retry == r.RetryCount {
			fmt.Printf("Retrier failed to fetch for %d times\n", retry)
			return Data{}, err
		}
		fmt.Printf("Retrier is waiting after error fetch for %v\n", r.WaitInterval)
		time.Sleep(r.WaitInterval)
	}

	return Data{}, nil
}
```

Затем мы можем добавить новые возможности повтора, обернув Repositoryэкземпляр с помощью Retrier:

```golang
repository := &cbreaker.Repository{}
retrier := &cbreaker.Retrier{
	RetryCount:   5,
	WaitInterval: time.Second,
	Fetcher:      repository,
}

data, err := repository.Fetch(cbreaker.Args{"id": "1"})
fmt.Printf("#1 repository.Fetch: %v\n", data)

data, err = retrier.Fetch(cbreaker.Args{})
fmt.Printf("#2 retrier.Fetch error: %v\n", err)

data, err = retrier.Fetch(cbreaker.Args{"id": "1"})
fmt.Printf("#3 retrier.Fetch: %v\n", data)
```

### решение 
Шаблон Decorator более удобен для добавления функциональных возможностей к объектам вместо целых структур 
во время выполнения. С отделкой также можно динамически удалять добавленные функции.
