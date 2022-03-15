package tests

import (
	"api/db"
	"api/domains/users/mocks"
	"api/domains/users/models"
	"api/support/response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	db.TruncateTable("users")
}

func TestGetAllUsers(t *testing.T) {
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserById(t *testing.T) {
	user := mocks.CreateUser()
	id := strconv.FormatUint(uint64(user.ID), 10)
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	testUser := models.User{}
	wrap := response.WrapData(testUser)
	json.Unmarshal(w.Body.Bytes(), &wrap)
	data, _ := wrap["data"]
	assert.Equal(t, data.(map[string]interface{})["name"], user.Name)
	assert.Equal(t, data.(map[string]interface{})["email"], user.Email)
	assert.Equal(t, fmt.Sprint(data.(map[string]interface{})["id"]), id)
}

func TestDeleteUser(t *testing.T) {
	user := mocks.CreateUser()
	id := strconv.FormatUint(uint64(user.ID), 10)
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoreUser(t *testing.T) {
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	user := models.User{
		Name:     "Test",
		Email:    "test@example.com",
		Password: "password",
	}
	data, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(data))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFailToStoreUser(t *testing.T) {
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	user := models.User{
		Name:     "",
		Email:    "",
		Password: "",
	}
	data, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(data))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestUpdateUser(t *testing.T) {
	oldUser := mocks.CreateUser()
	id := strconv.FormatUint(uint64(oldUser.ID), 10)
	router := SetupFeatureTests()
	w := httptest.NewRecorder()
	user := models.User{
		Name:  "Test 123",
		Email: "test123@example.com",
	}
	data, _ := json.Marshal(user)
	req, _ := http.NewRequest("PATCH", "/users/"+id, bytes.NewBuffer(data))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	testUser := models.User{}
	wrap := response.WrapData(testUser)
	json.Unmarshal(w.Body.Bytes(), &wrap)
	data2, _ := wrap["data"]
	assert.Equal(t, data2.(map[string]interface{})["name"], user.Name)
	assert.Equal(t, data2.(map[string]interface{})["email"], user.Email)
}
