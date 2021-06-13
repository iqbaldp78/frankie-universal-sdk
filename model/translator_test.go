package model

import (
	"reflect"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Test_translator(t *testing.T) {
	tests := []struct {
		name    string
		want    ut.Translator
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translator()
			if (err != nil) != tt.wantErr {
				t.Errorf("translator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("translator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectedfieldValidation(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectedfieldValidation(tt.args.fl); got != tt.want {
				t.Errorf("selectedfieldValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
