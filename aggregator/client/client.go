package client

import (
	"context"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error 		//being used by distance calaculator to aggregate the tolls
	GetInvoice(context.Context, int) (*types.Invoice, error)  		//being used by gateway to get the invoice
}
