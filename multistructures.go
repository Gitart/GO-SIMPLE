package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(`{
        "node1": {
            "value": "1",
            "node2": {
                "value": "2",
                "node4": {
                    "value": "4"
                }
            },
            "node3": {
                "value": "3"
            }
        }
    }`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat["node1"].(map[string]interface{})["node2"].(map[string]interface{})["node4"].(map[string]interface{})["value"])
}
