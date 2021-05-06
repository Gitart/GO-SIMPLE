package main

import (
    "fmt"
     // "github.com/tealeg/xlsx"
     // "github.com/extrame/xls"
        "github.com/360EntSecGroup-Skylar/excelize"
        "strings"
        "strconv"
)



func Trm (st string) string {
    return strings.Trim(st," ")
}

func StrToFloat(num string) float64 {
	f, _ := strconv.ParseFloat(num, 64)
	return f
}


func main() {
	sum:=0.00
    xlsx, err := excelize.OpenFile("./test.xlsx")
    
    if err != nil {
       fmt.Println(err)
       return
    }

    // Получение данных из ячейки
    // Get value from cell by given worksheet name and axis.
    // cell,_ := xlsx.GetCellValue("Sheet1", "B2")
    // fmt.Println(cell)
    
    // Get all the rows in the Sheet1.
    rows,_ := xlsx.GetRows("Sheet1")
    

    for idx, row := range rows {

        // for _, colCell := range row {
        //     fmt.Print(colCell, "\t")
        // }
        
       
       // Skip first row (Titles)
       if idx == 0 {
          continue
       }

        rr:=Trm(row[0])
       
        if rr == "" {
        	continue
        }
        
        rs := Trm(row[2])
        rv := 0.00
        rv  = StrToFloat(row[3])
        fmt.Printf("id %5v %-80s %15s %20s %9.2f %15v \n", idx, rr ,row[1], rs, rv, row[4])
        sum += StrToFloat(row[3])
    }


    fmt.Printf("****************************************** \n Общая сумма : %9.2f \n", sum + 0.00)
}




// func mains(){
//      xlsload("test.xls")
// }





// func xlsload(xlsfile string){
// if xlFile, err :=excelize.OpenFile(xlsfile, "utf-8"); err == nil {
	
// 	if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
		
// 		fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
		
// 		col1 := sheet1.Row(0).Col(0)
// 		col2 := sheet1.Row(0).Col(0)
		
// 		for i := 0; i <= (int(sheet1.MaxRow)); i++ {
// 			row1 := sheet1.Row(i)
// 			col1  = row1.Col(0)
// 			col2  = row1.Col(1)
			
// 			fmt.Print("\n", col1, ",", col2)
// 		}
// 	}
// }

// }







