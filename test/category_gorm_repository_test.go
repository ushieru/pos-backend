package test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/repository"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestEmptyListCategory(t *testing.T) {
	cs, clean := Bootstrap()
	t.Cleanup(func() {
		clean()
	})
	want := 0

	categories, err := cs.List()
	categoriesLength := len(categories)

	if err != nil {
		t.Fatalf(err.Message)
	}
	if categoriesLength != want {
		t.Fatalf("Categories length: %d, want: %d", categoriesLength, want)
	}
}

func TestCreateCategory(t *testing.T) {
	cs, clean := Bootstrap()
	t.Cleanup(func() {
		clean()
	})
	want := 1

	createCategoryDTO := &dto.UpsertCategoryRequest{
		Name: "Category Test",
	}

	category, err := cs.Save(createCategoryDTO)

	if err != nil {
		t.Fatalf(err.Message)
	}
	if int(category.ID) != want {
		t.Fatalf("Category id: %d, want: %d", category.ID, want)
	}
}

func TestUpdateCategoryFail(t *testing.T) {
	cs, clean := Bootstrap()
	t.Cleanup(func() {
		clean()
	})
	want := "Categoria no encontrada"

	_, err := cs.Update(1, &dto.UpsertCategoryRequest{
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
	cs, clean := Bootstrap()
	t.Cleanup(func() {
		clean()
	})
	want := "Categoria no encontrada"

	_, err := cs.Delete(0)

	if err == nil {
		t.Fatalf("Error is empty")
	}
	if err.Message != want {
		t.Fatalf("Error: %s, want: %s", err.Message, want)
	}
}

func Bootstrap() (*service.CategoryService, func() error) {
	database, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	categoryRepository := repository.NewCategoryGormRepository(database)
	categoryService := service.NewCategoryService(categoryRepository)
	db, _ := database.DB()
	return categoryService, db.Close
}