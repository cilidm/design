package strategy

import "testing"

func TestPayByCash(t *testing.T) {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func TestPayBank(t *testing.T){
	payment := NewPayment("Tom","0002",888,&Bank{})
	payment.Pay()
}