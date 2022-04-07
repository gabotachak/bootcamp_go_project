package inboundOrders

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Inbound_Orders, error)
	Get(ctx context.Context, id int) (domain.Inbound_Orders, error)
	Exists(ctx context.Context, id int) bool
	Save(ctx context.Context, io domain.Inbound_Orders) (domain.Inbound_Orders, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Inbound_Orders, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (domain.Inbound_Orders, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Exists(ctx context.Context, id int) bool {
	return s.repository.Exists(ctx, id)
}

func (s *Service) Save(ctx context.Context, io domain.Inbound_Orders) (domain.Inbound_Orders, error) {

	if s.Exists(ctx, io.ID) {
		return domain.Inbound_Orders{}, errors.New("el id ya existe")

	}
	return s.repository.Save(ctx, io)

}
