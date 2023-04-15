package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"

	"jwt-authentication/database"
	"jwt-authentication/models"
)

var (
	Ctx    context.Context
	DbRepo database.Repository
)

func DbConn(ctx context.Context, dbRepo database.Repository) {
	Ctx = ctx
	DbRepo = dbRepo
}

func AddTable(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.RunMigrate(Ctx, DbRepo)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Table created")
}

func DeleteTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := database.RunDeleteTable(Ctx, DbRepo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Table deleted")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user, err := database.RunAdd(Ctx, DbRepo, models.User1)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Println(user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := database.RunDelete(Ctx, DbRepo, "dwiw")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Your Account Is Deleted")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	allUser := database.RunGetAll(Ctx, DbRepo)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(allUser)
}
