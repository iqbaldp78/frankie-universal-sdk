package model

import (
	"errors"
	"frankie/universal/sdk/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	Device    string = "DEVICE"
	Biometric string = "BIOMETRIC"
	Combo     string = "COMBO"

	Signup       string = "SIGNUP"
	Login        string = "LOGIN"
	Payment      string = "PAYMENT"
	Confirmation string = "CONFIRMATION"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type InputValidation struct {
	CheckType       string     `json:"checkType" validate:"required,selectedfield"`
	ActivityType    string     `json:"activityType" validate:"required,selectedfield"`
	CheckSessionKey string     `json:"checkSessionKey" validate:"required"`
	ActivityData    []Activity `json:"activityData" validate:"required,dive,required"`
}

type Activity struct {
	KvpKey   string `json:"kvpKey" validate:"required"`
	KvpValue string `json:"kvpValue" validate:"required"`
	KvpType  string `json:"kvpType" validate:"required"`
}

func init() {
	validate = validator.New()
}

func (i *InputValidation) Validate() (utilsErr *utils.Error) {
	trans, err := translator()
	if err != nil {
		return utils.SetError(utils.OperationFailed, err)
	}

	if err := validate.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		errValue := ""
		for _, v := range errs.Translate(trans) {
			errValue = v
		}
		return utils.SetError(utils.IncorrectParam, errors.New(errValue))
	}

	return
}
