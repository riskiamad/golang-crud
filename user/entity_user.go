package user

import "time"

type User struct {
	ID        int       `gorm:"auto;primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(100)" json:"username"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Password  string    `gorm:"type:varchar(100)" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;default:NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:NULL" json:"updated_at"`
}
