# Responding to Requests

I wanted my GoViolin app to be able to respond to a request for a scale and display an image of the scale and also allow the
user to play a .mp3 file. This meant that the web app was going to need to be able to hear a request and respond to it. I wanted my web app to:

Generate some radio buttons.  
Display them on a html page.   
Allow the user to make a selection.   
Respond based on the user’s selection.  

For steps 1 and 2, generating the radio buttons and displaying them, I was able to accomplish this using a template and to generate a series of radio buttons using values stored in a struct.
For step 3, when the user makes a selection, I wanted to send the user’s selection in a POST request to the back-end. I was able to do this by creating a form and using jquery to submit the form when a radio button is changed.
I have written some code to demonstrate how a web app can respond to a request. This example uses two files, select.html and select.go which are shown below. When this code is run and the user navigates to localhost:8080, the user is asked if they prefer cats or dogs and presented with two radio buttons. When the user makes a selection, the page updates and their answer is displayed.

```html
<!DOCTYPE html>
<html>

<head>
<script type='text/javascript' src='https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js'></script>
<title>{{.PageTitle}}</title>
</head>

<script type='text/javascript'>
 $(document).ready(function() {
   $('input[name=animalselect]').change(function(){
     $('form').submit();
   });
});
</script>
<body>
  {{with $1:=.PageRadioButtons}}
  <p> Which do you prefer</p>

    <form action="/selected" method="post">
         {{range $1}}
           <input type="radio" name={{.Name}} value={{.Value}} {{if .IsDisabled}} disabled=true {{end}} {{if .IsChecked}}checked{{end}}> {{.Text}}
         {{end}}
    </form>
  {{end}}

  {{with $2:=.Answer}}
    <p>Your answer is {{$2}}</p>
  {{end}}

</body>
</html>
```

view rawselect.html hosted with ❤ by GitHub

```golang
package main

import (
  "net/http"
  "log"
  "html/template"
)


type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
  PageTitle        string
  PageRadioButtons []RadioButton
  Answer           string
}


func main() {
  http.HandleFunc("/", DisplayRadioButtons)
  http.HandleFunc("/selected", UserSelected)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){
 // Display some radio buttons to the user

   Title := "Which do you prefer?"
   MyRadioButtons := []RadioButton{
     RadioButton{"animalselect", "cats", false, false, "Cats"},
     RadioButton{"animalselect", "dogs", false, false, "Dogs"},
   }

  MyPageVariables := PageVariables{
    PageTitle: Title,
    PageRadioButtons : MyRadioButtons,
    }

   t, err := template.ParseFiles("select.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
   }

}

func UserSelected(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  // r.Form is now either
  // map[animalselect:[cats]] OR
  // map[animalselect:[dogs]]
 // so get the animal which has been selected
  youranimal := r.Form.Get("animalselect")

  Title := "Your preferred animal"
  MyPageVariables := PageVariables{
    PageTitle: Title,
    Answer : youranimal,
    }

 // generate page by passing page variables into template
    t, err := template.ParseFiles("select.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}
```
