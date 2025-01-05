package service

import (
	"context"
	"fmt"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/SmakTown-company/Backend/payment/internal/repository"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"os"
	"time"
)

type PaymentService struct {
	PaymentMethod  string
	ReturnURL      string
	vatCode        string
	PaymentClient  *yookassa.Client
	PaymentHandler *yookassa.PaymentHandler
	repo           repository.Payment
	ctx            context.Context
}

func NewPaymentService(paymentMethod, returnURL string, repo repository.Payment, ctx context.Context) *PaymentService {
	p := PaymentService{
		PaymentMethod: paymentMethod,
		ReturnURL:     returnURL,
		repo:          repo,
		ctx:           ctx,
		vatCode:       "1",
	}
	p.PaymentClient = yookassa.NewClient(os.Getenv("PAYMENT_ACCOUNT_ID"), os.Getenv("PAYMENT_SECRET_KEY"))
	p.PaymentHandler = yookassa.NewPaymentHandler(p.PaymentClient)
	return &p
}
func (p *PaymentService) MakePayment(payment model.Payment) (model.Payment, error) {
	amount := 0.0
	for _, v := range payment.ProductList {
		amount = amount + v.Amount
	}
	items := make([]*yoocommon.Item, len(payment.ProductList))
	for i, product := range payment.ProductList {
		items[i] = &yoocommon.Item{
			Description: product.Name,
			Quantity:    product.Quantity,
			Amount: &yoocommon.Amount{
				Value:    fmt.Sprintf("%.2f", product.Amount),
				Currency: product.Currency,
			},
			VatCode: p.vatCode,
		}
	}
	timeStart := time.Now()
	timeDeadline := timeStart.Add(time.Minute * 10)
	paymentReq, err := p.PaymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    fmt.Sprintf("%.2f", amount),
			Currency: payment.Currency,
		},
		PaymentMethod: yoopayment.PaymentMethodType(p.PaymentMethod),
		Confirmation:  yoopayment.Redirect{Type: "redirect", ReturnURL: p.ReturnURL},
		//Receipt: &yoopayment.Receipt{
		//	Customer: &yoocommon.Customer{
		//		Phone: payment.Phone,
		//		Email: payment.Email,
		//	},
		//	Items: items,
		//},
		CreatedAt: &timeStart, // Приведение к типу библиотеки
		ExpiresAt: &timeDeadline,
	})
	if err != nil {
		return model.Payment{}, err
	}
	//paymentReq.Confirmation
	url := paymentReq.Confirmation.(map[string]interface{})["confirmation_url"].(string)
	payment.URL = url
	payment.ExpiresAt = &timeDeadline
	payment.CreatedAt = &timeStart
	payment.ID = paymentReq.ID
	payment.Status = &paymentReq.Status
	payment.Amount = paymentReq.Amount
	payment.PaymentMethod = &paymentReq.PaymentMethod
	return payment, nil
}
func (p *PaymentService) SavePayment(payment model.Payment) error {
	return p.repo.WritePayment(payment)
}
func (p *PaymentService) UpdateUser(user model.User) (model.User, error) {
	return p.repo.WriteUser(user)
}

func (p *PaymentService) UpdateStatus(payment model.Payment) error {

}
