func makeID(s string)(x string){  
var y int = 1  
for i := 0; i < len(s)-1; i++ {  
    y *= int(s[i])  
}  
x = strconv.Itoa(y)  
x = x[0:8]  
x = "tab" + x  
return  
}  
