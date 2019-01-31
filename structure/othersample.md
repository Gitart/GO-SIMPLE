
// https://www.callicoder.com/golang-maps/
func CretaeTab (){
     tb := make(map[int]string,5)
     tb[1]="sss"
     fmt.Println(tb[1])

     ti := make(map[string]interface{},5)
     ti["Norm"]="Interface"
     ti["Dat"]=1
     ti["Iye"]=Mst{"var":"dsddd"}
     fmt.Println(ti["Norm"], ti["Dat"], ti["Iye"].(Mst)["var"]  )


     t2 := make(map[string]interface{},5)
     t2["Norm"]="Interface"
     fmt.Println(t2)


     t3:=make(Mst,5)
     t3["d1"]="ssss"
     t3["d2"]="ssss"
     t3["d3"]="ssss"
     t3["d4"]="ssss"
     t3["d5"]="ssss"
     t3["d6"]=1
     fmt.Println(t3)

     var Tl [5]Mst 
     Tl[1]=Mst{"sss":"ssss"}
     Tl[2]=Mst{"sss":"ssss", "Norm":"Tets Norm", "Nom":122.23}
     Tl[3]=Mst{"Name":"hhh"}
     fmt.Println(Tl[2]["sss"])  
     fmt.Println(Tl[3])  
     fmt.Println(Tl[4])  

     t7:=make([]Mst,5)
     t7[1]=Mst{"bb":"sss"}                // Обязательно добавить первый элемент через структуру
     t7[1]["Ter"]="Global position"        
     t7[1]["Terd"]="News papaer"
     fmt.Println("T7:=",t7)  

     var t8 [10]Mst
     t8[0] = Mst{"bb":"sss"}
     fmt.Println("T8:=",t8[0]["bb"])  

     type T struct {
           cn     string
           street string
     }

     m := make(map[string]T, 10)
     m["sss"]=T{"Nnn","YYY"}
     fmt.Println(m["sss"].cn)  


     var l []Mst
     l=append(l, Mst{"fff":"dddd"})
     fmt.Println(l[0]) 


     N:= make(map[string]interface{})
     N["sss"]="News Paper"
     fmt.Println(N["sss"]) 


     var P [10]Mst
     P[0]=Mst{"Number":"N-001"}
     P[1]=Mst{"Name":"Василий"}
     fmt.Println(P)    


     // Без кавычек в конце будет ошибка {}
     // https://play.golang.org/p/ISEp2GlVVb
     B := map[string]string{}
     B["ssss"]="Пример"
     fmt.Println(B)    

     K := map[string]string{}
     K["ssss"]="Пример"
     K["Nom"]="H-00123"
     fmt.Println(K)    


     O:= make(map[string][]string)
     O["sss"]=[]string{"ddd","fffff"}
     
     fmt.Printf("Данные в структуре  %T \n",O)         
     fmt.Printf("Cтруктура полностью %#v \n",O)         
     fmt.Println("Результат ", O["sss"])         

      var F [10]Mst
      F[0]=Mst{"sssss":"ddd"}
      F[0]["Gorn"]="Test word"
      F[1]=Mst{"sssss":"One element"}
      F[9]=Mst{"sssss":"Ten element"}
      fmt.Println("Результат ",F[0]["sssss"])              
      

      // Append
      var S []Mst
      ttr:=Mst{"Name":"Test-00", "Tabel":"T-0003"}
      S=append(S,ttr)
      
      ttr=Mst{"Name":"Test-02", "Tabel":"T-00100"}
      S=append(S,ttr)

      for _, ttl:=range(S){
             fmt.Println("Результат ..........",ttl["Name"])              	
      }
      
      fmt.Println("Результат ",S)              
      fmt.Println("Результат ",S[0]["Name"])              
}
