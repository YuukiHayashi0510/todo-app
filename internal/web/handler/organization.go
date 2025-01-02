package handler

import (
	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	repository organization.Repository
}

func (h *OrganizationHandler) List(c *gin.Context) {
	// service := organization.NewService(h.repository)

}
