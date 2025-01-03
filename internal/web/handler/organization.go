package handler

import (
	"errors"
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
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
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
		c.Set(middleware.ResponseContextKey, response.NewMissingRequiredParamsError())
		return
	}

	service := organization.NewService(h.repository)
	res, err := service.Create(c, req.OrganizationName)
	if err != nil {
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusCreated,
		Data:       res,
	})
}

func (h *OrganizationHandler) Update(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.OrganizationRequest)
	if empty.Is(req.OrganizationName) {
		c.Set(middleware.ResponseContextKey, response.NewMissingRequiredParamsError())
		return
	}

	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := organization.NewService(h.repository)
	res, err := service.Update(c, &organization.UpdateInput{
		OrganizationID:   pathParams.ID,
		OrganizationName: req.OrganizationName,
	})
	if err != nil {
		if errors.Is(err, organization.ErrOrganizationNotFound) {
			c.Set(middleware.ResponseContextKey, response.NewNotFoundError(err))
			return
		}
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
		Data:       res,
	})
}

func (h *OrganizationHandler) Delete(c *gin.Context) {
	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := organization.NewService(h.repository)
	err := service.Delete(c, pathParams.ID)
	if err != nil {
		if errors.Is(err, organization.ErrOrganizationNotFound) {
			c.Set(middleware.ResponseContextKey, response.NewNotFoundError(err))
			return
		}
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
	})
}

func (h *OrganizationHandler) Restore(c *gin.Context) {
	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := organization.NewService(h.repository)
	err := service.Restore(c, pathParams.ID)
	if err != nil {
		if errors.Is(err, organization.ErrOrganizationNotFound) {
			c.Set(middleware.ResponseContextKey, response.NewNotFoundError(err))
			return
		}
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
	})
}
