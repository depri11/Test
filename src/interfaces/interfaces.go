package interfaces

import (
	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/models"
)

type ServicePalindrome interface {
	CheckPalindrom(input models.Palindrome) *helpers.Response
}
