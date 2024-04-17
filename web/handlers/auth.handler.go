package handlers

import (
	"net/http"
	"time"
	"web-sales/pkg"
	"web-sales/web/models"
	"web-sales/web/services"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request) error
	Logout(w http.ResponseWriter, r *http.Request) error
}

type authHandler struct {
	serv services.AuthService
}

func NewAuthHandler(serv services.AuthService) AuthHandler {
	return &authHandler{serv: serv}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) error {
	login := new(models.Login)
	if err := pkg.GetJsonBody(r, login); err != nil {
		return err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	token, err := h.serv.Login(login, expirationTime)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	tokenMap := map[string]string{
		"token": token,
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: tokenMap, Message: "Login Sukses"}).Send(w)
	return nil
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Logged Out"}).Send(w)
	return nil
}
