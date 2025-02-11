package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type ProductController struct {
	Controllers
	ProductService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (productController *ProductController) List(response http.ResponseWriter, request *http.Request) {
	products, err := services.NewProductService().List(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(products.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json response")
		return
	}
}

func (productController *ProductController) Detail(response http.ResponseWriter, request *http.Request) {
	product, err := services.NewProductService().Detail(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(product.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json response")
		return
	}
}

func (productController *ProductController) CountProduct(response http.ResponseWriter, request *http.Request) {
	product, err := services.NewProductService().CountProduct(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(product.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json response")
		return
	}
}

func (productController *ProductController) Update(response http.ResponseWriter, request *http.Request) {
	product, err := services.NewProductService().Update(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(product.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json response")
		return
	}
}

func (productcontroller ProductController) Delete(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewProductService().Delete(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(result.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json response")
		return
	}
}
