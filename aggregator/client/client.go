package client

import (
	"context"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

type Client interface{
	Aggregate(context.Context, *types.AggregateRequest) (error)
}