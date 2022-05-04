package main

import (
	"fmt"
)

func main() {

type Y struct{
         Gh   string
         Name string
         Tags []string
}

t:=[]struct{
            Notify    string 
            Records   string
            Structure string
            Other     string
            Tg        Y 
           }{
              {"Kiev",     "Kievskaya",  "Структура 1", "Address 1",  Y{"Jonh",   "Kolobov", []string{"Arts",  "Traveling"}} }, 
              {"Kotlova",  "Obolon",     "Structur 2",  "Address 2",  Y{"Inecca", "Deeva",   []string{"Photos", "Reading", "Other Hoobyt"}}},
              {"август",   "Pechorsk",   "Structur 3",  "Address 3",  Y{"Inecca", "Deeva",   []string{"Photos", "Reading", "Other Hoobyt","Flyting", "Productive"}}},
              {"Киев",     "одинадцять", "Structur 4",  "Address 4",  Y{"Inecca", "Deeva",   []string{"Photos", "Reading", "Other Hoobyt","Flyting", "Productive"}}},
           }


          // Добавление тегов в первый элемент
          fmt.Println("Other [1] : ", t[1].Other)
          t[1].Tg.Tags = append(t[1].Tg.Tags, "Обновление Системный тег для первого")
          fmt.Println("Other [1] : ", t[1])
          addsrtring := []string{"Составный тег", "Ghrtytytytrytr0304435","sdfdsffdsfdsfdrty", "Норвегия","Бургундия"}



// Обход структуры	
for _, tt:=range(t){
     fmt.Println(tt.Structure, tt.Other, tt.Tg.Name, tt.Tg.Tags[0], len(tt.Tg.Tags))
         
         // Добавление тега в массив
         tt.Tg.Tags = append(tt.Tg.Tags, "Добавленный тег")
         tt.Tg.Tags = append(tt.Tg.Tags, "Составный тег", "Ghrtytytytrytr0304435","sdfdsffdsfdsfdrty")
         tt.Tg.Tags = append(tt.Tg.Tags, "Новый тег ")
         tt.Tg.Tags = append(tt.Tg.Tags, "Интеграционный тег ")
         tt.Tg.Tags = append(tt.Tg.Tags, "Компаудный тег ")
         tt.Tg.Tags = append(tt.Tg.Tags, "Системный тег ")
         
         // Прокурутка Tags
         for _, iop:=range(tt.Tg.Tags){
             fmt.Println("Tags :", iop)
               
         }

       fmt.Println( "\n")
 }
}

// © 2022 GitHub, In
