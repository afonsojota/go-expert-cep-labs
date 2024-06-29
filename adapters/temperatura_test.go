package adapters

import (
	"afonsojota/go-expert-cep-labs/entity"
	"reflect"
	"testing"
)

func TestGetWeather(t *testing.T) {
	type args struct {
		locale string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Temperatura
		wantErr bool
	}{
		{
			name: "Test GetWeather Error Invalid Locale",
			args: args{
				locale: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWeather(tt.args.locale)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() got = %v, want %v", got, tt.want)
			}
		})
	}
}
