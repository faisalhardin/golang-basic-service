package repo

import (
	"task1/entity"
	"testing"
)

func Test_GetFieldName(t *testing.T) {
	type args struct {
		summary entity.Summary
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful set",
			args: args{
				summary: entity.Summary{},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := GetFieldNames(tt.args.summary)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.Set() err = %v, wantErr = %v", err, tt.wantErr)
			}
			t.Log(got)

		})
	}
}
