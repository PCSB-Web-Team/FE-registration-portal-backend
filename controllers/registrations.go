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
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Errorf("error from server while creating registration: %v", err.Error())))
		return
	}

	if user, _ := c.DB.GetRegistration(req.Email); user != nil {
		if user.PaymentID == "" {
			result, err := c.DB.CreateRegistration(&req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Errorf("error from server while saving registration: %v", err.Error())))
				return
			}
			ctx.JSON(http.StatusCreated, result)
			return
		} else {
			ctx.JSON(http.StatusFound, user)
		}
	} else {
		result, err := c.DB.CreateRegistration(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Errorf("error from server while saving registration: %v", err.Error())))
			return
		}
		ctx.JSON(http.StatusCreated, result)
	}
}

func (c *registrationsController) GetRegistration(ctx *gin.Context) {
	var req models.GetRegistration
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	result, err := c.DB.GetRegistration(req.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusFound, result)
}
