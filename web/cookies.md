## Drop cookie to visitor's browser and http.SetCookie() example

There are times developers need to drop cookie to website visitors' browser to determine if a particular    
visitor has been to the website previously. Cookie can be used to customize how your website behave to    
a first time visitor or repeat visitor. One such example is the login page usage. Cookie will help your    
website to determine if the visitor is a first time visitor or not.   
For this tutorial, we will learn how to use http.SetCookie() function to drop cookie to website visitor.   
Here is an example code how to change the website message depending on the timestamp inside the visitor's cookie.    


```golang
 package main

 import (
         "net/http"
         "strconv"
         "time"
 )

 func home(w http.ResponseWriter, r *http.Request) {

         // setup cookie for deployment
         // see http://golang.org/pkg/net/http/#Request.Cookie

         // we will try to drop the cookie, if there's error
         // this means that the same cookie has been dropped
         // previously and display different message
         c, err := r.Cookie("timevisited") //

         expire := time.Now().AddDate(0, 0, 1)

         cookieMonster := &http.Cookie{
                 Name:  "timevisited",
                 Expires: expire,
                 Value: strconv.FormatInt(time.Now().Unix(), 10),
         }

         // http://golang.org/pkg/net/http/#SetCookie
         // add Set-Cookie header
         http.SetCookie(w, cookieMonster)

         if err != nil {
                 w.Write([]byte("Welcome! first time visitor!"))
         } else {
                 lasttime, _ := strconv.ParseInt(c.Value, 10, 0)
                 html := "Hey! Hello again!, your last visit was at "
                 html = html + time.Unix(lasttime, 0).Format("15:04:05")
                 w.Write([]byte(html))
         }
 }

 func main() {
         http.HandleFunc("/", home)
         http.ListenAndServe(":8080", nil)
 }
 ```
 
run the code above and point your browser to http://localhost:8080 and you should see this message for the first time :

Welcome! first time visitor!
refresh the page again, and this time you will see this message :

Hey! Hello again!, your last visit was at 15:16:02
If you are using Google Chrome browser, go to Views->Developer->Developer Tools and under the Resources tab, you should see Cookies->Localhost.
Delete the cookie named timevisited and refresh the page again. This time you should be able to see the original message again :

Welcome! first time visitor!
Hope this simple tutorial can be useful to you and happy coding!

### References :
http://en.wikipedia.org/wiki/HTTP_cookie
http://golang.org/pkg/net/http/#Cookie
