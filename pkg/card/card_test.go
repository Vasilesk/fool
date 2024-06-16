package card

import (
	"reflect"
	"testing"
)

func TestNewFromString(t *testing.T) {
	t.Parallel()

	type args struct {
		st string
	}

	tests := []struct {
		name    string
		args    args
		want    Card
		wantErr bool
	}{
		{
			name:    "base",
			args:    args{st: "6s"},
			want:    New(SuitSpades, 6),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewFromString(tt.args.st)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
