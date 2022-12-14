package models

type Registration struct {
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	RollNo      string `json:"roll_no" binding:"required"`
	Expectation string `json:"expectation"`
	PaymentID   string `json:"payment_id"`
	// Year        string `json:"year" binding:"required,oneof=FE SE TE"`
	// Division    string `json:"division" binding:"required"`
	// Department  string `json:"department" binding:"required,oneof=CS IT ENTC"`
}

type GetRegistration struct {
	Email string `json:"email" binding:"required"`
}
