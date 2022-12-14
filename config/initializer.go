package config

import (
	"log"
	"os"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/controllers"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/db"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/routes"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Initiators interface {
	Start(port string)
}

type app struct {
	registrationTableDBClient db.RegistrationsActions
	registrationsControllers  controllers.RegistrationsControllerInterface
	paymentsControllers       controllers.PaymentControllerInterface
	router                    *gin.Engine
}

func NewApp() Initiators {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var paymentClient services.PaymentInterface = services.NewPaymentClient(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_KEY_SECRET"))

	registrationTableDBClient := db.NewRegistrationTableInstance()
	registrationsControllers := controllers.NewRegistrationsController(registrationTableDBClient)
	paymentControllers := controllers.NewPaymentController(registrationTableDBClient, paymentClient)
	router := routes.NewRouter(registrationsControllers, paymentControllers)
	return &app{
		registrationTableDBClient,
		registrationsControllers,
		paymentControllers,
		router,
	}
}

func (app *app) Start(port string) {
	if err := app.router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server at %s, %s", port, err.Error())
	}
}
