package provider

import (
	"client/infrastructure/connect"
	"client/infrastructure/pb"
)

type QueryClientProvider struct {
	Category pb.CategoryServiceClient
	Product  pb.ProductServiceClient
}

func NewQueryClientProvider() (*QueryClientProvider, error) {
	connector := connect.QueryConnector{}
	connect, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	category := pb.NewCategoryServiceClient(connect)
	product := pb.NewProductServiceClient(connect)
	return &QueryClientProvider{Category: category, Product: product}, nil
}
