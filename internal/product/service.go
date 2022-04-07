package product

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, w domain.Product) (int, error)
	Update(ctx context.Context, w domain.Product) error
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

func (s *Service) GetAll(ctx *gin.Context) ([]domain.Product, error) {
	productos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s *Service) Get(ctx *gin.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Save(ctx context.Context, req domain.Product) (int, error) {
	if s.Exists(ctx, req.ProductCode) {
		return -1, errors.New("no se puede guardar el producto o el codigo de producto ya existe")
	}
	return s.repository.Save(ctx, req)

}

func (s *Service) Update(ctx *gin.Context, p domain.Product) error {
	if s.Exists(ctx, p.ProductCode) {
		prod, err := s.Get(ctx, p.ID)
		if err != nil {
			return errors.New("el producto no existe")
		}
		if prod.ProductCode != p.ProductCode {
			return errors.New("el product Code ya existe en otro producto")
		}
	}

	return s.repository.Update(ctx, p)

}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *Service) Exists(ctx context.Context, productCode string) bool {
	return s.repository.Exists(ctx, productCode)
}
