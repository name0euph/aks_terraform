package handler

import (
	"backend-namecard/usecase"
)

type INamecardHandler interface {

}

type namecardHandler struct {
	nu usecase.INamecardUsecase
}

func NewNamecardHandler(nu usecase.INamecardUsecase) INamecardHandler {
	return &namecardHandler{nu}
}

