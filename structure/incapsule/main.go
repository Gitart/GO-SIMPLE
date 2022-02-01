package main

import "fmt"

// SETT
type Sett struct {
       Title      string
       System     System
       Total      Total
       Basic      Basic
}


type System struct {
       Title string
       Date  string
}

type Total struct {
       Title string
       Date  string
}

// BASIC
type Basic struct {
     Title      string
     Date       string
     Incapsule     
}

func (b Basic) Long() string {return b.Title}
func (b Basic) Incs() Incapsule {return b.Incapsule}


// INCAPSULE
type Incapsule struct {
     Title string
     Date  string
     Subcapsule
}

func (b Incapsule) Rsub() Subcapsule { return b.Subcapsule}
func (b Incapsule) Duble() string {return "===> Duble Incapsule" + b.Title}


// SUBCAPSULE
type Subcapsule struct {
     Title string
     Date  string
     Terra
}

func (s Subcapsule) Teradata() Terra {return s.Terra}


// TERRAA
type Terra struct {
     Title string
     Date  string
}

func (t Terra) Titles() string {return "TITLE TERRA DATA :" + t.Title}


// main
func main(){
     var Dat Sett
     
     Dat.Title                                   = "Seting your system"
     Dat.System.Title                            = "System: Title"
     Dat.Basic.Title                             = "Bisaik title get"
     Dat.Basic.Incapsule.Date                    = "Дата в инкапсуле"
     Dat.Basic.Incapsule.Title                   = "TItle в инкапсуле"
     Dat.Basic.Incapsule.Subcapsule.Title        = "SUBCAPSULA: TItle в инкапсуле"
     Dat.Basic.Incapsule.Subcapsule.Terra.Title  = "TerraDATAA: TItle в инкапсуле"

     fmt.Println(Dat)
     fmt.Println(Dat.Basic.Long())
     fmt.Println(Dat.Basic.Title)


     fmt.Println(Dat.Basic.Incs().Duble())
     fmt.Println(Dat.Basic.Incs().Rsub())
     fmt.Println(Dat.Basic.Incs().Rsub().Teradata())
     fmt.Println("TITLE DATA:",   Dat.Basic.Incs().Rsub().Teradata().Title)
     fmt.Println("FUNTION DATA:", Dat.Basic.Incs().Rsub().Teradata().Titles())
}
