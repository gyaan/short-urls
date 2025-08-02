package mongo_client

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNew(t *testing.T) {
	type args struct {
		mongoDbConnectionUrl string
		mongoContextTimeout  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Not valid",
			args: args{
				mongoContextTimeout:  10,
				mongoDbConnectionUrl: "some random url",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.mongoDbConnectionUrl, tt.args.mongoContextTimeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestGetClient(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Client
	}{
		{
			name: "not valid",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
