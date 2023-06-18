package filereader

import (
	"io/fs"
	"reflect"
	mockrepo "task1/entity/mock"
	"testing"
)

var (
	mockReadLine = `{"type":"A","order_book":"911","price":"4650","stock_code":"BBRI"}`
)

func Test_ListFiles(t *testing.T) {

	mockFs := mockrepo.FileInfo{}

	type args struct {
		fileDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []string
		patch   func()
	}{
		{
			name: "Successful list file",
			args: args{
				fileDir: "",
			},
			want: []string{
				"filename",
			},
			wantErr: false,
			patch: func() {
				ioutilReadDir = func(dirname string) ([]fs.FileInfo, error) {
					return []fs.FileInfo{mockFs}, nil
				}

				mockrepo.GetIsDir = func() bool {
					return false
				}

				mockrepo.GetName = func() string {
					return "filename"
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.patch()
			got, err := ListFiles(tt.args.fileDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFiles() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ListFiles() = %v, want = %v", got, tt.want)
			}
		})
	}
}
