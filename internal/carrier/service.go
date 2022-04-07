package carrier

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type ServiceInterface interface {
	GetAllReport(ctx context.Context) ([]CarriersByLocality, error)
	GetReportDetails(ctx context.Context, idLocality int) ([]domain.Carrier, error)
	GetReport(ctx context.Context, idLocality int) (CarriersByLocality, error)
	LocalityExists(ctx context.Context, id int) bool
	CIDExists(ctx context.Context, cid string) bool
	Store(ctx context.Context, w domain.Carrier) (domain.Carrier, error)
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
Methods in the business layer.
In the service we communicate with the rest of the services and the carrier repository
**/

func (s *Service) GetAllReport(ctx context.Context) ([]CarriersByLocality, error) {
	return s.repository.GetAllReport(ctx)
}

func (s *Service) GetReportDetails(ctx context.Context, id int) ([]domain.Carrier, error) {
	if s.LocalityExists(ctx, id) {
		return s.repository.GetReportDetails(ctx, id)
	} else {
		return []domain.Carrier{}, errors.New("no existe una localidad con ese id")
	}
}

func (s *Service) GetReport(ctx context.Context, id int) (CarriersByLocality, error) {
	if s.LocalityExists(ctx, id) {
		return s.repository.GetReport(ctx, id)
	} else {
		return CarriersByLocality{}, errors.New("no existe una localidad con ese id")
	}
}

func (s *Service) Store(ctx context.Context, c domain.Carrier) (domain.Carrier, error) {
	if !s.LocalityExists(ctx, c.LocalityId) {
		return domain.Carrier{}, errors.New("no existe una localidad con ese id")
	}
	if s.CIDExists(ctx, c.CID) {
		return domain.Carrier{}, errors.New("ya existe un transportista con cid igual")
	}
	return s.repository.Store(ctx, c)

}

func (s *Service) LocalityExists(ctx context.Context, id int) bool {
	return s.repository.LocalityExists(ctx, id)
}

func (s *Service) CIDExists(ctx context.Context, cid string) bool {
	return s.repository.CIDExists(ctx, cid)
}
