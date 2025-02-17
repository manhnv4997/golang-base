package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type CustomerController struct {
	Controllers
	CustomerService *services.CustomerService
}

func NewCustomerController(customerService *services.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

func (customerController *CustomerController) List(response http.ResponseWriter, request *http.Request) {

	customers, err := customerController.CustomerService.List(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(customers.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) Detail(response http.ResponseWriter, request *http.Request) {
	customer, err := customerController.CustomerService.Detail(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(customer.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) CustomerOrders(response http.ResponseWriter, request *http.Request) {
	customerOrders, err := customerController.CustomerService.CustomerOrders(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(customerOrders.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) Count(response http.ResponseWriter, request *http.Request) {
	count, err := customerController.CustomerService.Count(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(count.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) Search(response http.ResponseWriter, request *http.Request) {
	count, err := customerController.CustomerService.Search(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(count.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) Store(response http.ResponseWriter, request *http.Request) {

	users, err := customerController.CustomerService.Store(response, request)
	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	encodeErr := utils.SuccessResponse(response, http.StatusOK, users)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (customerController *CustomerController) Update(response http.ResponseWriter, request *http.Request) {

	users, err := customerController.CustomerService.Update(response, request)
	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(users.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}
