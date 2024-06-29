package adapters

import (
	"afonsojota/go-expert-cep-labs/entity"
	"reflect"
	"testing"
)

func TestSearchZipCode(t *testing.T) {
	type args struct {
		zipCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Cep
		wantErr bool
	}{
		{
			name: "Test SearchZipCode Success",
			args: args{
				zipCode: "36790000",
			},
			want: &entity.Cep{
				Cep:          "36790-000",
				Log:          "",
				Neighborhood: "",
				Locale:       "Miraí",
				UF:           "MG",
				IBGE:         "3142205",
				DDD:          "32",
			},
			wantErr: false,
		},
		{
			name: "Test SearchZipCode When Error Invalid Zip Code",
			args: args{
				zipCode: "00000000",
			},
			want:    &entity.Cep{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchZipCode(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchZipCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchZipCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
