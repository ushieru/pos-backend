package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ushieru/pos/app/fiber/middlewares"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/service"
)

type ProductionCenterHandler struct {
	service    service.IProductionCenterService
	middleware *middlewares.AuthMiddleware
}

func (h *ProductionCenterHandler) setupRoutes(app *fiber.App) {
	productionCenter := app.Group("/api/production-centers")
	productionCenter.Use(h.middleware.CheckJWT)
	productionCenter.Get("/", h.listProductionCenters)
	productionCenter.Get("/:id", h.findProductionCenter)
	productionCenter.Get("/:id/tickets/:ticketId", h.findTicketByProductionCenter)
	productionCenter.Post("/", h.saveProductionCenter)
	productionCenter.Put("/:id", h.updateProductionCenter)
	productionCenter.Post("/:id/accounts/:accountId", h.addAccount)
	productionCenter.Post("/:id/categories/:categoryId", h.addCategory)
	productionCenter.Delete("/:id/accounts/:accountId", h.deleteAccount)
	productionCenter.Delete("/:id/categories/:categoryId", h.deleteCategory)
	productionCenter.Delete("/:id", h.deleteProductionCenter)
}

// @Router /api/production-centers [GET]
// @Security ApiKeyAuth
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {array} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) listProductionCenters(c *fiber.Ctx) error {
	searchCriteriaQueryRequest := new(dto.SearchCriteriaQueryRequest)
	if err := c.QueryParser(searchCriteriaQueryRequest); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Error parse query")
	}
	productionCenter, err := h.service.List(searchCriteriaQueryRequest)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id} [GET]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) findProductionCenter(c *fiber.Ctx) error {
	id := c.Params("id")
	productionCenters, err := h.service.Find(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenters)
}

// @Router /api/production-centers/{id}/tickets/{ticketId} [GET]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param ticketId path string true "Ticket ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.Ticket
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) findTicketByProductionCenter(c *fiber.Ctx) error {
	id := c.Params("id")
	ticketId := c.Params("ticketId")
	productionCenters, err := h.service.GetTicket(id, ticketId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenters)
}

// @Router /api/production-centers [POST]
// @Security ApiKeyAuth
// @Param dto body dto.CreateProductionCenterRequest true "Create ProductionCenter dto"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) saveProductionCenter(c *fiber.Ctx) error {
	dto := new(dto.CreateProductionCenterRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	productionCenter, err := h.service.Save(dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id} [PUT]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param dto body dto.CreateProductionCenterRequest true "Create ProductionCenter dto"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) updateProductionCenter(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(dto.CreateProductionCenterRequest)
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	productionCenter, err := h.service.Update(id, dto)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id}/accounts/{accountId} [POST]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param accountId path string true "Account ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) addAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	accountId := c.Params("accountId")
	productionCenter, err := h.service.AddAccount(id, accountId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id}/accounts/{accountId} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param accountId path string true "Account ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) deleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	accountId := c.Params("accountId")
	productionCenter, err := h.service.DeleteAccount(id, accountId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id}/categories/{categoryId} [POST]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param categoryId path string true "Category ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) addCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categoryId := c.Params("categoryId")
	productionCenter, err := h.service.AddCategory(id, categoryId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id}/categories/{categoryId} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Param categoryId path string true "Category ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) deleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	categoryId := c.Params("categoryId")
	productionCenter, err := h.service.DeleteCategory(id, categoryId)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

// @Router /api/production-centers/{id} [DELETE]
// @Security ApiKeyAuth
// @Param id path string true "Production Center ID"
// @Tags Production center
// @Accepts json
// @Produce json
// @Success 200 {object} domain.ProductionCenter
// @Failure default {object} domain.AppError
func (h *ProductionCenterHandler) deleteProductionCenter(c *fiber.Ctx) error {
	id := c.Params("id")
	productionCenter, err := h.service.Delete(id)
	if err != nil {
		return fiber.NewError(err.Code, err.Message)
	}
	return c.JSON(productionCenter)
}

func NewProductionCenterHandler(service service.IProductionCenterService, middleware *middlewares.AuthMiddleware, app *fiber.App) *ProductionCenterHandler {
	ch := ProductionCenterHandler{service, middleware}
	ch.setupRoutes(app)
	return &ch
}
