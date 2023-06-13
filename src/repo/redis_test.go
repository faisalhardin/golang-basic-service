package repo

import (
	"reflect"
	"task1/entity"
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
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

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
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

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
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

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
				Key:     "stock:BBCAB",
				Field:   "closePrice",
				Value:   "2000",
			},
			want:    "2000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

			got, err := red.HSet(tt.args.Key, tt.args.Field, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HSet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if got != 1 {
				t.Errorf("red.HSet() = %v, want = %v", got, 1)
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
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

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
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

			got, err := red.HGetAll(tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HGet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("red.HGet() = %v, want = %v", got, tt.want)
			}

			t.Log(got)

		})
	}
}

func Test_HSetSummary(t *testing.T) {
	mockSummary := entity.Summary{
		PreviousPrice: 8000,
		OpenPrice:     8050,
		HighestPrice:  8100,
		LowestPrice:   7950,
		ClosePrice:    8100,
		Volume:        900,
		Value:         7210000,
		IsNewDay:      false,
	}
	type args struct {
		Address string
		Key     string
		Summary entity.Summary
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
				Summary: mockSummary,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

			got, err := red.HSetSummary(tt.args.Key, tt.args.Summary)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HSetAll() err = %v, wantErr = %v", err, tt.wantErr)
			}
			t.Log(got)
		})
	}
}

func Test_HGetSummary(t *testing.T) {
	wantSummary := entity.Summary{
		PreviousPrice: 8000,
		OpenPrice:     8050,
		HighestPrice:  8100,
		LowestPrice:   7950,
		ClosePrice:    8100,
		Volume:        900,
		Value:         7210000,
		IsNewDay:      false,
	}
	type args struct {
		Address string
		Key     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    entity.Summary
	}{
		{
			name: "Successful set",
			args: args{
				Address: "127.0.0.1:6379",
				Key:     "stock:BBCA",
			},
			want:    wantSummary,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			red := NewRedisRepo(&RedisOptions{Address: tt.args.Address})

			got, err := red.HGetSummary(tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("red.HGet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("red.HGet() = %v, want = %v", got, tt.want)
			}

			t.Log(got)

		})
	}
}
