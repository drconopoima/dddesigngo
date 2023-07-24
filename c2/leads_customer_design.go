package c2

import (
	"context"
)

type LeadForm struct {
	email string
}

type Lead struct {
	id    string
	email string
}

type LeadCreator interface {
	CreateLead(ctx context.Context, form LeadForm) (Lead, error)
}

type Customer struct {
	leadID     string
	customerID string
}

func (c *Customer) CustomerID() string {
	return c.customerID
}

func (c *Customer) SetCustomerID(customerID string) {
	c.customerID = customerID
}

type LeadConverter interface {
	Convert(ctx context.Context, subscriptionSelection SubscriptionPlan) (Account, error)
}

func (l Lead) Convert(ctx context.Context, subscriptionSelection SubscriptionPlan) (Customer, error) {
	// TODO implement
	panic("implement me")
}
