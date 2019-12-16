package uulid

import (
	"bytes"
	"github.com/oklog/ulid/v2"
	"reflect"
	"testing"
	"time"
)

func TestTimeAsULID(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want ulid.ULID
	}{
		{
			"t1",
			args{t: time.Unix(1576481036, 999999999)},
			ulid.MustNew(ulid.Timestamp(time.Unix(1576481036, 999999999)), bytes.NewReader([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeAsULID(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeAsULID() = %v, want %v", got, tt.want)
			}
		})
	}
}
