package util_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/avelinobego/esocial/internal/shared/util"
)

type Customer struct {
	TotalPurchases int
	Active         bool
	LastPurchase   time.Time
}

// Especificações específicas
func IsPremiumCustomer(customer Customer) bool {
	return customer.TotalPurchases > 10000
}

func IsActiveCustomer(customer Customer) bool {
	return customer.Active &&
		time.Since(customer.LastPurchase) < 6*30*24*time.Hour
}

func TestIsActiveCustomer(t *testing.T) {
	customer := Customer{
		TotalPurchases: 15000,
		Active:         true,
		LastPurchase:   time.Now().Add(-5 * 30 * 24 * time.Hour),
	}

	eligibleSpec := util.And(IsPremiumCustomer, IsActiveCustomer)
	if eligibleSpec(customer) {
		fmt.Println("applyDiscount(customer)")
	}

}
