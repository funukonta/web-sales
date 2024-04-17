package handlers

import (
	"net/http"
	"web-sales/pkg"
	"web-sales/web/models"
	"web-sales/web/services"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request) error
	GetAll(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	serv services.UserService
}

func NewUserHandler(serv services.UserService) UserHandler {
	return &userHandler{serv: serv}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) error {
	user := new(models.RegisterUser)
	if err := pkg.GetJsonBody(r, user); err != nil {
		return err
	}

	newUser, err := h.serv.Create(user)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: newUser, Message: "User Created!"}).Send(w)
	return nil
}

func (h *userHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	users, err := h.serv.GetAll()
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: users, Message: "Berhasil ambil data Users"}).Send(w)
	return nil
}
