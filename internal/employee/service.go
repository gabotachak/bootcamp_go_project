package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, cardNumberID string) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, cardNumberID string) error
	GetInboundOrders(ctx context.Context) ([]employeeInboundOrders, error)
	GetInboundOrdersByEmployee(ctx context.Context, id int) (employeeInboundOrders, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, cardNumberID string) (domain.Employee, error) {
	return s.repository.Get(ctx, cardNumberID)
}

func (s *Service) Exists(ctx context.Context, cardNumberID string) bool {
	return s.repository.Exists(ctx, cardNumberID)
}

func (s *Service) Save(ctx context.Context, e domain.Employee) (int, error) {

	if s.Exists(ctx, e.CardNumberID) {
		return 0, errors.New("el card_number_id ya existe")

	}
	return s.repository.Save(ctx, e)

}
func (s *Service) Update(ctx context.Context, e domain.Employee) error {
	if s.Exists(ctx, e.CardNumberID) {
		return s.repository.Update(ctx, e)

	}
	return errors.New("el card_number_id no existe")

}
func (s *Service) Delete(ctx context.Context, cardNumberID string) error {
	return s.repository.Delete(ctx, cardNumberID)
}
func (s *Service) GetInboundOrders(ctx context.Context) ([]employeeInboundOrders, error) {
	return s.repository.GetInboundOrders(ctx)
}
func (s *Service) GetInboundOrdersByEmployee(ctx context.Context, id int) (employeeInboundOrders, error) {
	return s.repository.GetInboundOrdersByEmployee(ctx, id)
}
