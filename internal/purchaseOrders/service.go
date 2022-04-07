package purchaseOrders

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.PurchaseOrder, error)
	Get(ctx context.Context, purchaseOrderId int) (domain.PurchaseOrder, error)
	Exists(ctx context.Context, purchaseOrderId int) bool
	Store(ctx context.Context, p domain.PurchaseOrder) (domain.PurchaseOrder, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.PurchaseOrder, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, purchaseOrderId int) (domain.PurchaseOrder, error) {
	return s.repository.Get(ctx, purchaseOrderId)
}

func (s *Service) Exists(ctx context.Context, purchaseOrderId int) bool {
	return s.repository.Exists(ctx, purchaseOrderId)
}

func (s *Service) Store(ctx context.Context, p domain.PurchaseOrder) (domain.PurchaseOrder, error) {

	if s.Exists(ctx, p.ID) {
		return domain.PurchaseOrder{}, errors.New("la orden de compra ya existe")

	}
	return s.repository.
		Store(ctx, p)

}
