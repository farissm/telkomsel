package product

import (
	"gorm.io/gorm"
	"time"

	log "github.com/sirupsen/logrus"
)

type Repository interface {
	createProduct(product Product) (response Product, err error)
	getProduct(page int) (products []ProductResponse, err error)
	getProductById(produtId int) (product ProductResponse, err error)
	updateProductById(produtId int, productPayload Product) (err error)
	deleteProductById(produtId int) (err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) createProduct(product Product) (response Product, err error) {
	err = r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) getProduct(page int) (products []ProductResponse, err error) {
	var limitPage int = 5
	var offsetPage int = (page - 1) * limitPage

	sql := `SELECT * FROM products LIMIT ? OFFSET ?`
	
	if err = r.db.Raw(sql, limitPage, offsetPage).Find(&products).Error; err != nil {
		log.Errorf("Error get produts at repository %v", err)
		return
	}

	return
}

func (r *repository) getProductById(productId int) (product ProductResponse, err error) {

	sql := `SELECT * FROM products where id = ?`
	
	if err = r.db.Raw(sql, productId).Find(&product).Error; err != nil {
		log.Errorf("Error get produt by id at repository %v", err)
		return
	}

	return
}

func (r *repository) updateProductById(productId int, productPayload Product) (err error) {

	sql := `UPDATE products SET name = COALESCE(NULLIF(?, ''), name), description = COALESCE(NULLIF(?, ''), description), price = COALESCE(NULLIF(?, 0), price), Variety = COALESCE(NULLIF(?, ''), Variety), rating = COALESCE(NULLIF(?, 0), rating), stock = COALESCE(NULLIF(?, 0), stock), total_sold = COALESCE(NULLIF(?, 0), total_sold), updated_at = ? WHERE id = ?`
	
	if err = r.db.Exec(sql, productPayload.Name, productPayload.Description, productPayload.Price, productPayload.Variety, productPayload.Rating, productPayload.Stock, productPayload.TotalSold, time.Now().Format("2006-01-02 15:04:05"), productId).Error; err != nil {
		log.Errorf("Error update product at repository %v", err)
		return
	}

	return nil
}

func (r *repository) deleteProductById(productId int) (err error) {
	if err = r.db.Unscoped().Delete(&Product{}, productId).Error; err != nil {
		log.Errorf("Error delete product at repository %v", err)
		return 
	}

	return nil
}