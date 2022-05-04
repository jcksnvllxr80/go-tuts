package lead

import (
	"github.com/jcksnvllxr80/go-tuts/go-fiber-crm/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string
	Comapny string
	Email   string
	Phone   int
}
