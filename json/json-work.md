# How to Parse JSON in Golang (With Examples)

Updated on November 20, 2019

In this post, we will learn how to work with JSON in Go, in the simplest way possible.

We will look at different types of data that we encounter in Go, from structured data like *structs*, *arrays*, and *slices*, to unstructured data like *maps* and empty interfaces.

JSON is used as the de-facto standard for data serialization, and by the end of this post, you’ll get familiar with how to parse and encode JSON in Go

## [](#parsing-json-strings)Parsing JSON strings

The json package provided in Go’s standard library provides us with all the functionality we need. For any JSON string, the standard way to parse it is:

```go
import "encoding/json"
//...

// ...
myJsonString := `{"some":"json"}`

// `&myStoredVariable` is the address of the variable we want to store our
// parsed data in
json.Unmarshal([]byte(myJsonString), &myStoredVariable)
//...
```

What we will discuss in this post, is the different options you have for the type of `myStoredVariable`, and when you should use them.

There are two types of data you will encounter when working with JSON:

1.  Structured data
2.  Unstructured data

## [](#structured-data-decoding-json-to-structs)Structured data (Decoding JSON to Structs)

Since this is much easier, let’s deal with it first. This is the sort of data where you know the structure beforehand. For example, let’s say you have a bird object, where each bird has a `species` field and a `description` field :

```js
{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
}
```

To work with this kind of data, create a `struct` that mirrors the data you want to parse. In our case, we will create a bird struct which has a `Species` and `Description` attribute:

```go
type Bird struct {
  Species string
  Description string
}
```

And unmarshal it as follows:

```go
birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
//Species: pigeon, Description: likes to perch on rocks
```

