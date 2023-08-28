## No Save empty field in GORM

```go
// UpdateProduct
// ✨ Update a product
// При удалении галочки в CheckBox - пустое поле не сохраняется GORM !!!!
// Warning !
// ! Save   - Сохраняет галочку в CheckBox !!!!
// ! Upadte - HE Сохраняет галочку в CheckBox !!!!
func UpdateProduct(contact *Products) error {
	return db.DB.Save(contact).Error
}
```
