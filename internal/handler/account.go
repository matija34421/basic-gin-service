package handler

import (
	accountdto "basic-gin/internal/dto/accountDto"
	"basic-gin/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accService,
	}
}

func (h *AccountHandler) GetAccountsByClientId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id format: "+err.Error(), http.StatusBadRequest)
		return
	}

	accounts, err := h.accountService.GetAccountsByClientId(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, accounts)
}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id format: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.GetAccountById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, account)
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req struct {
		ClientId int `json:"client_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.CreateAccount(ctx, req.ClientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, account)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dto accountdto.UpdateAccountDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.UpdateAccount(ctx, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, account)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
