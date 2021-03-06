package controllers

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	"github.com/aceslick911/bcg-checkout/pkg/crypto"
	http_err "github.com/aceslick911/bcg-checkout/pkg/http-err"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		c.JSON(http.StatusOK, token)
	}
}
func GatewayRedirect(c *gin.Context) {

	location := url.URL{Path: "/docs/index.html"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
