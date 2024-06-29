package services

import "testing"

func TestSearchWeather(t *testing.T) {
	type args struct {
		zipCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test SearchWeather Success",
			args: args{
				zipCode: "36773460",
			},
			wantErr: false,
		},
		{
			name: "Test SearchWeather Not Found Zip Code",
			args: args{
				zipCode: "49045707",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SearchWeather(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchWeather() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
