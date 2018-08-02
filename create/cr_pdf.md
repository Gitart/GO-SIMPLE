# Create PDF file from HTML file
golang pdf html 

When I was coding with PHP and CodeIgniter framework, I need to generate PDF file to email receipts to my customers. In CodeIgniter/PHP world, there are multiple third party library - such as DOMPDF - that can render and convert a HTML file to PDF file with ease.
In Golang world, one can try to utilise the GoPDF packages (https://github.com/signintech/gopdf) or another different GoPDF package (bitbucket.org/zombiezen/gopdf/pdf) . However, I found them complicated for my own use and unable to render HTML codes.
In this example, I will use os.Exec() function to render HTML file and generate PDF file with domPDF.
First, you need to download the domPDF package from https://github.com/dompdf/dompdf/downloads, unzip, set the right permission to execute dompdf.php file.
Second, you need to have PHP installed on the server you are executing this code. See https://code.google.com/p/dompdf/wiki/Installation on how to get domPHP installed properly.
Here you go :

```go
 package main

 import (
         "bytes"
         "fmt"
         "io/ioutil"
         "net/http"
         "os"
         "os/exec"
         "path/filepath"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte(fmt.Sprintf("<html><body><form action='http://localhost:8080/pdf' method='post'><input type='submit' value='Generate PDF'></form></body></html>")))
 }

 func PDF(w http.ResponseWriter, r *http.Request) {

         // NOTE : receipt.html is a template file
         //        for the sake of this tutorial, we are keeping things simple

         // use os.Exec to execute domPDF
         // read
         // https://code.google.com/p/dompdf/wiki/Usage#Invoking_dompdf_via_the_command_line
         // on how to configure and use domPDF from command line

         // change the sweetlogic path to yours !!

         // without filepath.Abs...it won't work! why???!?!
         arg1, _ := filepath.Abs("/Users/sweetlogic/dompdf/dompdf.php")
         arg2, _ := filepath.Abs("/Users/sweetlogic/receipt.html")



         cmd := exec.Command("php", arg1, arg2)
         out, err := cmd.Output()

         if err != nil {
                 fmt.Println(string(out))
                 fmt.Println(err)
                 return
         }

         fmt.Print(string(out))

         // NOTE : In real world application, the output filename should be dynamic
         // and not static like in this example. For the sake of simplicity, we just
         // keep the output name to receipt.pdf
         // to change the output filename, see the domPDF command line instruction.

         // grab the generated receipt.pdf file and stream it to browser
         streamPDFbytes, err := ioutil.ReadFile("./receipt.pdf")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         b := bytes.NewBuffer(streamPDFbytes)

         // stream straight to client(browser)
         w.Header().Set("Content-type", "application/pdf")

         if _, err := b.WriteTo(w); err != nil {
                 fmt.Fprintf(w, "%s", err)
         }

         w.Write([]byte("PDF Generated"))
 }

 func main() {
         // http.Handler
         mux := http.NewServeMux()
         mux.HandleFunc("/", Home)
         mux.HandleFunc("/pdf", PDF)

         http.ListenAndServe(":8080", mux)
 }
 ```
 
### the content of the receipt.html file -- ripped off from one of my old project.

```html
 <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
 <html xmlns="http://www.w3.org/1999/xhtml">
 <head>
 <title><?php echo $title;?></title>
 <meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
 <style type="text/css">
 /**
  * NOTE : This file NEEDS a locally located stylesheet to generate the appropriate formatting for
  * transformation into a PDF.  If you alter this file (and you are encouraged to do so) just
  * keep in mind that all of your formatting must be located here.  You might also find that
  * there is limited or no support for a specific CSS style you want (ie: floating) and you'll
  * need to work around with old-school tables.  Sorry for that... ;)
  *
  *
  * NOTE : Things have changed since I used domPDF with PHP. See latest updates on how
  * domPDF renders CSS at https://github.com/zimmski/dompdf
  *
  * - SocketLoop.com - Adam
  */


 body {
     width: 100%;
     margin: 0.5in;
 }

 h1, h2, h3, h4, h5, h6, li, blockquote, p, th, td {
     font-family: Helvetica, Arial, Verdana, sans-serif; /*Trebuchet MS,*/
 }
 h1, h2, h3 {
     color: #5E88B6;
     font-weight: normal;
 }
 h4, h5, h6 {
     color: #000;
     font-size:12px;
 }
 h2 {
     margin: 0 auto auto auto;
     font-size: x-large;
 }
 h2 span {
     text-transform: uppercase;
 }

 li, blockquote, p, th, td {
     font-size: 80%;
 }
 ul {
     list-style: url('../images/system/bullet.gif') none;
 }
 table {
     width: 100%;
 }

 td {
     padding: 7px 15px;
 }
 td p {
     font-size: small;
     margin: 0;
 }

 td h1 {
     font-size:36px;
     font-weight:bold;
     margin: 0;
 }

 thead, tfoot {    background-color:black;color:white}

 th {
         color: #FFF;
         text-align: left;
         background-color:#000000;

 }
 .invoice_words {
     color: #000;
     font-size:22px;
     font-weight: bold;
     text-align: left;
     text-transform:uppercase;
 }
 .invoice_total {
     color: #000;
     font-size:12px;
     font-weight: bold;
     text-align: right;
     text-transform:uppercase;
 }

 #footer {
     border-top: 1px solid #CCC;
     text-align: right;
     font-size: 16px;
     color: #999999;
 }
 #footer a {
     color: #999999;
     text-decoration: none;
 }
 table.stripe {
     border-collapse: collapse;
     page-break-after: auto;
 }
 table.stripe td {
     border-bottom: 1pt solid black;
 }
 </style>
 </head>
 <body>

 <table>
         <tr>
             <td width="60%">
             <h1>
                 <img src="https://d1ohg4ss876yi2.cloudfront.net/preview/golang.png"/>
                 Receipt<br />
             </h1>
             <br />
             </td>
             <td>
                 <h2>
                     [company name]<br />
                 </h2>
                 <p>
                     [your street address]<br/>
                     Email: [your email address]<br />
                 </p>
             </td>
         </tr>
     </table>

     <h3> You can add more data here or modify the CSS.... use your imagination! </h3>

 </body>
 </html>
 ```


now, run the code and point your browser to http://localhost:8080 and you should see a button to generate PDF. Click on the button and if everything goes smoothly, you should see a PDF file being downloaded.
Hope this tutorial is useful to you. Not entirely a 'pure' Golang tutorial because it relies on PHP packages.....but hey... it works.
NOTE : This is the bitbucket.org/zombiezen/gopdf/pdf code example I that I've tried, but not useful for me. Maybe it can be useful to you

```
 package main

 import (
         "bitbucket.org/zombiezen/gopdf/pdf"
         "fmt"
         "net/http"
         "os"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte(fmt.Sprintf("<html><body><form action='http://localhost:8080/pdf' method='post'><input type='submit' value='Generate PDF'></form></body></html>")))
 }

 func PDF(w http.ResponseWriter, r *http.Request) {
         doc := pdf.New()
         canvas := doc.NewPage(pdf.A4Width, pdf.A4Height)

         canvas.Translate(100, 100)

         path := new(pdf.Path)
         path.Move(pdf.Point{0, 0})
         path.Line(pdf.Point{100, 0})
         canvas.Stroke(path)

         text := new(pdf.Text)
         text.SetFont(pdf.Helvetica, 14)
         text.Text("Hello, World!")
         canvas.DrawText(text)


         text2 := new(pdf.Text)
         //text.SetFont(pdf.Helvetica, 14)

         html := "<html><body><h1>Receipt #0001</h2><br><br><h2>Issue to You!</h2></body></html>"

         text2.Text(html) // will not render ya
         canvas.DrawText(text2)

         canvas.Close()

         // download straight to browser
         w.Header().Set("Content-type", "application/pdf")
         err := doc.Encode(w)
         if err != nil {
                  fmt.Println(err)
                 os.Exit(1)
         }

         // save a copy on web server
         pdfFile, err := os.Create("x.pdf")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         doc.Encode(pdfFile)

         fmt.Println("Save to x.pdf...")

         w.Write([]byte("PDF Generated"))
 }

 func main() {
         // http.Handler
         mux := http.NewServeMux()
         mux.HandleFunc("/", Home)
         mux.HandleFunc("/pdf", PDF)

         http.ListenAndServe(":8080", mux)
 }
 ```
 
