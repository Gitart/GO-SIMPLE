# Persistence — How should I work with databases and data stores?
# Accessing databases is typically part of the core business logic. Therefore, it probably makes sense to include an e.g. *sql.DB pointer in the concrete implementation of your service.

```go
type MyService struct {
	db     *sql.DB
	value  string
	logger log.Logger
}

func NewService(db *sql.DB, value string, logger log.Logger) *MyService {
	return &MyService{
		db:     db,
		value:  value,
		logger: logger,
	}
}
```
