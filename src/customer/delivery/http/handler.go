package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/MCABankCustomer/src/customer/domain"
	"github.com/rasteiro11/MCABankCustomer/src/customer/service"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogCore/pkg/validator"
)

var CustomerGroupPath = "/customers"

type (
	HandlerOpt func(*handler)
	handler    struct {
		customerService service.CustomerService
	}
)

func WithCustomerService(s service.CustomerService) HandlerOpt {
	return func(h *handler) {
		h.customerService = s
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("", CustomerGroupPath, http.MethodGet, h.FindAll)
	server.AddHandler("/:id", CustomerGroupPath, http.MethodGet, h.FindByID)
	server.AddHandler("", CustomerGroupPath, http.MethodPost, h.Create)
	server.AddHandler("/:id", CustomerGroupPath, http.MethodPut, h.Update)
	server.AddHandler("/:id", CustomerGroupPath, http.MethodDelete, h.Delete)
}

var _ Handler = (*handler)(nil)

var (
	ErrPathParam     = errors.New("path param is missing")
	ErrTypeAssertion = errors.New("type assertion error")
)

// FindAll godoc
// @Summary Get all customers
// @Description Retrieve a list of all customers
// @Tags customers
// @Produce json
// @Success 200 {array} customerResponse
// @Failure 500 {object} any
// @Router /customers [get]
func (h *handler) FindAll(c *fiber.Ctx) error {
	customers, err := h.customerService.GetAll(c.Context())
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}
	return rest.NewStatusOk(c, rest.WithBody(MapCustomersToHTTP(customers)))
}

// FindByID godoc
// @Summary Get a customer by ID
// @Description Retrieve a single customer by its ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} customerResponse
// @Failure 400 {object} any
// @Failure 404 {object} any
// @Router /customers/{id} [get]
func (h *handler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	customer, err := h.customerService.GetByID(c.Context(), uint(id))
	if err != nil {
		return rest.NewStatusNotFound(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(MapCustomerToHTTP(customer)))
}

// Create godoc
// @Summary Create a new customer
// @Description Create a customer with a name and email
// @Tags customers
// @Accept json
// @Produce json
// @Param request body createCustomerRequest true "Customer info"
// @Success 201 {object} customerResponse
// @Failure 400 {object} any
// @Failure 422 {object} any
// @Router /customers [post]
func (h *handler) Create(c *fiber.Ctx) error {
	req := &createCustomerRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if _, err := validator.IsRequestValid(req); err != nil {
		return rest.NewResponse(c, http.StatusBadRequest, rest.WithBody(err)).JSON(c)
	}

	customer, err := h.customerService.Create(c.Context(), MapCreateRequestToDomain(req))
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(MapCustomerToHTTP(customer)))
}

// Update godoc
// @Summary Update an existing customer
// @Description Update a customer's name and email by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param request body updateCustomerRequest true "Updated customer info"
// @Success 200 {object} customerResponse
// @Failure 400 {object} any
// @Failure 422 {object} any
// @Router /customers/{id} [put]
func (h *handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	req := &updateCustomerRequest{}
	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if _, err := validator.IsRequestValid(req); err != nil {
		return rest.NewResponse(c, http.StatusBadRequest, rest.WithBody(err)).JSON(c)
	}

	customer, err := h.customerService.Update(c.Context(), &domain.Customer{
		ID:    uint(id),
		Nome:  req.Nome,
		Email: req.Email,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(MapCustomerToHTTP(customer)))
}

// Delete godoc
// @Summary Delete a customer
// @Description Delete a customer by ID
// @Tags customers
// @Param id path int true "Customer ID"
// @Success 200 {string} string "Customer deleted successfully"
// @Failure 400 {object} any
// @Failure 422 {object} any
// @Router /customers/{id} [delete]
func (h *handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if err := h.customerService.Delete(c.Context(), uint(id)); err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody("Customer deleted successfully"))
}
