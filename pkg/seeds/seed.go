package seeds

import (
	"xorm.io/xorm"
)

type Seed struct {
	Name string
	Run  func(engine *xorm.Engine) error
}
