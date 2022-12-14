package controllers

import (
	"fmt"
	"net/http"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/db"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/models"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/utils"
	"github.com/gin-gonic/gin"
)

type RegistrationsControllerInterface interface {
	CreateRegistration(ctx *gin.Context)
	GetRegistration(ctx *gin.Context)
}

type registrationsController struct {
	DB db.RegistrationsActions
}

func NewRegistrationsController(db db.RegistrationsActions) RegistrationsControllerInterface {
	return &registrationsController{
		DB: db,
	}
}

func (c *registrationsController) CreateRegistration(ctx *gin.Context) {
	var req models.Registration
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if _, notExist := c.DB.GetRegistration(req.Email); notExist != nil {
		result, err := c.DB.CreateRegistration(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	} else {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(fmt.Errorf("'%s' is already registered", req.Email)))
	}
}

func (c *registrationsController) GetRegistration(ctx *gin.Context) {
	var req models.GetRegistration
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	result, err := c.DB.GetRegistration(req.Email)
	if err != nil {
		ctx.JSON(http.StatusNoContent, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusFound, result)
}
