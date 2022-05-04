package lead

import (
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Comapny string
	Email   string
	Phone   int
}
