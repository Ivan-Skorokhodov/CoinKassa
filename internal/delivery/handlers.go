package delivery

import (
	"CoinKassa/internal/delivery/response"
	"CoinKassa/internal/models"
	"CoinKassa/internal/usecase"
	"encoding/json"
	"net/http"
)

type Handler struct {
	usecase usecase.UsecaseInterface
}

func NewHandler(usecase usecase.UsecaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.SendErrorResponse("Method not allowed", http.StatusMethodNotAllowed, w)
		return
	}

	var inputData models.StoreRegisterInput
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		response.SendErrorResponse("Input data is not acceptable", http.StatusBadRequest, w)
		return
	}

	// TODO: validate data

	err := h.usecase.RegisterStore(inputData)
	if err != nil {
		response.SendErrorResponse(err.Error(), http.StatusBadRequest, w)
		return
	}

	response.SendOKResponse(w)

}
