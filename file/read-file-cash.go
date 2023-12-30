func main() {
 // Open a file
 file, err := os.Open("example.txt")
 if err != nil {
  // Handle the error
  fmt.Println("Error opening the file:", err)
  return
 }
 defer file.Close() // Close the file when done

 // Read from the file
 buffer := make([]byte, 1024)
 _, err = file.Read(buffer)
 if err != nil {
  // Handle the error
  fmt.Println("Error reading the file:", err)
  return
 }

 // Print the file content
 fmt.Println("File content:", string(buffer))
}