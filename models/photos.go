package models

type Photos struct {
	ID       uint   `gorm:"primaryKey;AUTO_INCREMENT;not null" json:"id"`
	Title    string `gorm:"type:varchar(255);not null;column:title" json:"title"`
	Caption  string `gorm:"type:varchar(100);not null;column:caption" json:"caption"`
	PhotoUrl string `gorm:"type:text;not null;column:photo_url" json:"photo_url"`
	User_ID  int64  `gorm:"not null;column:user_id" json:"user_id"`
}
