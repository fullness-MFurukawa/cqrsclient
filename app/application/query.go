package application

import (
	"client/infrastructure/pb"
	"client/infrastructure/provider"
	"context"
)

type QueryService struct {
	provider provider.QueryClientProvider
}

func (s *QueryService) CagtegoryList() (*pb.CategoriesResult, error) {
	param := &pb.CategoryParam{} // パラメータの生成

	resp, err := s.provider.Category.List(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *QueryService) ProductList() (*pb.ProductsResult, error) {
	param := &pb.ProductParam{}
	resp, err := s.provider.Product.List(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *QueryService) SearchCategoryById(id string) (*pb.CategoryResult, error) {
	param := &pb.CategoryParam{Id: &id}
	resp, err := s.provider.Category.ById(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *QueryService) SearchProductById(id string) (*pb.ProductResult, error) {
	param := &pb.ProductParam{Id: &id}
	resp, err := s.provider.Product.ById(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *QueryService) SearchProductByKeyword(keyword string) (*pb.ProductsResult, error) {
	param := &pb.ProductParam{Keyword: &keyword}
	resp, err := s.provider.Product.ByKeyword(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func NewQueryService() (*QueryService, error) {
	provider, err := provider.NewQueryClientProvider()
	if err != nil {
		return nil, err
	}
	return &QueryService{provider: *provider}, nil
}
