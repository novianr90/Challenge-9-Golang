package services

import (
	"challenge-9/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (ps *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	if err := ps.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (ps *ProductService) GetProductById(productId uint) (models.Product, error) {
	var product models.Product

	if err := ps.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (ps *ProductService) GetUserIdByProductId(productId uint) (models.Product, error) {
	var product models.Product

	if err := ps.DB.Select("user_id").First(&product, productId).Error; err != nil {
		return models.Product{}, errors.New("no data found")
	}

	return product, nil
}

func (ps *ProductService) GetAllProductsByUserId(userId uint) ([]models.Product, error) {
	var products []models.Product

	if err := ps.DB.Where("user_id = ?", userId).Find(&products).Error; err != nil {
		return []models.Product{}, err
	}

	fmt.Println("products")

	return products, nil
}

func (ps *ProductService) UpdateProduct(productId uint, product models.Product) error {
	var modelProduct models.Product
	if err := ps.DB.Model(&modelProduct).Where("id = ?", productId).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetAllProduct() ([]models.Product, error) {
	var products []models.Product

	if err := ps.DB.Find(&products).Error; err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (ps *ProductService) DeleteProduct(productId uint) error {
	var modelProduct models.Product

	if err := ps.DB.Where("id = ?", productId).Delete(&modelProduct).Error; err != nil {
		return err
	}

	return nil
}
