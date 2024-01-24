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

	//membandingkan dua nilai variable harus sama
	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "1232312"

	err := validate.Var(user, "required,number")

	if err != nil {
		fmt.Println(err.Error())
	}
}
