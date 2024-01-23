package go_validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidate(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nill")
	}
}
