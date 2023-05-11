package application

import (
	"client/infrastructure/pb"
	"client/infrastructure/provider"
	d "client/presentation/dto"
	"context"
)

type CommandService struct {
	provider provider.CommandClientProvider
}

func (s *CommandService) CreateCategory(dto d.CategoryDTO) (*pb.CategoryUpResult, error) {
	param := pb.CategoryUpParam{Name: &dto.Name}
	result, err := s.provider.Category.Create(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CommandService) UpdateCategory(dto d.CategoryDTO) (*pb.CategoryUpResult, error) {
	param := pb.CategoryUpParam{Id: &dto.Id, Name: &dto.Name}
	result, err := s.provider.Category.Update(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CommandService) DeleteCategory(dto d.CategoryDTO) (*pb.CategoryUpResult, error) {
	param := pb.CategoryUpParam{Id: &dto.Id}
	result, err := s.provider.Category.Delete(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CommandService) CreateProduct(dto d.ProductDTO) (*pb.ProductUpResult, error) {
	param := pb.ProductUpParam{Name: &dto.Name, Price: &dto.Price, CategoryId: &dto.CategoryId}
	result, err := s.provider.Product.Create(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CommandService) UpdateProduct(dto d.ProductDTO) (*pb.ProductUpResult, error) {
	param := pb.ProductUpParam{Id: &dto.Id, Name: &dto.Name, Price: &dto.Price, CategoryId: &dto.CategoryId}
	result, err := s.provider.Product.Update(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CommandService) DeleteProduct(dto d.ProductDTO) (*pb.ProductUpResult, error) {
	param := pb.ProductUpParam{Id: &dto.Id}
	result, err := s.provider.Product.Delete(context.Background(), &param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCommandService() (*CommandService, error) {
	provider, err := provider.NewCommandClientProvider()
	if err != nil {
		return nil, err
	}
	return &CommandService{provider: *provider}, nil
}
