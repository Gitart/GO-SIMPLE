 
# ⚒ Auto-generate reply email with text/template package


**Problem :**
You have customers that ask question via email that sometimes bordering repetitive questioning or simply refused to read the Frequently Asked Question section. What should you do?

**Solution :**
Instead of cracking your head to write reply email each time, why not create an email auto-generator? It can be done easily with the text/template package. For this example, I will keep it simple. No artificial intelligence stuff yet( I'm working on it ) and auto email out the reply yet(shouldn't be hard to implement)

```go
 package main

 import (
 	"fmt"
 	"log"
 	"os"
 	"text/template"
 )

 func main() {
 	// Define a email template.
 	const email = `
 Hey {{.Name}},
 {{if .Attended}}
 Have you read FAQ section yet?{{else}}
 It is a shame that you choose not to understand my first email reply!{{end}}
 {{with .Gift}}Thank you for your equiry about {{.}} by the way.
 {{end}}
 Best wishes,
 Gossie
 `

 	// Prepare some data to insert into the template.
 	type Recipient struct {
 		Name, Gift string
 		Attended   bool
 	}
 	var recipients = []Recipient{
 		{"Aunt Fonda", "dragon bone pottery set", true},
 		{"Scott", "ancient map to the castle in the sky", false},
 		{"Jessie", "", false},
 	}

 	// Create a new template and parse the letter into it.
 	t := template.Must(template.New("email").Parse(email))

 	// Execute the template for each recipient.
 	for _, r := range recipients {
 		err := t.Execute(os.Stdout, r)
 		fmt.Println("----------------")
 		if err != nil {
 			log.Println("executing template:", err)
 		}
 	}

 }
 
 ```
 
 [Samples](https://www.socketloop.com/tutorials/golang-auto-generate-reply-email-with-text-template-package/?utm_source=socketloop&utm_medium=tutesidebar)
 
