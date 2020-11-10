package controller

import (
    "bytes"
    "github.com/golang/mock/gomock"
    "github.com/gorilla/mux"
    "go-project/mocks"
    "go-project/models"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestCreateUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    userMock := mocks.NewMockUserRepositoryInterface(mockCtrl)
    userController := NewUserController(userMock)

    var jsonStr = []byte(`{"id":1,"name":"Eduardo","location":"SC","age":10}`)
    req, err := http.NewRequest(http.MethodPost, "/api/newuser", bytes.NewBuffer(jsonStr))
    if err != nil {
        t.Fatal(err)
    }

    user := models.User{ID: 1, Name: "Eduardo", Location: "SC", Age: 10}
    userMock.EXPECT().InsertUser(user).Return(user.ID)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(userController.CreateUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"id":1,"message":"User created successfully"}`
    if strings.Trim(rr.Body.String(), "\n") != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestGetUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    userMock := mocks.NewMockUserRepositoryInterface(mockCtrl)
    userController := NewUserController(userMock)

    req, err := http.NewRequest(http.MethodGet, "/api/user/1", nil)
    vars := map[string]string{"id": "1"}
    req = mux.SetURLVars(req, vars)

    if err != nil {
        t.Fatal(err)
    }

    userMock.EXPECT().GetUser(int64(1)).Return(models.User{ID: 1, Name: "Eduardo", Location: "SC", Age: 10}, nil)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(userController.GetUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"id":1,"name":"Eduardo","location":"SC","age":10}`
    if strings.Trim(rr.Body.String(), "\n") != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestGetAllUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    userMock := mocks.NewMockUserRepositoryInterface(mockCtrl)
    userController := NewUserController(userMock)

    req, err := http.NewRequest(http.MethodGet, "/api/user", nil)

    if err != nil {
        t.Fatal(err)
    }
    users := []models.User{{1, "Eduardo", "SC", 12}}
    userMock.EXPECT().GetAllUsers().Return(users, nil)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(userController.GetAllUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `[{"id":1,"name":"Eduardo","location":"SC","age":12}]`
    if strings.Trim(rr.Body.String(), "\n") != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestUpdateUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    userMock := mocks.NewMockUserRepositoryInterface(mockCtrl)
    userController := NewUserController(userMock)

    var jsonStr = []byte(`{"id":1,"name":"Eduardo","location":"SC","age":10}`)
    req, err := http.NewRequest(http.MethodPut, "/api/user/1",  bytes.NewBuffer(jsonStr))
    vars := map[string]string{"id": "1"}
    req = mux.SetURLVars(req, vars)

    if err != nil {
        t.Fatal(err)
    }
    user := models.User{ID: 1, Name: "Eduardo", Location: "SC", Age: 10}
    userMock.EXPECT().UpdateUser(int64(1), user).Return(int64(1))

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(userController.UpdateUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"id":1,"message":"User updated successfully. Total rows/record affected 1"}`
    if strings.Trim(rr.Body.String(), "\n") != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestDeleteUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    userMock := mocks.NewMockUserRepositoryInterface(mockCtrl)
    userController := NewUserController(userMock)

    req, err := http.NewRequest(http.MethodPut, "/api/deleteuser/1", nil)
    vars := map[string]string{"id": "1"}
    req = mux.SetURLVars(req, vars)

    if err != nil {
        t.Fatal(err)
    }
    userMock.EXPECT().DeleteUser(int64(1)).Return(int64(1))

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(userController.DeleteUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"id":1,"message":"User deleted successfully. Total rows/record affected 1"}`
    if strings.Trim(rr.Body.String(), "\n") != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}