[Try it here](https://play.golang.org/p/DtA6sEppLO)

> By convention, Go uses the same title cased attribute names as are present in the case insensitive JSON properties. So the `Species` attribute in our `Bird` struct will map to the `species`, or `Species` or `sPeCiEs` JSON property.

### [](#json-arrays)JSON Arrays

So what happens when you have an array of birds?

```js
[
  {
    "species": "pigeon",
    "decription": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
]
```

Since each element of the array is actually a `Bird`, you can actually unmarshal this, by just creating an array of birds :

```go
birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
var birds []Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf("Birds : %+v", birds)
//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
```

[Try it here](https://play.golang.org/p/thoUdxxmMa)

### [](#embedded-objects)Embedded objects

Now, consider the case when you have a property called `Dimensions`, that measures the `Height` and `Length` of the bird in question:

```js
{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
```

As with our previous examples, we need to mirror the structure of the object in question in our Go code. To add an embedded `dimensions` object, lets create a `dimensions` struct :

```go
type Dimensions struct {
  Height int
  Width int
}
```

Now, the `Bird` struct will include a `Dimensions` field:

```go
type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}
```

And can be unmarshaled using the same method as before:

```go
birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
var birds Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf(bird)
// {pigeon likes to perch on rocks {24 10}}
```

[Try it here](https://play.golang.org/p/zOUMUNH4w9)

### [](#primitives)Primitives

We mostly deal with complex objects or arrays when working with JSON, but it’s easy to forget that data like `3`, `3.1412` and `"birds"` are also perfectly valid JSON strings. We can unmarshal these values to their corresponding data type in Go:

```go
numberJson := "3"
floatJson := "3.1412"
stringJson := `"bird"`

var n int
var pi float64
var str string

json.Unmarshal([]byte(numberJson), &n)
fmt.Println(n)
// 3

json.Unmarshal([]byte(floatJson), &pi)
fmt.Println(pi)
// 3.1412

json.Unmarshal([]byte(stringJson), &str)
fmt.Println(str)
// bird
```

[Try it out](https://play.golang.org/p/usCx_5oESBd)

### [](#custom-attribute-names)Custom attribute names

I mentioned earlier that Go uses convention to ascertain the attribute name it should map a property to. Many times though, you want a different attribute name than the one provided in your JSON data.

```js
{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}
```

In the JSON data above, I would prefer `birdType` to remain as the `Species` attribute in my Go code. It is also not possible for me to provide a suitable attribute name for a key like `"what it does"`.

To solve this, we make use of struct field tags:

```go
type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
```

Now, we can explicitly tell our code which JSON property to map to which attribute.

```go
birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks}
```

[Try it here](https://play.golang.org/p/-_0XddCakR)

## [](#unstructured-data-decoding-json-to-maps)Unstructured data (Decoding JSON to maps)

If you have data whose structure or property names you are not certain of, you cannot use structs to unmarshal your data. Instead you can use maps. Consider some JSON of the form:

```js
{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
  },
  "animals": "none"
}
```

There is no struct we can build to represent the above data for all cases since the keys corresponding to the birds can change, which will change the structure.

To deal with this case we create a map of strings to empty interfaces:

```go
birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
var result map[string]interface{}
json.Unmarshal([]byte(birdJson), &result)

// The object stored in the "birds" key is also stored as
// a map[string]interface{} type, and its type is asserted from
// the interface{} type
birds := result["birds"].(map[string]interface{})

for key, value := range birds {
  // Each value is an interface{} type, that is type asserted as a string
  fmt.Println(key, value.(string))
}
```

[Try it here](https://play.golang.org/p/xbVxASrffo)

Each string corresponds to a JSON property, and its mapped `interface{}` type corresponds to the value, which can be of any type. The type is asserted from this `interface{}` type as is needed in the code. These maps can be iterated over, so a variable number of keys can be handled by a simple for loop.

## [](#encoding-json-from-go-data)Encoding JSON from Go data

The same rules that are used to decode a JSON string can be applied to encoding as well.

### [](#encoding-structured-data)Encoding structured data

Let’s consider our Go struct from before, and see the code required to get a JSON string from data of its type:

```go
package main

import (
	"encoding/json"
	"fmt"
)

// The same json tags will be used to encode data into JSON
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}
```

This will give the output:

```text
{"birdType":"Pigeon","what it does":"likes to eat seed"}
```

[Try it out](https://play.golang.org/p/jHGzfy-hEjT)

### [](#ignoring-empty-fields)Ignoring Empty Fields

In some cases, we would want to ignore a field in our JSON output, if its value is empty. We can use the `omitempty` property for this purpose.

For example, if the `Description` field is missing for the `pigeon` object, the key will not appear in the encoded JSON string incase we set this property:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
}
```

This will give us the output:

```text
{"birdType":"Pigeon"}
```

[Try it out](https://play.golang.org/p/jOihnSg8dGB)

If you want to learn about this in more detail, you can read my other post on [how “omitempty” works in Go.](https://www.sohamkamani.com/golang/2018-07-19-golang-omitempty/)

### [](#encoding-arrays-and-slices)Encoding Arrays and Slices

This isn’t much different from structs. We just need to pass the slice or array to the `json.Marshal` function, and it will encode data like you expect:

```go
pigeon := &Bird{
  Species:     "Pigeon",
  Description: "likes to eat seed",
}

// Now we pass a slice of two pigeons
data, _ := json.Marshal([]*Bird{pigeon, pigeon})
fmt.Println(string(data))
```

This will give the output:

```text
[{"birdType":"Pigeon","what it does":"likes to eat seed"},{"birdType":"Pigeon","what it does":"likes to eat seed"}]
```

[Try it out](https://play.golang.org/p/3a5lWlKSvPS)

### [](#encoding-maps)Encoding maps

We can use maps to encode unstructured data.

The keys of the map need to be strings, or a type that can convert to strings. The values can be any serializable type.

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
```

Output ([Try it out](https://play.golang.org/p/0A5pTy3onYZ)):

```text
{"birdSounds":{"eagle":"squak","pigeon":"coo"},"total birds":2}
```

## [](#what-to-use-structs-vs-maps)What to use (Structs vs maps)

As a general rule of thumb, if you *can* use structs to represent your JSON data, you should use them. The only good reason to use maps would be if it were not possible to use structs due to the uncertain nature of the keys or values in the data.
