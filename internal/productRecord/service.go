package productRecord

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetAll(ctx *gin.Context) ([]ProductRecordByProduct, error)
	Get(ctx context.Context, id int) (ProductRecordByProduct, error)
	Save(ctx context.Context, w domain.ProductRecord) (ProdRecSave, error)
	Exist(ctx context.Context, w domain.ProductRecord) bool
}
type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

// Send to the repository the GetAll request
func (s *Service) GetAll(ctx *gin.Context) ([]ProductRecordByProduct, error) {
	return s.repository.GetAll(ctx)
}

// Send to the repository the Get request for a specific id
func (s *Service) Get(ctx *gin.Context, id int) (ProductRecordByProduct, error) {
	if s.repository.ExistId(ctx, id) {
		return s.repository.Get(ctx, id)
	} else {
		return ProductRecordByProduct{}, errors.New("el id de producto ingresado no existe")
	}
}

// Send to the repository the Save request
func (s *Service) Save(ctx context.Context, req domain.ProductRecord) (ProdRecSave, error) {

	if !s.repository.ExistId(ctx, req.ProductId) {
		return ProdRecSave{}, errors.New("el id de producto no existe")
	}
	// if !s.repository.ValidateDate(ctx, req.LastUpdateDate){
	// 	return ProdRecSave{}, errors.New("la fecha debe ser mayor a la actual")
	// }
	pr, err := s.repository.Save(ctx, req)
	return pr, err

}
