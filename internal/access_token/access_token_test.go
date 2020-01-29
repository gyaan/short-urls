package access_token

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	type args struct {
		UserId          string
		tokenExpiryTime int64
		jwtSecret       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				UserId:          "5e307f7bcc81068098d1ae71",
				tokenExpiryTime: 432000,
				jwtSecret:       "+,g~9Ywa8)7D<nbR",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetToken(tt.args.UserId, tt.args.tokenExpiryTime, tt.args.jwtSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		tokenString string
		jwtSecret   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXCJ9.eyJpZCI6IjVlMzA3ZjdiY2M4MTA2ODA5OGQxYWU3MSIsImV4cCI6MTYwNjIzNjMyMX0.fx_UM7l5RCiw82bY8N6iywSPX4g4tuXtMXUCTPPMPTw",
				jwtSecret:   "+,g~9Ywa8)7D<nbR",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateToken(tt.args.tokenString, tt.args.jwtSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
