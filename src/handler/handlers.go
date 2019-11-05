package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/akspokl/go-task/src/db"
	"github.com/akspokl/go-task/src/model"
	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

var (
	commonCache = cache.New(5*time.Minute, 30*time.Second)
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	Modify(user, hashedPassword)

	if err := db.AddUser(user); err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}

	foundUser, err := db.GetUserByUsername(user.Username)

	if err != nil {
		HandleError(err, w, http.StatusNotFound)
		return
	}

	token := GenerateToken(foundUser)
	CacheToken(foundUser, token)

	json.NewEncoder(w).Encode(token)
	return
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := db.GetAllUsers()

	if err != nil {
		HandleError(err, w, http.StatusNotFound)
		return
	}

	response := MapToUserResponse(users)

	json.NewEncoder(w).Encode(response)
	return
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	username, found := commonCache.Get(token)

	if found {
		fmt.Print(username)
		json.NewEncoder(w).Encode(username)
	} else {
		w.WriteHeader(http.StatusNotFound)
		log.Println(username)
	}
	return
}

func GetGithubEvents(w http.ResponseWriter, _ *http.Request) {
	client := github.NewClient(nil)
	orgs, _, err := client.Activity.ListEvents(context.Background(), nil)
	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orgs)
	return
}

func GenerateToken(user *model.User) model.TokenResponse {
	buff := make([]byte, int(math.Round(float64(23)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return model.TokenResponse{Username: user.Username, Token: str}
}

func CacheToken(user *model.User, token model.TokenResponse) {
	commonCache.Set(token.Token, user.Username, cache.DefaultExpiration)
}

func MapToUserResponse(users []model.User) []model.UserResponse {
	response := []model.UserResponse{}
	for _, user := range users {
		n := model.UserResponse{Username: user.Username}
		response = append(response, n)
	}
	return response
}

func Modify(user model.User, password []byte) {
	user.Password = base64.StdEncoding.EncodeToString(password)
}
