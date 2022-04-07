package locality

import (
	"context"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetReport(ctx context.Context, id int) (LocalityReport, error)
	GetGeneralReport(ctx context.Context) ([]LocalityReport, error)
	Save(ctx context.Context, loc domain.Locality) (domain.Locality, error)
}
type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

/**
* Methods in the business layer.
* In the service we communicate with the rest of the services and the carrier repository
**/
func (s *Service) GetGeneralReport(ctx context.Context) ([]LocalityReport, error) {
	return s.repository.GetGeneralReport(ctx)
}

func (s *Service) GetReport(ctx context.Context, id int) (LocalityReport, error) {
	report, err := s.repository.GetReport(ctx, id)
	if err != nil {
		return LocalityReport{}, err
	}
	return report, nil
}

func (s *Service) Save(ctx context.Context, loc domain.Locality) (domain.Locality, error) {

	locality, err := s.repository.Save(ctx, loc)
	if err != nil {
		return domain.Locality{}, err
	}
	return locality, nil
}
