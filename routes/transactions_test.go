package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simpleBankingApplication/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMakePayment(t *testing.T) {
	router := gin.Default()
	router.POST("/make-payment", MakePayment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/make-payment", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"Invalid request"}`, w.Body.String())
}

func TestMakePaymentWithValidRequest(t *testing.T) {
	router := gin.Default()
	router.POST("/make-payment", MakePayment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/make-payment", bytes.NewBufferString(`{"user_id":2,"amount":100,"account_number":"0012345679","transaction_type":"debit"}`))
	router.ServeHTTP(w, req)

	respi := w.Body.String()

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message":"transaction saved successfully"}`,respi)
}
func TestGetPaymentByReference(t *testing.T) {
	router := gin.Default()
	router.GET("/get-payment/:reference", GetPaymentByReference)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get-payment/123456789", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Payment fetched successfully"}`, w.Body.String())
}

func TestGetPaymentByReferenceWithError(t *testing.T) {
	router := gin.Default()
	router.GET("/get-payment/:reference", GetPaymentByReference)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get-payment/invalid_reference", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error":"An error occurred while fetching the payment"}`, w.Body.String())
}

func TestCallPaymentService(t *testing.T) {
	request := models.Request{
		Account_id: "0012345679",
		Reference: "123456789",
		Amount:    100,
	}

	expectedResponse := models.Response{
		Account_id: "0012345679",
		Reference: "123456789",
		Amount:    100,
	}

	// Mock the external endpoint using httptest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "/make-payment", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		// Verify the request body
		var req models.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, request, req)

		// Send the expected response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	// Set the server URL as the endpoint for the test
	server.URL = "http://thirdparty.com/make-payment"

	// Call the method under test
	response, err := CallPaymentService(request)

	// Verify the response and error
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
func TestGetPaymentsByReference(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "/get-payment/123456789", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Send the expected response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Response{
			Account_id: "0012345679",
			Reference:  "123456789",
			Amount:     100,
		})
	}))
	defer server.Close()

	// Call the method under test
	response, err := GetPaymentsByReference("123456789")

	// Verify the response and error
	assert.NoError(t, err)
	assert.Equal(t, models.Response{
		Account_id: "0012345679",
		Reference:  "123456789",
		Amount:     100,
	}, response)
}