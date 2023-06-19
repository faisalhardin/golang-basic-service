package filereader

import (
	"io"
	"io/fs"
	"os"
	"reflect"
	"task1/entity"
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

func Test_ReadFilesWithChannel(t *testing.T) {

	mockBufioNewScanner := &mockrepo.BufioNewScanner{}

	type args struct {
		prefix    string
		filePaths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    entity.StockCodeToTransactionLogKeyValue
		patch   func()
	}{
		{
			name: "Successful read row of file too channel",
			args: args{
				filePaths: []string{""},
			},
			want: entity.StockCodeToTransactionLogKeyValue{
				StockCode:      "BBRI",
				TransactionLog: mockReadLine,
			},
			wantErr: false,
			patch: func() {
				osOpenFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
					return &os.File{}, nil
				}

				bufioNewScanner = func(r io.Reader) mockrepo.NewScanner {
					return mockBufioNewScanner
				}
				mockrepo.GetBuffioScanFunc = func() bool {
					return true
				}
				mockrepo.GetBuffioBytesFunc = func() []byte {
					mockrepo.GetBuffioScanFunc = func() bool {
						return false
					}
					return []byte(mockReadLine)
				}

				mockrepo.GetBuffioErrFunc = func() error {
					return nil
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			got := <-ReadFilesWithChannel(tt.args.prefix, tt.args.filePaths)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ReadFilesWithChannel() = %v, want = %v", got, tt.want)
			}
		})
	}
}
