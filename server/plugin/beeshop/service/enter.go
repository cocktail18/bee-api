package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/service/bee"
)

type ServiceGroup struct {
	BeeServiceGroup bee.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
