package main

import (
	"CRM/src/app/api/admin"
	"CRM/src/app/api/agent"
	"CRM/src/app/api/common"
	"CRM/src/app/api/user"
)

func initRouter() {
	commonService := common.New(app)
	commonService.InitRouter()

	adminService := admin.New(app)
	adminService.InitRouter()

	agentService := agent.New(app)
	agentService.InitRouter()

	userService := user.New(app)
	userService.InitRouter()
}
