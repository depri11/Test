package interfaces

import (
	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/models"
)

// Interface service represents a all method service
type ServicePalindrome interface {
	CheckPalindrom(input models.Palindrome) *helpers.Response
}
