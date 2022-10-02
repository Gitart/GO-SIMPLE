@echo off

echo  //This code generation                   >  m.go
echo  //Dont touch this code handles           >> m.go
echo  //Generation file automaticccally        >> m.go
echo  package main                             >> m.go
echo  import "fmt"                             >> m.go
echo  func %1(){fmt.Println("ok")}             >> m.go
echo  func Info(){fmt.Println("ok")}           >> m.go


