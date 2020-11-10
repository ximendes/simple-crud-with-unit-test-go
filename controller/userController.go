package controller

import (
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "github.com/gorilla/mux"
    "go-project/models"
    "go-project/repository"
    "log"
    "net/http" // used to access the request and response object of the api
    "strconv"
)

type UserController struct {
    repository repository.UserRepositoryInterface
}

func NewUserController(repository repository.UserRepositoryInterface) UserController{
    return UserController{repository}
}

func (controller UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User

    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    insertID := controller.repository.InsertUser(user)

    res := models.Response{
        ID:      insertID,
        Message: "User created successfully",
    }

    json.NewEncoder(w).Encode(res)
}

func (controller UserController) GetUser(w http.ResponseWriter, r *http.Request) {
   params := mux.Vars(r)

   id, err := strconv.Atoi(params["id"])
   fmt.Println(id)

   if err != nil {
       log.Fatalf("Unable to convert the string into int.  %v", err)
   }

   user, err:= controller.repository.GetUser(int64(id))

   if err != nil {
       log.Fatalf("Unable to get user. %v", err)
   }

   json.NewEncoder(w).Encode(user)
}

func (controller UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
   users, err := controller.repository.GetAllUsers()

   if err != nil {
      log.Fatalf("Unable to get all user. %v", err)
   }

   json.NewEncoder(w).Encode(users)
}

func (controller UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
   params := mux.Vars(r)

   id, err := strconv.Atoi(params["id"])

   if err != nil {
       log.Fatalf("Unable to convert the string into int.  %v", err)
   }

   var user models.User

   err = json.NewDecoder(r.Body).Decode(&user)

   if err != nil {
       log.Fatalf("Unable to decode the request body.  %v", err)
   }

   updatedRows :=  controller.repository.UpdateUser(int64(id), user)

   msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

   res := models.Response{
       ID:      int64(id),
       Message: msg,
   }

   json.NewEncoder(w).Encode(res)
}

func (controller UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
   params := mux.Vars(r)

   id, err := strconv.Atoi(params["id"])

   if err != nil {
       log.Fatalf("Unable to convert the string into int.  %v", err)
   }

   deletedRows :=  controller.repository.DeleteUser(int64(id))

   msg := fmt.Sprintf("User deleted successfully. Total rows/record affected %v", deletedRows)

   res := models.Response{
       ID:      int64(id),
       Message: msg,
   }

   json.NewEncoder(w).Encode(res)
}


