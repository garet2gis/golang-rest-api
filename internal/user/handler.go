package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/internal/apperror"
	"golang-rest-api/internal/handlers"
	"golang-rest-api/pkg/logging"
	"net/http"
	"strings"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct {
	service *Service
	logger  *logging.Logger
}

func NewHandler(storage Storage, logger *logging.Logger) handlers.Handler {
	service := NewService(storage, logger)
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, _ *http.Request) error {
	users, err := h.service.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal users, users: %+v", users)
	}
	w.Write(response)

	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	user, err := h.service.GetOneUser(context.Background(), id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user, user: %+v", user)
	}
	w.Write(response)

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var u CreateUserDTO

	h.logger.Debug("decoding new user entity")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return fmt.Errorf("failed to decode new user info")
	}

	user, err := h.service.CreateUser(context.Background(), u)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user, user: %+v", user)
	}
	w.Write(response)

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	var u UpdateUserDTO

	h.logger.Debug("decoding new user entity")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return fmt.Errorf("failed to decode new user info")
	}

	user, err := h.service.UpdateUser(context.Background(), id, u)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user, user: %+v", user)
	}
	w.Write(response)

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	err := h.service.DeleteUser(context.Background(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
