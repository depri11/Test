package palindrome

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/interfaces"
	"github.com/depri11/Test/src/models"
)

// handler struct for handling palindrome
type handler struct {
	service interfaces.ServicePalindrome
}

// NewHandler returns a new palindrome handler
func NewHandler(service interfaces.ServicePalindrome) *handler {
	return &handler{service}
}

// CheckPalindrom handles the request for checking palindrome
func (h *handler) Palindrome(w http.ResponseWriter, r *http.Request) {
	// Check if the request contains is equal or not to the request method
	if r.Method != "POST" {
		// Add the response return message
		response := helpers.ResponseAPI(405, "Check your HTTP method: Invalid HTTP method executed", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the response header to json

	var input models.Palindrome
	// Decode the request body into palindrome struct
	json.NewDecoder(r.Body).Decode(&input)

	// Checking Palindrome and return the response
	response := h.service.CheckPalindrom(input)
	json.NewEncoder(w).Encode(response)
}
