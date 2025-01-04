package web

import "github.com/YuukiHayashi0510/todo-app/internal/web/handler"

type Handlers struct {
	Organizations handler.OrganizationHandler
	Staffs        handler.StaffHandler
}
