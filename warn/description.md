## 👆 No Save empty field in GORM
Where used field CheckBox in form, after save to procedure in GO! 

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

## No save
👆 **No Save Check Box !!!**

```go
db.DB.Model(Products{}).
Where("id=?", f.Id).
Updates(&f)
```
