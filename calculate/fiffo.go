package fifo

import (
	"boiler/controllers/db"
	"boiler/controllers/employees"
	"boiler/controllers/ledger_book"
	"boiler/controllers/orders"
	"boiler/controllers/stock_list"
	"boiler/controllers/system"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Процедура фифо со склада
// она работатет только на вычитание
type (

	// Контрагент
	Comp struct {
		Title   string  `json:"title"`   // Наименование контрагента
		Qty     float64 `json:"qty"`     // Количество на складе
		Price   float64 `json:"price"`   // Цена
		Summ    float64 `json:"summ"`    // Сумма
		Ostatok float64 `json:"ostatok"` // Остаток
	}

	// Контрагент
	Itms struct {
		Id        int64   `json:"id"`         // Id items
		OrderId   int64   `json:"order_id"`   // Id Документа
		ProductId int64   `json:"product_id"` // Id продукта
		Product   string  `json:"product"`    // Продукт
		Qty       float64 `json:"qty"`        // Количество на складе
		Price     float64 `json:"price"`      // Цена
		Summ      float64 `json:"summ"`       // Сумма
		Fifo      float64 `json:"fifo"`       // Остаток после фифо
		Diff      float64 `json:"diff"`       // Разница между Остском и снятием
	}

	// Общее количество итоговые суммы
	Totals struct {
		DateTime   string  `json:"date_time"`   // Дата и  время выполнения
		TotalQty   float64 `json:"total_qty"`   // Общее количество товара
		TotalSum   float64 `json:"total_sum"`   // Общая сумма
		TotalCnt   int64   `json:"total_cnt"`   // Общее количество позиций
		Operation  string  `json:"operation"`   // Наименование операции
		Result     string  `json:"result"`      // Результат можно снимать или нет
		NotifyUser string  `json:"notify_user"` // Сообщение пользователю
		Qty        float64 `json:"qty"`         // Сколько снимаем количество
		Diff       float64 `json:"diff"`        // Разница между Qty и наличием сейчас на складах
		Companies  []Comp  `gorm:"-"`           // Компании Исключение обработки GORM
		Items      []Itms  `gorm:"-"`           // Продукты
	}
)

// FIFO: Test
func MethodFifo(c echo.Context) error {

	idproduct := system.StrToFloat(c.Param("idproduct"))
	qty := system.StrToFloat(c.Param("qty"))

	total := FiFo(idproduct, qty)

	Dat := echo.Map{
		"Totals": total,
	}

	return c.JSON(http.StatusOK, Dat)
}

// Метод FiFo
func FiFo(product_id, qty float64) Totals {

	inqty := qty
	// Контроль сколько товара всего на складах
	var total Totals

	//var comp []Comp
	var itms []Itms

	// Общий остаток продукта по всем складам
	db.DB.Table("stocks").
		Select("product_id, SUM(qty) AS total_qty, SUM(summ) AS total_sum, COUNT(*) AS total_cnt").
		Where("product_id=?", product_id).
		Group("product_id").
		Scan(&total)

	// Заполнение результатат
	total.DateTime = system.TodayDate()
	total.Qty = qty
	total.Diff = total.TotalQty - qty
	total.Operation = "Списание по методу FiFo"

	// Опредление можно ли производить списание со складов
	if total.TotalQty-qty <= 0 {
		total.Result = "empty"
		total.NotifyUser = "Нельзя снимать со склада нет достаточного количества"
	} else {
		total.Result = "full"
		total.NotifyUser = "Можно снимать со склада есть необходимое количество на складах"
	}

	// Определяем у кого сколько есть этого товара на складе
	// db.DB.Table("stocks").
	//	Select("product_id, qty, summ AS summ, contragent_name AS title, price").
	//	Where("product_id=?", product_id).
	//	Scan(&comp)

	// Поиск в итемс
	db.DB.Table("order_items").
		Select("id, order_id, product_id, product, qty, summ, price, fifo").
		Where("product_id=?", product_id).
		//Where("qty!=fifo").
		Scan(&itms)

	// Если на складах есть столько прокручиваем
	// ТТН которые учатсвовали в наполненни этих
	// складов

	// Остатко после вычитания
	ostatok := 0.0

	// Снятие количества товара по очереди
	for i, itm := range itms {

		// Остаток есть
		if itm.Qty != itm.Fifo {

			fmt.Println(i, qty, itm.Qty, itm.Fifo, itm.Product)

			if itm.Qty <= itm.Fifo+qty+ostatok {
				ostatok = itm.Qty - itm.Fifo
				ItemsFifoUpdate(itm.Id, itm.Qty)
				break
			}

			if itm.Qty >= itm.Fifo+qty+ostatok {
				ItemsFifoUpdate(itm.Id, itm.Fifo+qty)
				break
			}
		}

	}

	// Update fifo in TTN
	for _, i := range itms {

		// ItemsFifoUpdate(i.Id, i.Qty)

		// Информация о текущем документе
		document := orders.OrderInfo(i.OrderId)

		// Запись в Главную книгу (trigger)
		lb := ledger_book.LedgerBooks{
			DocId:          i.OrderId,
			Ttn:            document.Ttn + " FIFO",
			ProductId:      i.ProductId,
			ProductName:    i.Product,
			DocNumber:      document.Num,
			ContragentId:   document.CompanyId,
			ContragentName: document.Company,
			ContractId:     document.ContractId,
			ContractNumber: document.Contract,
			StockFrom:      document.StockFrom,
			StockTo:        document.StockId,
			Stock:          stock_list.StockName(document.StockId),
			Qty:            i.Qty,
			Price:          i.Price,
			Summ:           i.Summ * i.Qty,
			Remain:         i.Qty,
			Operation:      "виробництво",
			UserId:         document.UserId,
			UserName:       employees.GetNameEmployeeID(document.UserId),
			AccountNumber:  "287-02",
			Note:           "FiFo",
		}

		if i.Qty > 0 {

			// Главная книга
			ledger_book.Add(lb)

			// Партия
			AddParty(ProductsParties{
				DocumentId:   i.OrderId,
				Qty:          inqty,
				ProductId:    i.ProductId,
				ProductName:  i.Product,
				Price:        i.Price,
				StockId:      document.StockId,
				ContragentId: document.Company,
				Remain:       i.Qty - inqty,
				InQty:        inqty,
				OutQty:       i.Qty,
			})
		}
	}

	total.Items = itms

	return total
}

// Обновление в order_items FiFo
func ItemsFifoUpdate(id int64, fifo float64) {

	// fix for GORM
	if fifo == 0 {
		fifo = 0.000001
	}

	i := orders.OrderItems{Fifo: fifo}

	db.DB.Model(orders.OrderItems{}).
		Where("id=?", id).
		Updates(&i)
}
