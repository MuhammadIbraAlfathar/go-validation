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

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "23131321312"

	err := validate.Var(user, "required,number,min=10,max=20")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateStruct(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Username string `validate:"required,max=50"`
	}

	type CreatePassword struct {
		Password        string `validate:"required,eqfield"`
		ConfirmPassword string `validate:"required,eqfield"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Email:    "testing@gmail.com",
		Username: "testing123",
	}

	createPassword := CreatePassword{
		Password:        "test",
		ConfirmPassword: "test",
	}

	err2 := validate.VarWithValue(createPassword.Password, createPassword.ConfirmPassword, "eqfield")
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateErrors(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Username string `validate:"required,max=50"`
	}

	type CreatePassword struct {
		Password        string `validate:"required,eqfield"`
		ConfirmPassword string `validate:"required,eqfield"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Email:    "testingcom",
		Username: "testing123",
	}

	createPassword := CreatePassword{
		Password:        "test",
		ConfirmPassword: "test",
	}

	err2 := validate.VarWithValue(createPassword.Password, createPassword.ConfirmPassword, "eqfield")
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			fmt.Println("Error" + fieldError.Error() + "On Tag:" + fieldError.Tag())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()

	registerUser := RegisterUser{
		Email:           "test@gmail.com",
		Password:        "test123",
		ConfirmPassword: "test123",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}

}
