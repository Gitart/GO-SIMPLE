type UserRepository struct{db *gorm.DB}

func NewRepository(db *gorm.DB) UserRepository {
    return UserRepository{db: db}
}

func (r UserRepository) Get(userID uint) (*User, error) {
    entity := new(User)
    err := r.db.Limit(limit: 1).Where(query: "user_id = ?", userID).Find(entity).Error
    return entity, err
}

func (r UserRepository) Create(entity *User) error {
    return r.db.Create(entity).Error
}
func (r UserRepository) Update(entity *User) error {
    return r.db.Model(entity).Update(entity).Error
}

func (r UserRepository) Delete(entity *User) error {
    return r.db.Delete(entity).Error
}
