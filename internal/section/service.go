package section

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id interface{}) (domain.Section, error)
	Exists(ctx context.Context, sn string) bool
	Save(ctx context.Context, d domain.Section) (int, error)
	Update(ctx context.Context, d domain.Section) error
	Delete(ctx context.Context, id interface{}) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Section, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Save(ctx context.Context, d domain.Section) (int, error) {
	if s.Exists(ctx, d.SectionNumber) {
		return -1, errors.New("SectionNumber already exists")
	}

	idSection, err := s.repository.Save(ctx, d)
	if err != nil {
		return idSection, err
	}
	return idSection, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, d domain.Section) error {
	if s.Exists(ctx, d.SectionNumber) {
		section, err := s.Get(ctx, d.ID)
		if err != nil {
			return errors.New("item not found")
		}
		if section.SectionNumber != d.SectionNumber {
			return errors.New("section number already used")
		}
	}
	return s.repository.Update(ctx, d)
}

func (s *Service) Get(ctx context.Context, id int) (domain.Section, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Exists(ctx context.Context, sn string) bool {
	return s.repository.Exists(ctx, sn)
}
