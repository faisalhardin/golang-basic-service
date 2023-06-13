package repo

import (
	"reflect"
	"testing"
)

func Test_Ping(t *testing.T) {
	type args struct {
		Address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful ping",
			args: args{
				Address: "127.0.0.1:6379",
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			err := red.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("red.Ping() err = %v, wantErr = %v", err, tt.wantErr)
			}

		})
	}
}

func Test_SetKeyValue(t *testing.T) {
	type args struct {
		Address string
		Key     string
		Value   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "test",
				Value:   "success",
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			got, err := red.SetKeyValue(tt.args.Key, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.Set() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != "OK" {
				t.Errorf("red.Set() = %v, want = %v", got, "OK")
			}

		})
	}
}

func Test_Get(t *testing.T) {
	type args struct {
		Address string
		Key     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "test",
			},
			want:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			got, err := red.Get(tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.Get() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("red.Get() = %v, want = %v", got, tt.want)
			}

		})
	}
}

func Test_HSet(t *testing.T) {
	type args struct {
		Address string
		Key     string
		Field   string
		Value   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "stock:BBCA",
				Field:   "closePrice",
				Value:   "2000",
			},
			want:    "2000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			_, err := red.HSet(tt.args.Key, tt.args.Field, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HSet() err = %v, wantErr = %v", err, tt.wantErr)
			}

		})
	}
}

func Test_HGet(t *testing.T) {
	type args struct {
		Address string
		Key     string
		Param   string
		Value   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "stock:BBCA",
				Param:   "closePrice",
				Value:   "2000",
			},
			want:    "2000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			got, err := red.HGet(tt.args.Key, tt.args.Param)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HGet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("red.HGet() = %v, want = %v", got, tt.want)
			}

		})
	}
}

func Test_HGetAll(t *testing.T) {
	type args struct {
		Address string
		Key     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []string
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "stock:BBCA",
			},
			want:    []string{"closePrice", "2000", "openPrice", "2050"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := New(&Options{Address: tt.args.Address})

			got, err := red.HGetAll(tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HGet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("red.HGet() = %v, want = %v", got, tt.want)
			}

		})
	}
}
