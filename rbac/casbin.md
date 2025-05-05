## Casbin

```go
package rbac

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func Rbac(e echo.Context) error {
	// Initialize the Casbin enforcer
	enforcer, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		log.Println("Failed to create enforcer: %v", err)
	}

	// Load the policy from the CSV file
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Println("Failed to load policy: %v", err)
	}

	// Check if the user has permission to perform the action on the object
	ok, _ := enforcer.Enforce("alice", "data2", "read")
	fmt.Println(ok)

	ok, _ = enforcer.Enforce("admins", "data3", "write")
	fmt.Println("admins write - ", ok)

	ok, _ = enforcer.Enforce("admins", "data3", "read")
	fmt.Println("admins read - ", ok)

	return nil
}

```
## rbac_model.conf
https://casbin.org/docs/supported-models


```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```


## rbac_policy.csv

```csv
p, alice, data1, read
p, bob, data2, write
p, data2_admin, data2, read
p, data2_admin, data2, write
g, alice, data2_admin2
p, alice, data2, read
p, admins, data3, read
p, admins, data3, delete
```

