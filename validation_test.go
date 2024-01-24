package go_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
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

func TestNestedStruct(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type RegisterUser struct {
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
		Address         Address
	}

	validate := validator.New()

	registerUser := RegisterUser{
		Email:           "test@gmail.com",
		Password:        "test123",
		ConfirmPassword: "test123",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidateCollection(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type RegisterUser struct {
		Email           string    `validate:"required,email"`
		Password        string    `validate:"required,min=5"`
		ConfirmPassword string    `validate:"required,min=5,eqfield=Password"`
		Address         []Address `validate:"required,dive"`
	}

	validate := validator.New()

	registerUser := RegisterUser{
		Email:           "test@gmail.com",
		Password:        "test123",
		ConfirmPassword: "test123",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidateBasicCollection(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type RegisterUser struct {
		Email           string    `validate:"required,email"`
		Password        string    `validate:"required,min=5"`
		ConfirmPassword string    `validate:"required,min=5,eqfield=Password"`
		Address         []Address `validate:"required,dive"`
		Hobby           []string  `validate:"required,dive,required,min=5"`
	}

	validate := validator.New()

	registerUser := RegisterUser{
		Email:           "test@gmail.com",
		Password:        "test123",
		ConfirmPassword: "test123",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobby: []string{
			"Gaming",
			"Coding",
			"",
			"tes",
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type User struct {
		Id       string `validate:"varchar,min=5"`
		Email    string `validate:"varchar,email"`
		Username string `validate:"varchar"`
	}

	user := User{
		Id:       "123",
		Email:    "test@gmail.com",
		Username: "testt",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func MusValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("mustupper", MusValidUsername)
	if err != nil {
		fmt.Println(err.Error())
	}

	type User struct {
		Id       string `validate:"required,min=5"`
		Email    string `validate:"required,email"`
		Username string `validate:"mustupper"`
	}

	user := User{
		Id:       "123",
		Email:    "test@gmail.com",
		Username: "TESTT",
	}

	err = validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}
	value := field.Field().String()

	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidParam(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("validPin", MustValidPin)
	if err != nil {
		fmt.Println(err.Error())
	}

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,validPin=6"`
	}

	login := Login{
		Phone: "088678686786",
		Pin:   "133321",
	}

	err = validate.Struct(login)
	if err != nil {
		fmt.Println(err)
	}

}
