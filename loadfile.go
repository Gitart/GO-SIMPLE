Entefunc webUploadHandler(w http.ResponseWriter, r *http.Request) {

 file, header, err := r.FormFile("file") // the FormFile function takes in the POST input id file

 if err != nil {
    fmt.Fprintln(w, err)
    return
 }
 defer file.Close()

 // My error comes here

 messageId := r.URL.Query()["id"][0]
 out, err := os.Create("./upload/" + messageId + ".mp3")

 if err != nil {
    fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
    return
 }
 defer out.Close()

 // write the content from POST to the file
 _, err = io.Copy(out, file)
 if err != nil {
    fmt.Fprintln(w, err)
 }

 fmt.Fprintf(w,"File uploaded successfully : ")
 fmt.Fprintf(w, header.Filename)

}

