### Parsing the set of links from a text file
Assuming the text file has a set of links each on a new line, the below function will read the text file and return a slice of all links.

```golang
func readLines(path string) ([]string, error) {
   file, err := os.Open(path)
   if err != nil {
       return nil, err
   }
   defer file.Close()

   var lines []string
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
       lines = append(lines, scanner.Text())
   }
   return lines, scanner.Err()
}
```

Iterate over each link to generate broken links for each parsed link
Following code will iterate over each link to generate reports for each site having broken links. The logic for crawling links and detecting broken links will go in the for loop and will be illustrated in the next coming steps.

```golang
linksToCrawl,err := readLines("links.txt")

if err != nil {
   fmt.Println("Error reading file: ", err)
}

// crawl each website in input file one consecutively
for i := 0; i < len(linksToCrawl); i++ {
      //crawling logic for each link
}
```

Crawl over the web page and inbound the slice of found links in the channel
The most important part starts here when there are a lot of stuff you should know about. The terms going to be used here are Channels/Buffered Channels, Tokenizers, Wait Group and Goroutines. All terms except tokenizers are going to be used for crawling the links with handled concurrency. Tokenization is done by creating a Tokenizer for an io.Reader r. It is the caller’s responsibility to ensure that r provides UTF-8 encoded HTML.

The buffer size of channels is not going to matter much but I have still kept it 100 so that there are not many links in the channel and thus leading in unmanageable goroutines. My program crashed many times crawling one huge website because there were 1000’s of goroutines being executed concurrently.

Before any more explanation, I need to also specify the struct created for each link. Here it is as follows: -

```golang
// every link found stored as a node
type Node struct {
   link string
   parent string
   linkText string
   isOutsideLink bool
   statusCode int
}
```

Thus, the below code snippet uses number of variables and a function named getLinks(). Let me explain you the use of those variables and the function.
Urls: — It is a buffered channel of slice of type Node. The reason being slice of type Node is that whenever a set of links for a particular link is found, the set of links is directly inserted in a new slice links and further inbounded in the channel.
getLinks(resp,root,Urls): — The getLinks function takes the http.Response, link to get links for, channel to inbound the set of links.
isCrawled[root.link]: — This variable is a Node-bool map in which it keeps the track of crawled links.
counter: — This variables is a counter for links to be crawled in the Urls chan.

```golang
//kept depth as 4 for the depth-first-search
func crawl(root Node,depth int){
   Urls:=make(chan []Node,100) // buffer size 100
   resp, err := http.Get(string(root.link)) //get the html content
   if err != nil {
      log.Panic("error is ", err)
   }
   go getLinks(resp,root,Urls) //find links for root
   defer resp.Body.Close()
   counter :=1  //counter for links to be crawled in the Urls chan
   isCrawled[root.link] = true //isCrawled is
   for i:=0;i<depth;i++{
      for counter > 0{
         counter --
         next:= <- Urls  //inbound the found links
         //further logic for iterating over the next slice for broken links --> check step 4
      }
   }
   return
}
```

Lets explore the getLinks function now having html package’s Tokenizer, net/url package’s ResolveReference to fix the relative paths to absolute paths because the href tag can have any value such as “//”, “../”,”../../”,”./” etc, and Wait groups. The net/url package’s function ResolveReference resolves such relatives paths to absolute paths.
The golang’s package sync has a type WaitGroup where each WaitGroup has a set of goroutines and it waits until all the goroutines are finished executing.

