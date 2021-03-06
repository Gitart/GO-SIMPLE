# Вычисление количества дней от текущей даты

## Описание вариантов вычисления дат промокодов

**ds** - дата старта  
**de** - дата финиша  

| ds | de | Status|
|---|---|-------|
|-10 | +2 | Акция идет и еще будет активна 2 дня, а началась 10 дней назад|
|-10 | 0  | Акция закончилась сегодня а началась 10 дней назад|
|-10 | -3 | Акция закончилась 3 дня назад а началась 10 дней назад|
| 10 | 20 | Акция начнется через 10 дней закончится через 20 |

```go    
// Активна или нет промоакция
func PromoIsActive(start, finish time.Time) bool {
   s,e:= DaysPromo(start, finish)
   if s<0 && e>0 {
      return true
   }
      return false
}

// Количество дней от текщей даты до старта promo
// Количество дней от текщей даты до финиша promo
func DaysPromo(start, finish time.Time) (float64, float64) {
    hrs   := DaysDate(start)
    hrd   := DaysDate(finish)
    return hrs,hrd
}

// Количество дней между текущей датой и вводимой
// Прошедшие дни со знаоком минус (-)
// Будущие дни со знаком плюс (+)
func DaysDate(date time.Time) float64 {
     fx   := services.ToFixed  
     today:= time.Now()
     hrs  := fx(today.Sub(date).Hours()/24,0)
     return hrs*-1
}
```

## Округление
```go

// *****************************************************************
// Округление до N знаков
// Библиотека с высокой точностьюdecimal.Decimal
// *****************************************************************
func ToFixedNum(num float64, precision int32)  decimal.Decimal {
     sumd:= decimal.NewFromFloat(num).Round(precision)
     return sumd
 }

// *****************************************************************
// Округление до N знаков
// *****************************************************************
func ToFixed(num float64, precision int) float64 {
     output := math.Pow(10, float64(precision))
     return float64(round(num * output)) / output
}

func round(num float64) int {
     return int(num + math.Copysign(0.5, num))
}
```

