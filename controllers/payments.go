package controllers

import (
	"fmt"
	"net/http"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/db"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/models"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/services"
	"github.com/PCSB-Web-Team/FE-registration-portal-backend/utils"
	"github.com/gin-gonic/gin"
)

type PaymentControllerInterface interface {
	CreateOrder(ctx *gin.Context)
	VerifyPayment(ctx *gin.Context)
	GetPayment(ctx *gin.Context)
}

type paymentController struct {
	DB            db.RegistrationsActions
	paymentClient services.PaymentInterface
}

func NewPaymentController(db db.RegistrationsActions, paymentClient services.PaymentInterface) PaymentControllerInterface {
	return &paymentController{
		DB:            db,
		paymentClient: paymentClient,
	}
}

func (c *paymentController) CreateOrder(ctx *gin.Context) {
	var req models.GetRegistration
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("error while parsing request: %v", err.Error())))
		return
	}

	result, notExist := c.DB.GetRegistration(req.Email)
	if notExist != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(fmt.Errorf("error while fetching user: %v", notExist.Error())))
		return
	}
	if result.PaymentID != "" {
		ctx.JSON(http.StatusFound, result)
		return
	}

	createdOrder, err := c.paymentClient.CreateOrder(result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Errorf("error while creating order: %v", err.Error())))
		return
	}
	ctx.JSON(http.StatusCreated, createdOrder["id"])
}

func (c *paymentController) VerifyPayment(ctx *gin.Context) {
	var successfulPaymentResponse models.SuccessfulPayment
	if err := ctx.ShouldBindJSON(&successfulPaymentResponse); err != nil {
		fmt.Printf("error from razorpay while verifying payment: %v\n", err.Error())
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	// order_id := ctx.DefaultQuery("id", "Guest")
	// generated_signature := hmac_sha256(order_id+"|"+successfulPaymentResponse.Razorpay_payment_id, os.Getenv("RAZORPAY_KEY_SECRET"))

	// if generated_signature == successfulPaymentResponse.Razorpay_signature {
	// }

	// type paymentOptions struct {
	// 	notes map[string]string
	// }

	// payment, _ := json.Marshal(paymentStatus)

	result, err := c.DB.GetRegistration(successfulPaymentResponse.Email)
	if err != nil {
		fmt.Printf("error from server while verifying payment: %v\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	result.PaymentID = successfulPaymentResponse.Razorpay_payment_id
	registeredUser, err := c.DB.CreateRegistration(result)
	if err != nil {
		fmt.Printf("error from server while verifying payment: %v\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, registeredUser)
}

func (c *paymentController) GetPayment(ctx *gin.Context) {
	result, err := c.paymentClient.FetchPayment(ctx.Query("id"))
	if err != nil {
		fmt.Printf("error from server while fetching payment: %v\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusFound, result["notes"])
}
