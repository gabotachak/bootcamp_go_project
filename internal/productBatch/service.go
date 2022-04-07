package productBatch

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	Save(ctx context.Context, pb domain.ProductBatch) (int, error)
	ReportAll(ctx context.Context) ([]domain.ReportProductBatch, error)
	ReportBySection(ctx context.Context, sectionId int) (domain.ReportProductBatch, error)
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) Save(ctx context.Context, pb domain.ProductBatch) (int, error) {
	return s.repository.Save(ctx, pb)
}

func (s *Service) ReportAll(ctx context.Context) ([]domain.ReportProductBatch, error) {
	return s.repository.ReportAll(ctx)
}

func (s *Service) ReportBySection(ctx context.Context, sectionId int) (domain.ReportProductBatch, error) {
	if s.repository.SectionExists(ctx, sectionId) {
		return s.repository.ReportBySection(ctx, sectionId)
	}
	return domain.ReportProductBatch{}, errors.New("section id does not exist")
}
