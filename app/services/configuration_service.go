package services

import (
	"log"
)

type ConfigurationService struct{}

func NewConfigurationService() *ConfigurationService {
	return &ConfigurationService{}
}

func (configurationService *ConfigurationService) Handle() (interface{}, error) {
	log.Println("OK....!!!")

	return "Xử lý thành công", nil
}
