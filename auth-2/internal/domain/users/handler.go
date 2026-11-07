package users

import (
	"auth-two/internal/domain/users/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
    // Set response content type.
    w.Header().Set("Content-Type", "application/json")

    // Create request DTO.
    var req dto.UserCreateRequest

    // Decode JSON request body.
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"message":"Invalid JSON body"}`, http.StatusBadRequest)
        return
    }

	fmt.Println("json decoded data ->", json.NewDecoder(r.Body))

    // Validate required fields.
    if req.Name == "" || req.Email == "" || req.Password == "" {
        http.Error(w, `{"message":"Name, email and password are required"}`, http.StatusBadRequest)
        return
    }

    // Call service layer.
	// fmt.Println("users ->", req)

    // Send successful response.
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "User created successfully",
		"data": req,
    })
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all users retrived successfull"))
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("single user retrived successfull"))
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updated successfull users"))
}

func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user updated successfull"))
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user deleted successfull"))
}