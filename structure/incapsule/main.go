package main

import "fmt"
import "encoding/json"

var (
      cashSetting = Sett{}
)


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

func (b Basic) Long() string {
      return b.Title
}

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

func (t Terra) Titles() string {
      return "TITLE TERRA DATA :" + t.Title
}

// func (t Terra) Dats() Solo {
//       return  t.Date
// }

// Terra show json
func (m Terra) ShowJSONStatistics() {
      data, _ := json.MarshalIndent(m, "", "  ")
      fmt.Printf("Terra data : %s\n", string(data))
}


// SOLO
type Solo struct {
     Title string
     Date  string
}

// main
func main(){
     var Dat Sett
     
     Dat.Title                                   = "Seting your system"
     Dat.System.Title                            = "System: Title"
     Dat.Basic.Title                             = "Bisaik title get"
     Dat.Basic.Incapsule.Date                    = "Дата в инкапсуле"
     Dat.Basic.Incapsule.Title                   = "Title в инкапсуле"
     Dat.Basic.Incapsule.Subcapsule.Title        = "SUBCAPSULA: TItle в инкапсуле"
     Dat.Basic.Incapsule.Subcapsule.Terra.Title  = "TerraDATA:  TItle в инкапсуле"

     // Cash setting
     cashSetting.System.Title                    = "Cash: sysytem title"
     cashSetting.Basic.Incapsule.Date            = "Cash Date setting"

     fmt.Println(Dat)
     fmt.Println(Dat.Basic.Long())
     fmt.Println(Dat.Basic.Title)

     fmt.Println(Dat.Basic.Incs().Duble())
     fmt.Println(Dat.Basic.Incs().Rsub())
     fmt.Println(Dat.Basic.Incs().Rsub().Teradata())
     fmt.Println("TITLE DATA:",    Dat.Basic.Incs().Rsub().Teradata().Title)
     fmt.Println("FUNCTION DATA:", Dat.Basic.Incs().Rsub().Teradata().Titles())

     Dat.Basic.Incs().Rsub().Teradata().ShowJSONStatistics()

     fmt.Println(cashSetting.Basic.Incapsule.Date)
     

}
