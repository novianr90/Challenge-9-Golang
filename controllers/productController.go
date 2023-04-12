package controllers

import (
	"challenge-9/helpers"
	"challenge-9/models"
	"challenge-9/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductController struct {
	ProductService *services.ProductService
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var (
		productDto ProductDto

		data = c.MustGet("data").(map[string]any)

		userData = data["user"].(models.User)

		contentType = helpers.GetContentType(c)

		userId = userData.ID
	)

	if contentType == appJson {
		if err := c.ShouldBindJSON(&productDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&productDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}

	product := models.Product{
		Name:        productDto.Name,
		Description: productDto.Description,
		UserID:      userId,
	}

	result, err := pc.ProductService.CreateProduct(product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          result.ID,
		"name":        result.Name,
		"description": result.Description,
	})
}

func (pc *ProductController) GetAllProductByUserId(c *gin.Context) {
	var (
		products []models.Product

		data = c.MustGet("data").(map[string]any)

		userData = data["user"].(models.User)

		isAdmin = data["isAdmin"].(bool)

		err error
	)

	if isAdmin {
		products, err = pc.ProductService.GetAllProduct()
	} else {
		products, err = pc.ProductService.GetAllProductsByUserId(userData.ID)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var (
		contentType = helpers.GetContentType(c)

		mapId = c.MustGet("mapId").(map[string]uint)

		productDto ProductDto

		err error
	)

	if contentType == appJson {
		if err = c.ShouldBindJSON(&productDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {
		if err = c.ShouldBind(&productDto); err != nil {
			if err = c.ShouldBindJSON(&productDto); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, err)
				return
			}
		}
	}

	updatedProduct := models.Product{
		ID:          mapId["productId"],
		Name:        productDto.Name,
		Description: productDto.Description,
		UserID:      mapId["userId"],
	}

	err = pc.ProductService.UpdateProduct(updatedProduct.ID, updatedProduct)

	if err != nil {
		if err = c.ShouldBindJSON(&productDto); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"product": updatedProduct,
	})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	mapId := c.MustGet("mapId").(map[string]uint)

	if err := pc.ProductService.DeleteProduct(mapId["productId"]); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product sucesfully deleted",
	})
}
