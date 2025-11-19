package delivery

import (
	"CoinKassa/internal/delivery/response"
	"CoinKassa/internal/models"
	"CoinKassa/internal/usecase"
	"CoinKassa/pkg/logs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	usecase   usecase.UsecaseInterface
	validator *validator.Validate
}

func NewHandler(usecase usecase.UsecaseInterface) *Handler {
	return &Handler{
		usecase:   usecase,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *Handler) RegisterStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logs.PrintLog(r.Context(), "[delivery] RegisterStore", "Method not allowed")
		response.SendErrorResponse("Method not allowed", http.StatusMethodNotAllowed, w)
		return
	}

	var inputData models.StoreRegisterInput
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		logs.PrintLog(r.Context(), "[delivery] RegisterStore", "Input data is not acceptable")
		response.SendErrorResponse("Input data is not acceptable", http.StatusBadRequest, w)
		return
	}

	err := h.validator.Struct(inputData)
	if err != nil {
		logs.PrintLog(r.Context(), "[delivery] RegisterStore", fmt.Sprintf("Validation error: %s", err.Error()))
		response.SendErrorResponse(err.Error(), http.StatusBadRequest, w)
		return
	}

	err = h.usecase.RegisterStore(r.Context(), inputData)
	if err != nil {
		logs.PrintLog(r.Context(), "[delivery] RegisterStore", err.Error())
		response.SendErrorResponse(err.Error(), http.StatusBadRequest, w)
		return
	}

	response.SendOKResponse(w)
}
