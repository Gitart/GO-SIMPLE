# My note on Map Iterations and Handling Arbitrary Types with Interface in Go

[

![Joe Chasinga](https://miro.medium.com/fit/c/96/96/2*Zw3HZJzsXzJj4daj5CMu5Q.jpeg)

](https://medium.com/@jochasinga?source=post_page-----feef83ab9db2----------------------)

[Joe Chasinga](https://medium.com/@jochasinga?source=post_page-----feef83ab9db2----------------------)

Follow

[Nov 10, 2014](https://medium.com/code-zen/dynamically-creating-instances-from-key-value-pair-map-and-json-in-go-feef83ab9db2?source=post_page-----feef83ab9db2----------------------) · 4 min read

I’m writing this as my personal note as well as a quick guide to others who may find this useful. In this write\-up, I’ll be converting a JSON string to Go’s native instances. But I’m taking the long haul of constructing an empty interface to deal with arbitrary data types coming from JSON and encoding them with type assertions. If you’re looking for an idiomatic way to convert JSON to Go instances, see this [playground](http://play.golang.org/p/R3iouPgx4Y) by [Ed Robinson](http://www.eddrobinson.net/).

In Go, we iterate values in a map like this:

for key, value := range mymap {}

However, to transform key\-value pairs in our map into instances, I incorporate a counter variable to serve as a missing index in the loop:

type Person struct {
        Name string
        Age int
}
persons := map\[string\]int{
        "John"  : 35,
        "Jane"  : 23,
        "Mary"  : 12,
        "Shane" : 35,
        "Phil"  : 87,
}
counter := 0
pslice := make(\[\]\*Person, len(persons))
for k, v := range persons {
        pslice\[counter\] = &Person{k, v}
        counter++
}

This returns a slice filled with all the \*Person instances.

By using *Unmarshal* function in Go’s *json* package and an interface or struct, we can convert JSON to Go’s instances. To deal with JSON unpredictable structure, we introduce an empty interface that acts as an all\-purpose container to store the decoded data regardless of their types. The *json* package use the following map forms to store values from JSON bytes:

map\[string\]interface{}
\[\]interface{}

Here is what you might do provided a JSON byte array:

// JSON blob as an slice of bytes (char) in Go
b := \[\]byte(\`{"Persons" :\[
                 {"Name": "John", "Age" : 35 },
                 {"Name": "Jane", "Age" : 23 },
                 {"Name": "Mary", "Age" : 12 },
                 {"Name": "Shane", "Age": 35 },
                 {"Name": "Phil", "Age" : 87 }
        \]}\`)

First, we declare an empty interface and store it in a variable that can be passed to *json.Unmarshal* function as a pointer.

i := interface{}
err := json.Unmarshal(b, &in)
if err != nil {
        panic("OMG!!")
}
fmt.Print(i)

This is what you should see:

map\[Persons:\[map\[Name:John Age:35\] map\[Name:Jane Age:23\] map\[Age:12 Name:Mary\] map\[Name:Shane Age:35\] map\[Name:Phil Age:87\]\]\]

This is equivalent to assigning a map in this form:

i := map\[string\]interface{}{
        "Persons" : \[\]interface{}{
                map\[string\]interface{}{"Name":"John", "Age": 35},
                map\[string\]interface{}{"Name":"Jane", "Age": 23},
                map\[string\]interface{}{"Name":"Mary", "Age": 12},
                map\[string\]interface{}{"Name":"Shane", "Age":35},
                map\[string\]interface{}{"Name":"Phil", "Age": 87},
        },
}

Using interfaces allow you to deal with arbitrary data types before eventually converting them to appropriate ones.

Now, to convert an interface to an appropriate type, I use type assertions. Type assertion works like:

var in interface{}
in = 23
rightVal := in.(int)
wrongVal := in.(string)fmt.Println(rightVal)    // Correct type! This prints 23
fmt.Println(wrongVal)    // This will result in a panic!

I did expect an assertion to return a bool type (true or false), but I thought that wouldn’t make much sense considering returning false wouldn’t do one much good but keep guessing.

Now, to assert the type of our *i* interface, just use type assertion on the first level like so:

p := i.(map\[string\]interface{})
fmt.Println(p)  // This will prints exactly the same value as i,
                // but with concrete type

p is now a *map\[string\]interface{}* type, meaning we finally know (or Go does) that the first parent level is a map with key(s) of type string, and that’s it. The value(s) are still indeterminable as interface{}*,* which can be just about any type, even another interface.
Now, use type switch to dig into the structure and type assert the first level.

for k, v := range p {
        switch val := v.(type) {
        case string:
                fmt.Println(k, "is string", val)
        case int:
                fmt.Println(k, "is int", val)
        case \[\]interface{}:
                fmt.Println(k, "is an array")
                for i, v := range val {
                        fmt.Println(i, v)
                }
        default:
                fmt.Println(k, "is unknown type")
        }
}

Here we are just printing out to notify us of the assertion. If everything is right, this is what we should see:

Persons is an array:
0 map\[Name:John Age:35\]
1 map\[Name:Jane Age:23\]
2 map\[Name:Mary Age:12\]
3 map\[Name:Shane Age:35\]
4 map\[Name:Phil Age:87\]

Which is exactly this:

Persons := \[5\]map\[string\]interface{}

Or more verbosely:

Persons : \[
        map\[string\]interface{}{"Name": "John", "Age": 35},
        map\[string\]interface{}{"Name": "Jane", "Age": 23},
        ...
        map\[string\]interface{}{"Name": "Phil", "Age": 87},
\]

I noticed that the maps each contains both a string value and integer value. Go let that because those values are still indetermined interfaces. Therefore, “John” and 35 are not yet recognized as a string and integer until they’re type asserted.

My goal at this point was to convert the array of interfaces into a simple persons *map\[string\]int* like the first step (for the sake of my own education) so that I could create instances from it. Now I kept type asserting the next level:

// Type assertion on the value of key "Persons"
// which turns out to be an array of interface{}
n := m\["Persons"\].(\[\]interface{})

Try printing out *n*, I see the array under the key “Persons”.
Now that *n* is an array of interface{}’s, which I knew at this point that each member is of type *map\[string\]interface{}*, i.e. *\[{“Name”: “John”, “Age”:35}, …\]*, I just jumped into creating instances at this level by chaining type assertions all in one go.

// Declare a slice to be filled with \*Person's instances
persons := make(\[\]\*Person, len(n))

// Iterate through the n array (or slice)
for i := range n { // type assert on each member of n
        name := n\[i\].(map\[string\]interface{})\["Name"\].(string)
        age  := n\[i\].(map\[string\]interface{})\["Age"\].(float64) // Create an instance in every loop
        persons\[i\] = &Person{name, int(age)}
}
