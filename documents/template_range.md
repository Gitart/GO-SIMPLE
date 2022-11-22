## Прокурутка в темплейте полей которые есть в массиве


```html
   {{ $data   := .data }}
   {{ $okeys  := .keysOrder }}

    {{range $key := $okeys}}
           <h2>{{$key}} = {{ index $data $key}} </h2>
    {{end}}
  ```
  
  
  
  // ***********************************************************
// 
// Один счет
// http://localhost:1968/account/1248
//
// ***********************************************************
func Account(w http.ResponseWriter, r *http.Request){
     var Dt Mst
     
     Idaccount := r.URL.Path[len("/account/"):]
     Su        := Acc_id_Get(Idaccount).Summ
     keysOrder    := []string {"Idaccount","Account"}
     var data = Mst {
      // "Dat"         :  AccGet(),                            // Информация по счету
      "Idaccount"   :  Idaccount,
      "Sum"         :  Su,                                     // Сумма по счету  
      "Propis"      :  UaMoney(Su, true) ,             
      "Items"       :  ItemsGet(Idaccount),                    // Итемс 
      "Title"       : "Постачальник",
      "Address"     : "Адреса ",
      "Post"        : "Фiзична особа-пiдприэмець ",
      "Name"        : "Артур",
      "Account"     : "Рахунок: UA 52 300528  ",
      "Bank"        : "Банк: АТ «БАНК»",
      "Edrpou"      : "ЕДРПОУ: 11111038",
      
      "Post_name"   : "ТОВАРИСТВО З ОБМЕЖЕНОЮ ВІДПОВІДАЛЬНІСТЮ ",
      "Post_bank"   : "п/р UA373226690000026009300099265 у банку ПАТ ",
      "Post_addr"   : "вул.",
      "Post_tel"    : "тел.: 111111809",
      "Post_edrpou" : "код за ЄДРПОУ 11111 ІПН 111111111",
     }
     
     
     Dt = Mst{
      "data": data,
      "keysOrder": keysOrder,
     } 
    
    
    
    
    
    
    fp        := path.Join("tmp", "account.html")
	  tmpl, err := template.ParseFiles(fp)

	// Error
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

	if err := tmpl.Execute(w, Dt); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}