```golang
func getLinks(resp *http.Response, parent Node,Urls chan []Node) {
   wg.Add(1) //add link to waitGroup
   var links = make([]Node,0) //to get a set of links
   z := html.NewTokenizer(resp.Body) // a new tokenizer for the response of the html
   for {
      tt := z.Next()
      switch {
      case tt == html.ErrorToken:
         Urls <- links    
         wg.Done()
         return
      case tt == html.StartTagToken:
         t := z.Token()            // taken token
         isAnchor := t.Data == "a" // checking whether it is anchor
         
         if isAnchor{
            for _, a := range t.Attr { //going through all attributes of anchor i.e., t
               base,err := url.Parse(websiteURL1)
               if err != nil {
                  log.Println("err is ",err)
               }
               z.Next()
               t1 := z.Token()
               
               if a.Key == "href"{
                  a.Val = strings.TrimSpace(a.Val)

                  if a.Val != parent.link {
                     var newNode Node //create a node for the found link
                     if !strings.Contains(a.Val,"tel:") && !strings.Contains(a.Val,"mailto:"){
                        //parsing the url
                        u, err := url.Parse(a.Val)
                        if err != nil {
                           log.Println(err)
                        }else {
                           uri := base.ResolveReference(u)
                           a.Val = uri.String()
                           //checking outsideLink
                           if strings.Contains(a.Val, websiteURL1) {
                              newNode = Node{a.Val, parent.link,linkText,false, 0}
                           } else {
                              newNode = Node{a.Val, parent.link,linkText,true, 0}
                           }
                        }

                     }

                     if !strings.Contains(newNode.link,"mailto:") && !isCrawled[newNode.link]{
                        links = append(links, newNode)
                     }
                     break
                  }
               }
            }
         }
      }
   }
   wg.Done()
   return
}
```
Check the http response of each link in the inbounded slice of links.

```golang
next:= <- Urls
for _, url := range next {
   if _, done := isCrawled[url.link] ; !done {
      if url.link!=""{
         timeout := time.Duration(10 * time.Second)
         client := http.Client{Timeout: timeout}
         resp, err := client.Get(strings.TrimSpace(url.link)) 
         if resp != nil {
            if resp.StatusCode == 404 {
               brokenLinks = append(brokenLinks, url) //appended to broken links slice
            }else {
               if !url.isOutsideLink && url.linkText=="Link" {
                  counter ++
                  go getLinks(resp, url, Urls) //crawl the remaining links in chan
                  isCrawled[url.link] = true
               }
            }
         }
      }
   }
}
```

Once the slice of nodes are inbounded in the channel, based on the depth specified, each link in the slice will be checked for the http response. Based on the http response, if it is 404 then it is tagged as a broken link. If the http response is not 404 then the counter is incremented and a goroutine is started to get links for that particular link. Ofcourse, after the link is crawled, it is flagged as crawled so that no further crawling is done on the same link again.

Send a report of all broken links via email
The report for broken links is sent via email through iterating over broken links slice.

See the below: -

```golang
for i := 0; i < len(brokenLinks); i++ {
   if brokenLinks[i].link != ""{
      fmt.Printf("Broken link : %s \n", brokenLinks[i].link)
      contentForMail = append (contentForMail,"Link Type: "+brokenLinks[i].linkText+"<br>Link URL: <a href='"+brokenLinks[i].link+"'>"+brokenLinks[i].link+"</a>" + "<br>Source: " + brokenLinks[i].parent+"<br><br>")
      results = append(results, "Link Type: "+brokenLinks[i].linkText+"<br>Link URL: <a href='"+brokenLinks[i].link+"'>"+brokenLinks[i].link+"</a>" + "<br>Source: <a href='"+brokenLinks[i].parent+"'>" + brokenLinks[i].parent+"</a><br><br>")
   }
}
```

Thus this particular results variable is passed to sendMail function to send mail to the reporter.

```golang
func sendMail(result string){
   m := gomail.NewMessage()
   m.SetHeader("From","*sender email *")
   m.SetHeader("To", "*receiver email*")
   m.SetHeader("Subject", "["+websiteURL1+"] Broken Links Detected")
   m.SetBody("text/html", result)
   d := gomail.NewDialer("smtp.gmail.com", 25, "*sender email*", "*sender password*")

   // Send the email to Bob, Cora and Dan.
   if err := d.DialAndSend(m); err != nil {
      log.Panic(err)
   }
}
```
