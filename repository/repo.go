package repository

// import (
// 	"fmt"
// 	"frankie/universal/sdk/model"
// 	"frankie/universal/sdk/utils"
// 	"log"

// 	"github.com/go-playground/validator/v10"
// )

// type Irepo struct{}

// type Repo interface {
// 	Validate(input model.InputValidation) *utils.Error
// }

// func NewRepo() Repo {
// 	return &Irepo{}
// }

// var validate *validator.Validate

// func init() {
// 	validate = validator.New()
// }

// func validateStruct(input *model.InputValidation) (err error) {

// 	// returns nil or ValidationErrors ( []FieldError )
// 	err = validate.Struct(input)
// 	if err != nil {

// 		// this check is only needed when your code could produce
// 		// an invalid value for validation such as interface with nil
// 		// value most including myself do not usually have code like this.
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			fmt.Println(err)
// 			return
// 		}

// 		for _, err := range err.(validator.ValidationErrors) {

// 			fmt.Println(err.Namespace())
// 			fmt.Println(err.Field())
// 			fmt.Println(err.StructNamespace())
// 			fmt.Println(err.StructField())
// 			fmt.Println(err.Tag())
// 			fmt.Println(err.ActualTag())
// 			fmt.Println(err.Kind())
// 			fmt.Println(err.Type())
// 			fmt.Println(err.Value())
// 			fmt.Println(err.Param())
// 			fmt.Println()
// 		}

// 		// from here you can create your own error messages in whatever language you wish
// 		return
// 	}
// 	return
// 	// save user to database
// }

// func (i *Irepo) Validate(input model.InputValidation) (utilsErr *utils.Error) {

// 	validate = validator.New()
// 	log.Println("input", input)

// 	err := validateStruct(&input)
// 	log.Println("err", err)
// 	// err := validate.Struct(input)
// 	// log.Println("err", err)
// 	// for _, e := range err.(validator.ValidationErrors) {
// 	// 	fmt.Println(e)
// 	// }
// 	return
// }
