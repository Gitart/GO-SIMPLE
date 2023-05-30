## Work with choice from

![image](https://github.com/Gitart/GO-SIMPLE/assets/3950155/34e3302f-ce61-44c6-b834-159974ce7a7b)


## We have list - select
```html
                <div class="modal-body">
                    <select id="ch_clients" class="form-select" multiple style="height: 300px">
                        {{range .Company}}
                            <option id="ch-{{.Id}}" value="{{.Id}}">{{.Title}}</option>
                        {{end}}
                    </select>
                </div>
```

## Get selected values in list
```js
    var selectedValues = $('select[multiple]').val();
    $.ajax({
        method: "POST",
        url:  "/documents/choice",
        data: JSON.stringify({ selectedValues: selectedValues }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) {
            console.log("Выбрано елемента " + data)
        }
    });
```

## Get selected ID`s in list
```js
 var selectedIDs = $('#ch_clients[multiple] option:selected').map(function() {
        return this.id;
    }).get();
    console.log(selectedIDs)
```    

## Backend

**Get all values in backend**

```go

package orders

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type SelectionPayload struct {
	SelectedValues []string `json:"selectedValues"`
}

// Выбранные контрагенты в окне
func Choises(e echo.Context) error {
	s := SelectionPayload{}
	e.Bind(&s)

	cnt := len(s.SelectedValues)

	// Process the selected values
	for _, value := range s.SelectedValues {
		fmt.Println(value)
	}

	return e.JSON(200, cnt)
}
```
