package routes

import (
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(registrationController controllers.RegistrationsControllerInterface, paymentController controllers.PaymentControllerInterface) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/register", registrationController.CreateRegistration)
	router.POST("/register/verify", registrationController.GetRegistration)
	router.POST("/register/create-order", paymentController.CreateOrder)
	router.POST("/register/verify-order", paymentController.VerifyPayment)
	return router
}
