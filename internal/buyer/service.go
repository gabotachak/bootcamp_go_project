package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	Save(ctx context.Context, b domain.Buyer) (int, error)
	GetAll() ([]domain.Buyer, error)
	Get(cardNumberID string) (domain.Buyer, error)
	Update(b domain.Buyer) (domain.Buyer, error)
	Delete(cardNumberID string) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) Save(ctx context.Context, b domain.Buyer) (int, error) {

	if s.Exists(ctx, b.CardNumberID) {
		return -100, errors.New("ya existe un elemento con ese CardNumberID")
	} else {
		idBuyer, err := s.repository.Save(ctx, b)
		if err != nil {
			return idBuyer, err
		}
		return idBuyer, nil
	}

}

func (s *Service) Exists(ctx context.Context, cardNumberID string) bool {
	return s.repository.Exists(ctx, cardNumberID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func (s *Service) Get(ctx context.Context, CardNumberID string) (domain.Buyer, error) {
	return s.repository.Get(ctx, CardNumberID)
}

func (s *Service) Update(ctx context.Context, b domain.Buyer) error {
	if s.Exists(ctx, b.CardNumberID) {
		return s.repository.Update(ctx, b)
	}

	return errors.New("el id no existe en la base de datos")
}

func (s *Service) Delete(ctx context.Context, CardNumberID string) error {
	return s.repository.Delete(ctx, CardNumberID)
}

func (s *Service) GetPurchaseOrders(ctx context.Context) ([]BuyerPurchaseOrders, error) {
	return s.repository.GetPurchaseOrders(ctx)
}
func (s *Service) GetPurchaseOrdersByBuyer(ctx context.Context, id int) (BuyerPurchaseOrders, error) {
	return s.repository.GetPurchaseOrdersByBuyer(ctx, id)
}
