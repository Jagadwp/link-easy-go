package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/gavv/httpexpect"
)

func TestCreateShortUrl(t *testing.T) {
	handler := InitHandler()

	// run server using httptest
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
	idFloat := obj.Value("data").Object().Value("id").Number().Raw()
	userID := int(idFloat)

	// fmt.Println("Data adalah:\n", userID)
	fmt.Println("tipenya:\n", reflect.TypeOf(userID))

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Insert url, then call get one By ID and found that url", func(t *testing.T) {
		dataForInsert := dto.CreateUrlRequest{
			Title:        "Tutorial docker golang postgres",
			OriginalLink: "https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23&t=106s",
			UserID:       userID,
		}

		auth.POST("/urls/generate").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})
}

func TestGetUrlByUserId(t *testing.T) {
	handler := InitHandler()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"username": "caesaryo_slf",
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

	t.Run("Expected Find ALL short url list", func(t *testing.T) {
		auth.GET("/urls").
			Expect().
			Status(http.StatusOK).JSON().Object()

	})

}

func TestUpdateUrl(t *testing.T) {
	handler := InitHandler()

	// run server using httptest
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

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Update url, then call get one By ID and found that url", func(t *testing.T) {
		dataForInsert := dto.UpdateUrlRequest{
			Title:        "Tutorial docker golang postgres",
			ShortLink:    "tutordocker",
			OriginalLink: "https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23&t=106s",
		}

		auth.PUT("/urls/{id}").WithPath("id", 12).
			WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Update url, then call get one By ID and not found that url", func(t *testing.T) {
		dataForInsert := dto.UpdateUrlRequest{
			Title:        "Tutorial docker golang postgres",
			ShortLink:    "tutor_docker",
			OriginalLink: "https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23&t=106s",
		}

		auth.PUT("/urls/{id}").WithPath("id", 999).
			WithJSON(dataForInsert).
			Expect().
			Status(http.StatusNotFound)
	})
}

func TestDeleteUrl(t *testing.T) {
	handler := InitHandler()

	// run server using httptest
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

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected delete url", func(t *testing.T) {
		auth.DELETE("/urls/{id}").WithPath("id", 14).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected delete url, then call get one By ID But NOT found that url", func(t *testing.T) {
		auth.DELETE("/urls/{id}").WithPath("id", 999).
			Expect().
			Status(http.StatusNotFound)
	})
}

func TestRedirectUrl(t *testing.T) {
	handler := InitHandler()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	t.Run("Expected redirect url", func(t *testing.T) {
		e.GET("/{id}").WithPath("id", "myits").
			Expect().
			Status(http.StatusOK)
	})
}
