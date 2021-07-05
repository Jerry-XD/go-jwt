package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-jwt/forms"
	"go-jwt/routes"

	"github.com/google/go-cmp/cmp"
)

var (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU1MzU1OTQsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoidXNlciJ9.UJh3IoQ6AvnHCkbCOTTGHlPLbRxPfi6HzBI0Fqky_Fs"
)

var router = routes.NewRouter()

func TestCreate(t *testing.T) {
	createData := &forms.AdminCreateInput{
		Name:  "test",
		Email: "test@test.com",
		Age:   25,
	}
	jsonBytes, _ := json.Marshal(createData)

	req := httptest.NewRequest(http.MethodPost, "/admin/data", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Error(fmt.Sprintf("Error http code => expected: %d | actual: %d", http.StatusOK, rec.Code))
	}
}

func TestGetList(t *testing.T) {
	var listDataExpected = []forms.Admin{
		{
			ID:    "1",
			Name:  "Test1",
			Email: "gmail.com",
			Age:   10,
		},
		{
			ID:    "2",
			Name:  "Test2",
			Email: "gmail.com",
			Age:   20,
		},
		{
			ID:    "3",
			Name:  "Test3",
			Email: "gmail.com",
			Age:   30,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/admin/data", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Error(fmt.Sprintf("Error http code => expected: %d | actual: %d", http.StatusOK, rec.Code))
	}

	res := &forms.ListResponse{}
	_ = json.Unmarshal(rec.Body.Bytes(), &res)

	for _, v := range res.Data {
		log.Println(v)
	}

	if !cmp.Equal(listDataExpected, res.Data) {
		t.Error(fmt.Sprintf("Error get data => expected: %+v | actual: %+v", listDataExpected, res.Data))
	}

}
