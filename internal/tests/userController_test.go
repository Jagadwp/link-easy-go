package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jagadwp/link-easy-go/internal/services/helper"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/gavv/httpexpect"
)

func TestGetUserData(t *testing.T) {
	handler := InitHandler()

	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"username": "jagadwp",
		"password": "admin",
	}
	// get token

	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("data").Object().Value("token").String().Raw()

	// fmt.Println("Data adalah:\n", token)
	// fmt.Println("tipenya:\n", reflect.TypeOf(token))

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected find user data", func(t *testing.T) {
		auth.GET("/users/{id}").WithPath("id", 1).
			Expect().
			Status(http.StatusOK).JSON().Object()
	})

	t.Run("Expected find user data, But NOT found that user", func(t *testing.T) {
		auth.GET("/users/{id}").WithPath("id", 999).
			Expect().
			Status(http.StatusNotFound).JSON().Object()
	})

}

func TestUpdateUserData(t *testing.T) {
	handler := InitHandler()

	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"username": "faizul",
		"password": "admin",
	}
	// get token

	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("data").Object().Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Update user, then call get one By ID and found that user", func(t *testing.T) {
		newPassword, _ := helper.Hash("passbaru")

		dataForUpdate := dto.UpdateUserRequest{
			Username: "User Y",
			Fullname: "User Y Updated",
			Email:    "new@mail.com",
			Password: newPassword,
			Admin:    true,
		}

		auth.PUT("/users/{id}").WithPath("id", 10).
			WithJSON(dataForUpdate).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Update user, then call get one By ID But NOT found that user", func(t *testing.T) {
		newPassword, _ := helper.Hash("passbaru")

		dataForUpdate := dto.UpdateUserRequest{
			Username: "User X",
			Fullname: "User X Updated",
			Email:    "updated@mail.com",
			Password: newPassword,
			Admin:    true,
		}

		auth.PUT("/users/{id}").WithPath("id", 999).
			WithJSON(dataForUpdate).Expect().
			Status(http.StatusNotFound)
	})
}

func TestDeleteUserData(t *testing.T) {
	handler := InitHandler()

	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"username": "faizul",
		"password": "admin",
	}
	// get token

	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("data").Object().Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected delete user, then call get one By ID and found that user", func(t *testing.T) {
		auth.DELETE("/users/{id}").WithPath("id", 11).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected delete user, then call get one By ID But NOT found that user", func(t *testing.T) {
		auth.DELETE("/users/{id}").WithPath("id", 999).
			Expect().
			Status(http.StatusNotFound)
	})
}
