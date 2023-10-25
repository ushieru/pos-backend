package repository

import (
	"fmt"

	"github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type CriteriaGormRepository struct{}

func (r *CriteriaGormRepository) FiltersToScopes(filters []domain_criteria.Filter) []func(*gorm.DB) *gorm.DB {
	scopes := make([]func(*gorm.DB) *gorm.DB, 0)
	for _, filter := range filters {
		whereStatement := fmt.Sprintf("%s %s %s", filter.Field, filter.Operator, filter.Value)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(whereStatement)
		})
	}
	return scopes
}
