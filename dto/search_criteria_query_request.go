package dto

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/domain/criteria"
)

type SearchCriteriaQueryRequest struct {
	Filters []domain_criteria.Filter `query:"filters"`
}

func (dto *SearchCriteriaQueryRequest) Validate() *domain.AppError {
	return nil
}
