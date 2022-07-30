package palindrome

import (
	"encoding/json"
	"net/http"

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
		HandlerMessage := []byte(`{
		 "success": false,
		 "message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		json.NewEncoder(w).Encode(HandlerMessage)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var input models.Palindrome
	json.NewDecoder(r.Body).Decode(&input)

	result := h.service.CheckPalindrom(input)
	json.NewEncoder(w).Encode(result)
}
