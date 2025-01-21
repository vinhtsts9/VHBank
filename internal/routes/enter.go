package routers

import (
	"Golang-Masterclass/simplebank/internal/routes/manage"
	"Golang-Masterclass/simplebank/internal/routes/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
