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

func (h *handler) FindAll(c *fiber.Ctx) error {
	customers, err := h.customerService.GetAll(c.Context())
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}
	return rest.NewStatusOk(c, rest.WithBody(MapCustomersToHTTP(customers)))
}

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

func (h *handler) Create(c *fiber.Ctx) error {
	req := &CreateCustomerRequest{}

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

func (h *handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	req := &UpdateCustomerRequest{}
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
