package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Seller) error
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

func (s *Service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (domain.Seller, error) {
	seller, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Seller{}, err
	}
	return seller, nil
}

func (s *Service) Save(ctx context.Context, sel domain.Seller) (int, error) {
	if s.Exists(ctx, sel.CID) {
		return -1, errors.New("Ya existe un elemento con ese seller cid")
	}
	seller, err := s.repository.Save(ctx, sel)
	if err != nil {
		return 0, err
	}
	return seller, nil
}

func (s *Service) Update(ctx context.Context, sel domain.Seller) error {
	if s.Exists(ctx, sel.CID) {
		return errors.New("No se puede guardar o no existe")
	}
	return s.repository.Update(ctx, sel)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Exists(ctx context.Context, cid int) bool {
	return s.repository.Exists(ctx, cid)
}
