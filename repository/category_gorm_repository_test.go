package repository

import (
	"testing"

	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
)

func TestEmptyListCategory(t *testing.T) {
	db, clean := getInMemoryDB()
	s := getCategoryGormRepository(db)
	t.Cleanup(func() {
		clean()
	})
	want := 0

	categories, err := s.List(&dto.SearchCriteriaQueryRequest{})
	categoriesLength := len(categories)

	if err != nil {
		t.Fatalf(err.Message)
	}
	if categoriesLength != want {
		t.Fatalf("Categories length: %d, want: %d", categoriesLength, want)
	}
}

func TestCreateCategory(t *testing.T) {
	db, clean := getInMemoryDB()
	s := getCategoryGormRepository(db)
	t.Cleanup(func() {
		clean()
	})
	want := 1

	createCategoryDTO := &dto.UpsertCategoryRequest{
		Name: "Category Test",
	}

	category, err := s.Save(createCategoryDTO)

	if err != nil {
		t.Fatalf(err.Message)
	}
	if int(category.ID) != want {
		t.Fatalf("Category id: %d, want: %d", category.ID, want)
	}
}

func TestUpdateCategoryFail(t *testing.T) {
	db, clean := getInMemoryDB()
	s := getCategoryGormRepository(db)
	t.Cleanup(func() {
		clean()
	})
	want := "Categoria no encontrada"

	_, err := s.Update(1, &dto.UpsertCategoryRequest{
		Name: "category test",
	})

	if err == nil {
		t.Fatalf("Error is empty")
	}
	if err.Message != want {
		t.Fatalf("Error: %s, want: %s", err.Message, want)
	}
}

func TestDeleteCategoryFail(t *testing.T) {
	db, clean := getInMemoryDB()
	s := getCategoryGormRepository(db)
	t.Cleanup(func() {
		clean()
	})
	want := "Categoria no encontrada"

	_, err := s.Delete(0)

	if err == nil {
		t.Fatalf("Error is empty")
	}
	if err.Message != want {
		t.Fatalf("Error: %s, want: %s", err.Message, want)
	}
}

func getCategoryGormRepository(db *gorm.DB) *service.CategoryService {
	r := NewCategoryGormRepository(db)
	return service.NewCategoryService(r)
}
