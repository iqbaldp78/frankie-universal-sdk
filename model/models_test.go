package model

import (
	"errors"
	"frankie/universal/sdk/utils"
	"reflect"
	"testing"
)

func TestInputValidation_Validate(t *testing.T) {
	tests := []struct {
		name         string
		i            *InputValidation
		wantUtilsErr *utils.Error
	}{
		{
			name: "normal flow",
			i: &InputValidation{
				CheckType:       "DEVICE",
				ActivityType:    "SIGNUP",
				CheckSessionKey: "string",
				ActivityData: []Activity{
					{
						KvpKey:   "ip.address",
						KvpValue: "1.23.45.123",
						KvpType:  "general.string",
					},
				},
			},
			wantUtilsErr: nil,
		}, {
			name: "checkType field missing",
			i: &InputValidation{
				ActivityType:    "SIGNUP",
				CheckSessionKey: "string",
				ActivityData: []Activity{
					{
						KvpKey:   "ip.address",
						KvpValue: "1.23.45.123",
						KvpType:  "general.string",
					},
				},
			},
			wantUtilsErr: utils.SetError(utils.IncorrectParam, errors.New("CheckType is a required field")),
		}, {
			name: "checkType not in selected list",
			i: &InputValidation{
				CheckType:       "TEST",
				ActivityType:    "SIGNUP",
				CheckSessionKey: "string",
				ActivityData: []Activity{
					{
						KvpKey:   "ip.address",
						KvpValue: "1.23.45.123",
						KvpType:  "general.string",
					},
				},
			},
			wantUtilsErr: utils.SetError(utils.IncorrectParam, errors.New("CheckType must be DEVICE | BIOMETRIC | COMBO")),
		},
		{
			name: "activityType not in selected list",
			i: &InputValidation{
				CheckType:       "DEVICE",
				ActivityType:    "TEST",
				CheckSessionKey: "string",
				ActivityData: []Activity{
					{
						KvpKey:   "ip.address",
						KvpValue: "1.23.45.123",
						KvpType:  "general.string",
					},
				},
			},
			wantUtilsErr: utils.SetError(utils.IncorrectParam, errors.New("ActivityType must be SIGNUP | LOGIN | PAYMENT | CONFIRMATION | _<Vendor Specific List>")),
		},
		{
			name: "activityType checking vendor specific list",
			i: &InputValidation{
				CheckType:       "DEVICE",
				ActivityType:    "_TEST",
				CheckSessionKey: "string",
				ActivityData: []Activity{
					{
						KvpKey:   "ip.address",
						KvpValue: "1.23.45.123",
						KvpType:  "general.string",
					},
				},
			},
			wantUtilsErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUtilsErr := tt.i.Validate(); !reflect.DeepEqual(gotUtilsErr, tt.wantUtilsErr) {
				t.Errorf("InputValidation.Validate() = %v, want %v", gotUtilsErr, tt.wantUtilsErr)
			}
		})
	}
}
