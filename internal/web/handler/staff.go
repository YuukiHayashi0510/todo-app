package handler

import (
	"errors"
	"net/http"

	"github.com/YuukiHayashi0510/todo-app/internal/domain/repository"
	"github.com/YuukiHayashi0510/todo-app/internal/domain/staff"
	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/postgres"
	"github.com/YuukiHayashi0510/todo-app/internal/web/middleware"
	"github.com/YuukiHayashi0510/todo-app/internal/web/request"
	"github.com/YuukiHayashi0510/todo-app/internal/web/response"
	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	repository staff.Repository
}

func NewStaffHandler(db *postgres.DB) StaffHandler {
	return StaffHandler{repository: repository.NewStaffRepository(db)}
}

func (h *StaffHandler) List(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.StaffRequest)

	service := staff.NewService(h.repository)
	res, err := service.Search(c, &staff.SearchInput{
		Staff: staff.Staff{
			OrganizationID: req.OrganizationID,
			StaffID:        req.StaffID,
			StaffName:      req.StaffName,
			Email:          req.Email,
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

func (h *StaffHandler) Create(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.StaffRequest)
	if empty.Any(req.OrganizationID, req.StaffName, req.Email) {
		c.Set(middleware.ResponseContextKey, response.NewMissingRequiredParamsError())
		return
	}

	service := staff.NewService(h.repository)
	res, err := service.Create(c, &staff.CreateInput{
		OrganizationID: req.OrganizationID,
		Email:          req.Email,
		StaffName:      req.StaffName,
	})
	if err != nil {
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusCreated,
		Data:       res,
	})
}

func (h *StaffHandler) Update(c *gin.Context) {
	req := c.MustGet(middleware.ValidationContextKey).(*request.StaffRequest)
	if empty.Any(req.OrganizationID, req.StaffName, req.Email) {
		c.Set(middleware.ResponseContextKey, response.NewMissingRequiredParamsError())
		return
	}

	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := staff.NewService(h.repository)
	res, err := service.Update(c, &staff.UpdateInput{
		OrganizationID: req.OrganizationID,
		StaffID:        pathParams.ID,
		Email:          req.Email,
		StaffName:      req.StaffName,
	})
	if err != nil {
		if errors.Is(err, staff.ErrStaffNotFound) {
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

func (h *StaffHandler) Delete(c *gin.Context) {
	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := staff.NewService(h.repository)
	err := service.Delete(c, pathParams.ID)
	if err != nil {
		if errors.Is(err, staff.ErrStaffHasAlreadyDeleted) {
			c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
			return
		} else if errors.Is(err, staff.ErrStaffNotFound) {
			c.Set(middleware.ResponseContextKey, response.NewNotFoundError(err))
			return
		}
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
		Data: gin.H{
			"message": "deleted successfully",
		},
	})
}

func (h *StaffHandler) Restore(c *gin.Context) {
	var pathParams request.PathParams
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
		return
	}

	service := staff.NewService(h.repository)
	err := service.Restore(c, pathParams.ID)
	if err != nil {
		if errors.Is(err, staff.ErrStaffIsNotDeleted) {
			c.Set(middleware.ResponseContextKey, response.NewBadRequestError(err))
			return
		} else if errors.Is(err, staff.ErrStaffNotFound) {
			c.Set(middleware.ResponseContextKey, response.NewNotFoundError(err))
			return
		}
		c.Set(middleware.ResponseContextKey, response.NewInternalServerError(err))
		return
	}

	c.Set(middleware.ResponseContextKey, &response.Response{
		HttpStatus: http.StatusOK,
		Data: gin.H{
			"message": "restored successfully",
		},
	})
}
