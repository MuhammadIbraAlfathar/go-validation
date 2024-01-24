package go_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidate(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nill")
	}
}

func TestValidateVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}
