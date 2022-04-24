package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(user User) (User, error) {
	if err := r.db.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
