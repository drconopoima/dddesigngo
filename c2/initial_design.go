package c2

import (
	"context"
)

type SubscriptionPlan int

const (
	unknownSubscriptionPlan SubscriptionPlan = iota
	basic
	premium
	exclusive
)

type AccountRole int

const (
	unknownAccountRole AccountRole = iota
	lead
	customer
	churned
	lostLead
)

type PaymentDetails struct {
	stripeTokenID string
}

type AccountCreateForm struct {
	AccountRole      AccountRole
	Email            string
	SubscriptionPlan SubscriptionPlan
	PaymentDetails   PaymentDetails
}

type AccountModifyForm struct {
	ID               string
	AccountRole      AccountRole
	Email            string
	SubscriptionPlan SubscriptionPlan
	PaymentDetails   PaymentDetails
}

type Account struct {
	ID             string
	Email          string
	PaymentDetails PaymentDetails
}

type AccountManager interface {
	AddAccount(ctx context.Context, form AccountCreateForm) (Account, error)
	ModifyAccount(ctx context.Context, form AccountModifyForm) (Account, error)
}
