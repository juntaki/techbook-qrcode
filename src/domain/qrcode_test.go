package domain

import (
	"reflect"
	"testing"
)

func TestQRCode_GetQRCode(t *testing.T) {
	type fields struct {
		ID string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QRCode{
				ID: tt.fields.ID,
			}
			if got := q.GetQRCode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QRCode.GetQRCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
