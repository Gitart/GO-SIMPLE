## GORM




## Save to Database
```go

// Запись в базу данных без сохранения PromptItems + Items
func (p *Prompts) SavePromptforAi() string {

	// Сохраняем только сам Prompts, без PromptItems
	// Если нужно сохранить PromptItems, то нужно использовать отдельный метод
	p.CreatedAt = time.Now() // Устанавливаем дату создания
	rec := db.DB.Create(&p)

	// Проверяем на ошибки при сохранении Prompts
	if rec.Error != nil {
		fmt.Println("🔴 Error saving prompt:", rec.Error)
		return "🔴 Error saving prompt: " + rec.Error.Error()
	}


	return ""
}

```

## Models

```go
package ai

import (
	"time"
)

type (
	Settings struct {
		Id        int64     `json:"id"`         // Идентификатор настроек
		UserId    int64     `json:"user_id"`    // Идентификатор пользователя
		Username  string    `json:"username"`   // Имя пользователя
		Email     string    `json:"email"`      // Email пользователя
		Title     string    `json:"title"`      // Название настроек
		CreatedAt time.Time `json:"created_at"` // Дата создания настроек
		KeyApy    string    `json:"key_api"`    // Ключ API для доступа к AI сервисам
	}

	Prompts struct {
		Id         int64         `json:"id"`                                             // Идентификатор промпта
		Title      string        `json:"title"`                                          // Название промпта
		CreatedAt  time.Time     `json:"created_at"`                                     // Дата создания промпта
		Header     string        `json:"header"`                                         // Заголовок промпта
		Prompt     string        `json:"prompt"`                                         // Текст промпта
		Rating     float64       `json:"rating"`                                         // Оценка промпта
		Category   string        `json:"category"`                                       // Категория промпта
		Sources    string        `json:"sources"`                                        // Источник промпта
		Activ      int64         `json:"activ"`                                          // Использовать ли этот промпт
		Importance float64       `json:"importance"`                                     // Важность промпта
		Tokens     int64         `json:"tokens"`                                         // Количество токенов в промпте
		Items      []PromptItems `json:"items" gorm:"foreignKey:IdPrompt;references:Id"` // Список промптов
	}

	PromptItems struct {
		Id         int64     `json:"id"`         // Идентификатор промпта
		IdPrompt   int64     `json:"id_prompt"`  // Идентификатор промпта
		Title      string    `json:"title"`      // Название промпта
		CreatedAt  time.Time `json:"created_at"` // Дата создания промпта
		Header     string    `json:"header"`     // Заголовок промпта
		Prompt     string    `json:"prompt"`     // Текст промпта
		Activ      int64     `json:"activ"`      // Использовать ли этот промпт
		Importance float64   `json:"importance"` // Важность промпта
		Category   string    `json:"category"`   // Категория промпта -то что будет использоваться в AI

	}
)
```



## Var

```go

var (
	PromptExample = Prompts{
		Title:      "Test Gemini Prompt",
		Header:     "Test Gemini Header",
		Prompt:     "This is a test prompt for Gemini AI.",
		Rating:     4.5,
		Category:   "Test",
		Sources:    "Test Source",
		Activ:      1,
		Importance: 1.0,
		Tokens:     100,
		Items: []PromptItems{
			{
				Title:      "Пример для Gemini AI",
				CreatedAt:  time.Now(),
				Header:     "Пример для Gemini AI",
				Prompt:     "Это пример запроса для ",
				Activ:      1,
				Importance: 1.0,
				Category:   "Киев",
			},
			{
				Title:      "Пример для Gemini AI 2",
				CreatedAt:  time.Now(),
				Header:     "Пример для Gemini AI",
				Prompt:     "Это пример запроса для Gemini AI.",
				Activ:      1,
				Importance: 1.0,
				Category:   "Киев",
			},
			{
				Title:      "Состояние ",
				CreatedAt:  time.Now(),
				Header:     "Состояние ",
				Prompt:     " Gemini AI.",
				Activ:      1,
				Importance: 4.0,
				Category:   "Киев",
			},
		},
	}
)
```

