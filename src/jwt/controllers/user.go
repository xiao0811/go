package controllers

import (
	"encoding/json"
	"jwt/helper"
	"jwt/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code:http.StatusBadRequest, Msg:"bad params"})
		return
	}
}
