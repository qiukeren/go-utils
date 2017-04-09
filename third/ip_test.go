package third

import (
	"reflect"
	"testing"
)

func TestOutIp(t *testing.T) {
	tests := []struct {
		name    string
		wantIp  *Ip
		wantErr bool
	}{
		{"no nil test", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp, err := OutIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("OutIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIp, tt.wantIp) {
				t.Errorf("OutIp() = %v, want %v", gotIp, tt.wantIp)
			}
		})
	}
}
