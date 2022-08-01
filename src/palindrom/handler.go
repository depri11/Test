package palindrome

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/interfaces"
	"github.com/depri11/Test/src/models"
)

type handler struct {
	service interfaces.ServicePalindrome
}

func NewHandler(service interfaces.ServicePalindrome) *handler {
	return &handler{service}
}

func (h *handler) Palindrome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Add the response return message
		response := helpers.ResponseAPI(405, "Check your HTTP method: Invalid HTTP method executed", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var input models.Palindrome
	json.NewDecoder(r.Body).Decode(&input)

	response := h.service.CheckPalindrom(input)
	json.NewEncoder(w).Encode(response)
}
