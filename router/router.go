package router

import (
    "go-project/controller"
    "go-project/repository"
    "net/http"

    "github.com/gorilla/mux"
)

func Router() *mux.Router {
    userRepository := repository.UserRepository{}
    userController := controller.NewUserController(userRepository)
    router := mux.NewRouter()

    router.HandleFunc("/api/user/{id}", userController.GetUser).Methods(http.MethodGet, http.MethodOptions)
    router.HandleFunc("/api/user", userController.GetAllUser).Methods(http.MethodGet, http.MethodOptions)
    router.HandleFunc("/api/newuser", userController.CreateUser).Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/api/user/{id}", userController.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
    router.HandleFunc("/api/deleteuser/{id}", userController.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)

    router.Use(applyCors)
    return router
}

func applyCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Add("Access-Control-Allow-Origin", "*")
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Add("Access-Control-Allow-Headers", "access-control-allow-headers,access-control-allow-methods,access-control-allow-origin,authorization")
        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}