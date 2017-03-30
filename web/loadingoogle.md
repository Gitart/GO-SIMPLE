## Google Drive API upload and rename example


For this tutorial, we will explore how to upload file to Google Drive using RESTful method and rename the file with Drive API.

Before you start, please turn on the Google Drive API if you haven't do so and download the credential file 
- client_secret.json by following the Step 1: Turn on the Drive API section found in

https://developers.google.com/drive/v3/web/quickstart/go#prerequisites

Once you've downloaded the credential file, move it to the same directory as the source code below before running the program.

### NOTE:
You will need to change the file to be uploaded - img.png to something else that you have.

When prompted for the authorization code for the first time, cut-n-paste the URL from your terminal to your browser, 
then you will get a string(token), cut-n-paste that string into your terminal where you execute the program.

In case you get Authorization error in the API JSON reply, you will need to authorize the Drive API v3 scope 
at https://developers.google.com/oauthplayground/

Next, do a go get command
```
>go get google.golang.org/api/drive/v3
```

Finally, enable the Drive API by going to https://developers.google.com/drive/v3/web/enable-sdk and follow the  
instruction listed below To enable the Drive API, complete these steps:
Run this code example below and observe the changes to your own Google Drive content at https://drive.google.com/drive/my-drive


```golang
 package main

 import (
         "crypto/rand"
         "encoding/json"
         "fmt"
         "github.com/antonholmquist/jason"
         "golang.org/x/net/context"
         "golang.org/x/oauth2"
         "golang.org/x/oauth2/google"
         "google.golang.org/api/drive/v3"
         "io/ioutil"
         "log"
         "net/http"
         "net/url"
         "os"
         "os/user"
         "path/filepath"
         "strconv"
         "strings"
 )

 // NOTE : we don't want to visit CSRF URL to get the authorization code
 // and paste into the terminal each time we want to send an email
 // therefore we will retrieve a token for our client, save the token into a file
 // you will be prompted to visit a link in your browser for authorization code only ONCE
 // and subsequent execution of the program will not prompt you for authorization code again
 // until the token expires.

 // getClient uses a Context and Config to retrieve a Token
 // then generate a Client. It returns the generated Client.
 func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
         cacheFile, err := tokenCacheFile()
         if err != nil {
                 log.Fatalf("Unable to get path to cached credential file. %v", err)
         }
         tok, err := tokenFromFile(cacheFile)
         if err != nil {
                 tok = getTokenFromWeb(config)
                 saveToken(cacheFile, tok)
         }
         return config.Client(ctx, tok)
 }

 // getTokenFromWeb uses Config to request a Token.
 // It returns the retrieved Token.
 func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
         authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
         fmt.Printf("Go to the following link in your browser then type the "+
                 "authorization code: \n%v\n", authURL)

         var code string
         if _, err := fmt.Scan(&code); err != nil {
                 log.Fatalf("Unable to read authorization code %v", err)
         }

         tok, err := config.Exchange(oauth2.NoContext, code)
         if err != nil {
                 log.Fatalf("Unable to retrieve token from web %v", err)
         }
         return tok
 }

 // tokenCacheFile generates credential file path/filename.
 // It returns the generated credential path/filename.
 func tokenCacheFile() (string, error) {
         usr, err := user.Current()
         if err != nil {
                 return "", err
         }
         tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
         os.MkdirAll(tokenCacheDir, 0700)
         return filepath.Join(tokenCacheDir,
                 url.QueryEscape("google-drive-golang.json")), err
 }

 // tokenFromFile retrieves a Token from a given file path.
 // It returns the retrieved Token and any read error encountered.
 func tokenFromFile(file string) (*oauth2.Token, error) {
         f, err := os.Open(file)
         if err != nil {
                 return nil, err
         }
         t := &oauth2.Token{}
         err = json.NewDecoder(f).Decode(t)
         defer f.Close()
         return t, err
 }

 // saveToken uses a file path to create a file and store the
 // token in it.
 func saveToken(file string, token *oauth2.Token) {
         fmt.Printf("Saving credential file to: %s\n", file)
         f, err := os.Create(file)
         if err != nil {
                 log.Fatalf("Unable to cache oauth token: %v", err)
         }
         defer f.Close()
         json.NewEncoder(f).Encode(token)
 }

 func randStr(strSize int, randType string) string {

         var dictionary string

         if randType == "alphanum" {
                 dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "alpha" {
                 dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "number" {
                 dictionary = "0123456789"
         }

         var bytes = make([]byte, strSize)
         rand.Read(bytes)
         for k, v := range bytes {
                 bytes[k] = dictionary[v%byte(len(dictionary))]
         }
         return string(bytes)
 }

 func main() {

         ctx := context.Background()

         // process the credential file
         credential, err := ioutil.ReadFile("client_secret.json")
         if err != nil {
                 log.Fatalf("Unable to read client secret file: %v", err)
         }

         // In order for POST upload attachment to work
         // You need to authorize the Gmail API v1 scope
         // at https://developers.google.com/oauthplayground/
         // otherwise you will get Authorization error in the API JSON reply

         // Use DriveScope for this example. Because of we want to Manage the files in
         // Google Drive.

         // See the rest at https://godoc.org/google.golang.org/api/drive/v3#pkg-constants

         config, err := google.ConfigFromJSON(credential, drive.DriveScope)
         if err != nil {
                 log.Fatalf("Unable to parse client secret file to config: %v", err)
         }

         client := getClient(ctx, config)

         // initiate a new Google Drive service
         driveClientService, err := drive.New(client)
         if err != nil {
                 log.Fatalf("Unable to initiate new Drive client: %v", err)
         }

         // get our token
         cacheFile, err := tokenCacheFile()
         if err != nil {
                 log.Fatalf("Unable to get path to cached credential file. %v", err)
         }

         token, err := tokenFromFile(cacheFile)
         if err != nil {
                 log.Fatalf("Unable to get token from file. %v", err)
         }

         // we will use MULTIPART upload method
         // see https://developers.google.com/drive/v3/web/manage-uploads

         // read file for upload purpose
         fileName := "img.png" // <------------ CHANGE HERE!
         fileBytes, err := ioutil.ReadFile(fileName)
         if err != nil {
                 log.Fatalf("Unable to read file for upload: %v", err)
         }

         fileMIMEType := http.DetectContentType(fileBytes)

         // use Multipart upload method because we need both the media and its metadata(filename)
         // Simple upload method will cause the filename to be named "Untitled"

         postURL := "https://www.googleapis.com/upload/drive/v3/files?uploadType=multipart"

         // extract auth or access token from Token file
         // see https://godoc.org/golang.org/x/oauth2#Token
         authToken := token.AccessToken

         boundary := randStr(32, "alphanum")

         uploadData := []byte("\n" +
                 "--" + boundary + "\n" +
                 "Content-Type: application/json; charset=" + string('"') + "UTF-8" + string('"') + "\n\n" +
                 "{ \n" +
                 string('"') + "name" + string('"') + ":" + string('"') + fileName + string('"') + "\n" +
                 "} \n\n" +
                 "--" + boundary + "\n" +
                 "Content-Type:" + fileMIMEType + "\n\n" +
                 string(fileBytes) + "\n" +

                 "--" + boundary + "--")

         // post to Drive with RESTful method
         request, _ := http.NewRequest("POST", postURL, strings.NewReader(string(uploadData)))
         request.Header.Add("Host", "www.googleapis.com")
         request.Header.Add("Authorization", "Bearer "+authToken)
         request.Header.Add("Content-Type", "multipart/related; boundary="+string('"')+boundary+string('"'))
         request.Header.Add("Content-Length", strconv.FormatInt(request.ContentLength, 10))

         // debug
         //fmt.Println(request)

         response, err := client.Do(request)
         if err != nil {
                 log.Fatalf("Unable to be post to Google API: %v", err)
         }

         defer response.Body.Close()
         body, err := ioutil.ReadAll(response.Body)

         if err != nil {
                 log.Fatalf("Unable to read Google API response: %v", err)
         }

         // output the response from Drive API
         fmt.Println(string(body))

         // we need to extract the uploaded file ID to execute Update command
         jsonAPIreply, _ := jason.NewObjectFromBytes(body)

         uploadedFileID, _ := jsonAPIreply.GetString("id")
         fmt.Println("Uploaded file ID : ", uploadedFileID)

         // -------------------------------------------------------------------------------
         // just for fun, let's rename our uploaded file with Google Drive API

         // see https://godoc.org/google.golang.org/api/drive/v3#File
         renamedFile := drive.File{Name: "renamed-" + fileName}

         _, err = driveClientService.Files.Update(uploadedFileID, &renamedFile).Do()

         if err != nil {
                 log.Fatalf("Unable to rename(update) uploaded file in Drive:  %v", err)
         } else {
                 fmt.Println("Renamed " + fileName + " to " + renamedFile.Name)
         }

         // list out the uploaded files
         // see https://godoc.org/google.golang.org/api/drive/v3#FilesService.List
         filesListCall, err := driveClientService.Files.List().OrderBy("name").Do()

         if err != nil {
                 log.Fatalf("Unable to list files in Drive:  %v", err)
         }
         // NOTE : Will list out the files inside the Trash folder as well

         fmt.Println("Files:")
         if len(filesListCall.Files) > 0 {
                 for num, file := range filesListCall.Files {
                         fmt.Printf("[%d]: %s (%s) Trashed : %v\n", num, file.Name, file.Id, file.Trashed)
                 }
         } else {
                 fmt.Println("No files found.")
         }

 }
 ```
 
 
Sample output:

```
Go to the following link in your browser then type the authorization code: https://accounts.google.com/o/oauth2/auth?accesstype=offline&clientid=[redacted]&state=state-token
4/sh-[redacted]-8Q
Saving credential file to: /Users/admin/.credentials/google-drive-golang.json

```json
{
"kind": "drive#file",
"id": "0B_wYQdperIb5Vk5vaDFjdmpibE0",
"name": "img.jpg",
"mimeType": "image/jpeg"
}
```

Uploaded file ID : 0B_wYQdperIb5Vk5vaDFjdmpibE0
Renamed img.jpg to renamed-img.jpg
Files:
[0]: renamed-img.jpg (0B_wYQdperIb5Vk5vaDFjdmpibE0) Trashed : false
