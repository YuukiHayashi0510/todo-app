package handler

import (
	"net/http"

	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
	"github.com/YuukiHayashi0510/todo-app/internal/app/repository"
	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/postgres"
	"github.com/YuukiHayashi0510/todo-app/internal/web/middleware"
	"github.com/YuukiHayashi0510/todo-app/internal/web/request"
	"github.com/YuukiHayashi0510/todo-app/internal/web/response"
	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	repository organization.Repository
}

func NewOrganizationHandler(db *postgres.DB) OrganizationHandler {
	return OrganizationHandler{repository: repository.NewOrganizationRepository(db)}
}

func (h *OrganizationHandler) List(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.OrganizationRequest)

	service := organization.NewService(h.repository)
	res, err := service.Search(c, &organization.SearchInput{
		Organization: organization.Organization{
			OrganizationID:   req.OrganizationID,
			OrganizationName: req.OrganizationName,
		},
	})
	if err != nil {
		c.Set(middleware.ResponseContextKey, response.NewInternalServerErrorResponse(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
		Data:       res,
	})
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.OrganizationRequest)
	if empty.Is(req.OrganizationName) {
		c.Set(middleware.ResponseContextKey, response.NewMissingRequiredParamsErrorResponse())
		return
	}

	service := organization.NewService(h.repository)
	res, err := service.Create(c, req.OrganizationName)
	if err != nil {
		c.Set(middleware.ResponseContextKey, response.NewInternalServerErrorResponse(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusCreated,
		Data:       res,
	})
}
