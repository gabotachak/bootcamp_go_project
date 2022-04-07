package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	if s.Exists(ctx, w.WarehouseCode) {
		return -1, errors.New("ya existe un elemento con ese warehouse code")
	}
	return s.repository.Save(ctx, w)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, w domain.Warehouse) error {
	if s.Exists(ctx, w.WarehouseCode) {
		warehouse, err := s.Get(ctx, w.ID)
		if err != nil {
			return errors.New("item inexistente")
		}
		if warehouse.WarehouseCode != w.WarehouseCode {
			return errors.New("ya existe un elemento con ese warehouse code")
		}
	}
	return s.repository.Update(ctx, w)
}

func (s *Service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Exists(ctx context.Context, warehouseCode string) bool {
	return s.repository.Exists(ctx, warehouseCode)
}
