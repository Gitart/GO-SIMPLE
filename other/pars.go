package main

import (
        "fmt"

        "github.com/fatih/structtag"
)

func main() {
        tag := `json:"foo,omitempty,string" xml:"foo"`

        // parse the tag
        tags, err := structtag.Parse(string(tag))
        if err != nil {
                panic(err)
        }

        // iterate over all tags
        for _, t := range tags.Tags() {
                fmt.Printf("tag: %+v\n", t)
        }

        // get a single tag
        jsonTag, err := tags.Get("json")
        if err != nil {
                panic(err)
        }

        // change existing tag
        jsonTag.Name = "foo_bar"
        jsonTag.Options = nil
        tags.Set(jsonTag)

        // add new tag
        tags.Set(&structtag.Tag{
                Key:     "hcl",
                Name:    "foo",
                Options: []string{"squash"},
        })

        // print the tags
        fmt.Println(tags) // Output: json:"foo_bar" xml:"foo" hcl:"foo,squash"
}
