package services

import (
	"github.com/vadym-98/golang-microservices/mvc/domain"
	"github.com/vadym-98/golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct{}

var ItemsService itemsService

func (i *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
