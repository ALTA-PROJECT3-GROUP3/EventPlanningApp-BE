package helper

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransPay(OrderID string, GrossAmount int64) *snap.Response {
	// Set up Midtrans configuration
	midtrans.ServerKey = config.MidtransServerKey
	midtrans.ClientKey = config.MidtransClientKey
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(config.MidtransServerKey, midtrans.Sandbox)
	// Create Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  OrderID,
			GrossAmt: GrossAmount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}
	snapResp, _ := snap.CreateTransaction(req)
	return snapResp
}

func MidtransStatusPay(orderID string) (*coreapi.TransactionStatusResponse, error) {
	midtrans.ServerKey = config.MidtransServerKey
	midtrans.ClientKey = config.MidtransClientKey
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(config.MidtransServerKey, midtrans.Sandbox)

	res, err := c.CheckTransaction(orderID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
