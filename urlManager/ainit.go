package urlManager

import (
	"github.com/jinzhu/gorm"
	"github.com/zhang555/crawler1/db"
)

var (
	DB *gorm.DB
)

func init() {

	DB = db.DB
}
