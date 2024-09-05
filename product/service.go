package product

import (
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateProduct(product Product) (response ProductResponse, err error)
	GetProducts(page int) (response []ProductResponse, err error)
	GetProductById(productId int) (response ProductResponse, err error)
	UpdateProduct(productId int, product Product) (err error)
	DeleteProduct(productId int) (err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateProduct(product Product) (response ProductResponse, err error) {
	result, err := s.repository.createProduct(product)
	if err != nil {
		log.Errorf("error create product at service level. Error: %s", err.Error())
		return response, err
	}

	response.ID				= result.ID
	response.Name			= result.Name
	response.Description	= result.Description
	response.Price			= result.Price
	response.Variety		= result.Variety
	response.Rating			= result.Rating
	response.Stock			= result.Stock
	response.TotalSold		= result.TotalSold

	return response, nil
}

func (s *service) GetProducts(page int) (response []ProductResponse, err error) {
	result, err := s.repository.getProduct(page)
	if err != nil {
		log.Errorf("error get product in service level. Error: %s", err.Error())
		return response, err
	}

	return result, err
}

func (s *service) GetProductById(productId int) (response ProductResponse, err error) {
	result, err := s.repository.getProductById(productId)
	if err != nil {
		log.Errorf("error get product by id in service level. Error: %s", err.Error())
		return response, err
	}

	return result, err
	
}

func (s *service) UpdateProduct(productId int, product Product) (err error) {
	if err := s.repository.updateProductById(productId, product); err != nil {
		log.Errorf("error update product by id in service level. Error: %s", err.Error())
		return err
	}

	return nil
	
}

func (s *service) DeleteProduct(productId int) (err error) {
	if err := s.repository.deleteProductById(productId); err != nil {
		log.Errorf("error delete product by id in service level. Error: %s", err.Error())
		return err
	}

	return nil
}