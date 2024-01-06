package models

import "time"

type Users struct {
	ID        uint       `gorm:"primaryKey;AUTO_INCREMENT;not null" json:"id"`
	Username  string     `gorm:"type:varchar(100);not null;column:username" json:"username"`
	Email     string     `gorm:"type:varchar(50);not null;unique;column:email" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null;column:password" json:"password"`
	Photos    []Photos   `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE" json:"photos"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
