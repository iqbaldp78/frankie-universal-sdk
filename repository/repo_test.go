package repository

// import (
// 	"frankie/universal/sdk/model"
// 	"frankie/universal/sdk/utils"
// 	"reflect"
// 	"testing"
// )

// func TestNewRepo(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want Repo
// 	}{
// 		{name: "init", want: &Irepo{}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewRepo(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_irepo_Validate(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		i            *Irepo
// 		input        model.InputValidation
// 		wantUtilsErr *utils.Error
// 	}{
// 		{name: "init", i: &Irepo{}, input: model.InputValidation{}, wantUtilsErr: nil},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotUtilsErr := tt.i.Validate(tt.input); !reflect.DeepEqual(gotUtilsErr, tt.wantUtilsErr) {
// 				t.Errorf("Irepo.Validate() = %v, want %v", gotUtilsErr, tt.wantUtilsErr)
// 			}
// 		})
// 	}
// }
