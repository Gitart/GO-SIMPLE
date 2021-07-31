 package main 

import (
	// "os"
    "os/user"
    "fmt"
 )


 func ChekUser(){
 	user, err := user.Current()
    if err != nil {
        panic(err)
    }

    // Current User
    fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
    fmt.Println("Username: " + user.Username)
    fmt.Println("Home Dir: " + user.HomeDir)

    // Get "Real" User under sudo.
    // More Info: https://stackoverflow.com/q/29733575/402585
    // fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
 }
