package author

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID         int `gorm:"primaryKey;autoIncrement" json:"Id"`
	AuthorName string
}
