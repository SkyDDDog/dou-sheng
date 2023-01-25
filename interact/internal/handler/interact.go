package handler

import "interact/internal/service"

type InteractService struct {
	service.UnimplementedInteractServiceServer
}

func NewInteractService() *InteractService {
	return &InteractService{}
}
