package provider

import (
	"client/infrastructure/connect"
	"client/infrastructure/pb"
)

type CommandClientProvider struct {
	Category pb.CategoryUpServiceClient
	Product  pb.ProductUpServiceClient
}

func NewCommandClientProvider() (*CommandClientProvider, error) {
	connector := connect.CommandConnector{}
	connect, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	category := pb.NewCategoryUpServiceClient(connect)
	product := pb.NewProductUpServiceClient(connect)
	return &CommandClientProvider{Category: category, Product: product}, nil
}
