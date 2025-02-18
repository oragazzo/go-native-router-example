package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oragazzo/go-native-router-example/internal/models"
	"github.com/oragazzo/go-native-router-example/internal/storage"

	"github.com/google/uuid"
)

type UserHandler struct {
	storage *storage.MemoryStorage
}

func NewUserHandler(storage *storage.MemoryStorage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.storage.GetUsers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserList{Users: users})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.ID = uuid.New().String()
	user = h.storage.AddUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := h.storage.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.storage.UpdateUser(id, updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.storage.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
