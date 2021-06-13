package model

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func translator() (ut.Translator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}

	_ = validate.RegisterTranslation("selectedfield", trans, func(ut ut.Translator) error {
		err := ut.Add("selectedfieldcheckType", "{0} must be DEVICE | BIOMETRIC | COMBO", true)
		if err != nil {
			return err
		}

		err = ut.Add("selectedfieldactivityType", "{0} must be SIGNUP | LOGIN | PAYMENT | CONFIRMATION | _<Vendor Specific List>", true)
		if err != nil {
			return err
		}
		return nil
	}, func(ut ut.Translator, fe validator.FieldError) string {
		if fe.Field() == "CheckType" {
			t, _ := ut.T("selectedfieldcheckType", fe.Field())
			return t
		} else if fe.Field() == "ActivityType" {
			t, _ := ut.T("selectedfieldactivityType", fe.Field())
			return t
		}
		return ""
	})

	validate.RegisterValidation("selectedfield", selectedfieldValidation)
	return trans, nil
}

func selectedfieldValidation(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	fieldName := fl.FieldName()

	if fieldName == "CheckType" {
		x0 := val == Biometric
		x1 := val == Device
		x2 := val == Combo
		if x0 || x1 || x2 {
			return true
		}
	} else if fieldName == "ActivityType" {
		x0 := val == Signup
		x1 := val == Login
		x2 := val == Payment
		x3 := val == Confirmation

		//check underscore
		rune := string(val[0])
		if rune == "_" {
			return true
		}

		if x0 || x1 || x2 || x3 {
			return true
		}
	}

	return false
}
