package pagination

import (
	"reflect"
	"testing"
)

func Test_page_GetPagination(t *testing.T) {
	type fields struct {
		currentPage int64
		totalItem   int64
		data        interface{}
		perPageItem int
	}
	tests := []struct {
		name    string
		fields  fields
		want    Response
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				currentPage: 2,
				totalItem:   100,
				data: []struct{ name string }{
					{name: "gyaneshwar"},
					{name: "gyaan"},
				},
				perPageItem: 10,
			},
			want: Response{
				TotalPage:    10,
				TotalItem:    100,
				PreviousPage: 1,
				NextPage:     3,
				Data: []struct{ name string }{
					{name: "gyaneshwar"},
					{name: "gyaan"},
				},
				CurrentPage: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &page{
				currentPage: tt.fields.currentPage,
				totalItem:   tt.fields.totalItem,
				data:        tt.fields.data,
				perPageItem: tt.fields.perPageItem,
			}
			got, err := p.GetPagination()
			if (err != nil) != tt.wantErr {
				t.Errorf("page.GetPagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("page.GetPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}
