package repository

import (
	"testing"

	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
)

func TestTableGormRepository(t *testing.T) {

	t.Run("Create New Table", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s := getTableGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		if _, err := s.Save(&dto.CreateTableRequest{Name: "asdqwe"}); err != nil {
			t.Fail()
		}
	})
}

func getTableGormRepository(db *gorm.DB) *service.TableService {
	r := NewTableGormRepository(db)
	return service.NewTableService(r)
}
