package services

import (
	"fmt"
	"time"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/models"
	"github.com/razorpay/razorpay-go"
)

type PaymentInterface interface {
	CreateOrder(studentDetails *models.Registration) (map[string]interface{}, error)
	FetchPayment(paymentId string) (map[string]interface{}, error)
}

type paymentClient struct {
	client *razorpay.Client
}

func NewPaymentClient(RAZORPAY_KEY_ID string, RAZORPAY_KEY_SECRET string) PaymentInterface {
	client := razorpay.NewClient(RAZORPAY_KEY_ID, RAZORPAY_KEY_SECRET)
	return &paymentClient{
		client,
	}
}

func (c *paymentClient) CreateOrder(studentDetails *models.Registration) (map[string]interface{}, error) {
	receipt := fmt.Sprintf("%v_%v", time.Now().Unix(), studentDetails.RollNo)
	data := map[string]interface{}{
		"amount":          36000,
		"currency":        "INR",
		"receipt":         receipt,
		"partial_payment": false,
		"notes": map[string]interface{}{
			"email": studentDetails.Email,
		},
	}
	orderResponseBody, err := c.client.Order.Create(data, nil)
	if err != nil {
		return nil, err
	}
	return orderResponseBody, nil
}

func (c *paymentClient) FetchPayment(paymentId string) (map[string]interface{}, error) {
	paymentStatus, err := c.client.Payment.Fetch(paymentId, nil, nil)
	if err != nil {
		return nil, err
	}
	return paymentStatus, nil
}
