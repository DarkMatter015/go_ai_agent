package handlers

import (
	"agent/cmd/internal/models"
	"encoding/json"
	"net/http"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /users", h.addUser)
	http.HandleFunc("PUT /users/{id}", h.updateUser)
	http.HandleFunc("DELETE /users/{id}", h.deleteUser)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.usecases.GetAllUsers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {

	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := h.usecases.AddNewUser(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.UserResponse{ID: id})

}

func (h Handlers) updateUser(w http.ResponseWriter, r *http.Request) {

	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id := r.PathValue("id")
	if id == "" || id == "/" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "ID is required"})
		return
	}

	user, err := h.usecases.UpdateUser(req, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h Handlers) deleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" || id == "/" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "ID is required"})
		return
	}

	user, err := h.usecases.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(user)

}
