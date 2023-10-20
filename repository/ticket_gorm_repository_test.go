package repository

import (
	"testing"

	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
)

func TestTicketGormRepository(t *testing.T) {
	t.Run("List tickets", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, _ := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		tickets, err := s.List()
		if err != nil {
			t.Fatalf(err.Message)
		}
		if len(tickets) != 0 {
			t.Fatalf("Ticket is not empty")
		}
	})

	t.Run("Create New ticket", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, _ := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		ticket, err := s.Save(&domain.Account{})
		if err != nil {
			t.Fatalf(err.Message)
		}
		if ticket == nil || ticket.ID == 0 {
			t.Fatalf("Ticket is not empty")
		}
	})

	t.Run("Find ticket", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, _ := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		s.Save(&domain.Account{})
		ticket, err := s.Find(1)
		if err != nil {
			t.Fatalf(err.Message)
		}
		if ticket == nil || ticket.ID == 0 {
			t.Fatalf("Ticket is not empty")
		}
	})

	t.Run("Delete ticket", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, _ := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		s.Save(&domain.Account{})
		s.Delete(1)
		ticket, err := s.Find(1)
		if ticket != nil {
			t.Fatalf("Ticket found")
		}
		if err == nil {
			t.Fatalf("Ticket found")
		}
	})

	t.Run("Delete ticket that does not exist", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, _ := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		ticket, deleteErr := s.Delete(1)
		if deleteErr == nil {
			t.Fatalf("Error not found")
		}
		if ticket != nil {
			t.Fatalf("Ticket found")
		}
	})

	t.Run("Add product", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, ps := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		a := &domain.Account{Model: domain.Model{ID: 1}}
		s.Save(a)
		ps.Save(&dto.UpsertProductRequest{Name: "ptest", Description: "dtest", Price: 1})
		_, err := s.AddProduct(1, 1, a)
		ticket, fErr := s.Find(1)
		if err != nil {
			t.Fatalf(err.Message)
		}
		if fErr != nil {
			t.Fatalf(fErr.Message)
		}
		if ticket == nil {
			t.Fatalf("Ticket nil")
		}
		if len(ticket.TicketProducts) == 0 {
			t.Fatalf("Ticket empty")
		}
	})

	t.Run("Delete product", func(t *testing.T) {
		db, clean := getInMemoryDB()
		s, ps := getTicketGormRepository(db)
		t.Cleanup(func() {
			clean()
		})
		a := &domain.Account{Model: domain.Model{ID: 1}}
		s.Save(a)
		ps.Save(&dto.UpsertProductRequest{Name: "ptest", Description: "dtest", Price: 1})
		_, err := s.AddProduct(1, 1, a)
		ticket, fErr := s.Find(1)
		if err != nil {
			t.Fatalf(err.Message)
		}
		if fErr != nil {
			t.Fatalf(fErr.Message)
		}
		if ticket == nil {
			t.Fatalf("Ticket nil")
		}
		if len(ticket.TicketProducts) == 0 {
			t.Fatalf("Ticket empty")
		}
	})
}

func getTicketGormRepository(db *gorm.DB) (*service.TicketService, *service.ProductService) {
	pr := NewProductGormRepository(db)
	ps := service.NewProductService(pr)
	r := NewTicketGormRepository(db, ps)
	return service.NewTicketService(r), ps
}
