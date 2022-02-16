## Sample
```golang
// curl -v -F "should_be_bound=test" -F "ShouldNotBeBound=nope" http://localhost:8080/
func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		type Thing struct {
			ShouldBeBound    string `form:"should_be_bound"`
			ShouldNotBeBound string // `form:"-"`
		}

		fields := Thing{}
		if err := c.Bind(&fields); err != nil {
			return err
		}

		log.Printf("%+v\n", fields)
		return c.String(http.StatusOK, "OK\n")
	})

	log.Fatal(e.Start(":8080"))
}
```

## Sample
```go
c.Bind() binds query params along with body only with get/delete methods. Documentations mentions that https://echo.labstack.com/guide/binding/

You can bind body and query separately but know that - if body has empty fields it will empty fields that query has bound. ie. overwrite data that BindQueryParams had filled (this is how standard library json unmarshaller works)

		if err := (&echo.DefaultBinder{}).BindQueryParams(c, u); err != nil {
			return err
		}
		if err := (&echo.DefaultBinder{}).BindBody(c, u); err != nil {
			return err
		}
ps. you need to add struct tags for query params also

type User struct {
	Name  string `json:"name" query:"name"`
	Email string `json:"email" query:"email"`
}
Also see the difference of these requests

curl -v -H 'Content-Type: application/json'   -d '{"name":"Joe"}' "http://localhost:8080/users?name=test&email=test1"

{"name":"Joe","email":"test1"}
vs

curl -v -H 'Content-Type: application/json'   -d '{"name":"Joe","email":""}' "http://localhost:8080/users?name=test&email=test1"
{"name":"Joe","email":""}
```

## BindingQuery

```go
type User struct {
	Name  string `json:"name" query:"name"`
	Email string `json:"email" query:"email"`
}

func main() {
	e := echo.New()

	e.POST("/users", func(c echo.Context) error {

		name := c.QueryParam("name")
		email := c.QueryParam("email")
		fmt.Printf("%s ||| %s \n", name, email)

		u := new(User)
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, u); err != nil {
			return err
		}
		if err := (&echo.DefaultBinder{}).BindBody(c, u); err != nil {
			return err
		}

		//No output here
		fmt.Printf("%s ---%s \n", u.Name, u.Email)
		return c.JSON(http.StatusCreated, u)

	})
	e.Logger.Fatal(e.Start(":8080"))
}
```
