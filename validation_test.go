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

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()
	password := "test123"
	confirmPassword := "test123"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}
