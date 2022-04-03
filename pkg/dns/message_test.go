package dns

import (
	"reflect"
	"testing"
)

func TestHeaderMarshal(t *testing.T) {
	type args struct {
		h *Header
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"allZero",
			args{&Header{}},
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"ID set",
			args{&Header{ID: 0x1234}},
			[]byte{0x12, 0x34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"QR set",
			args{&Header{QR: true}},
			[]byte{0, 0, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"Opcode set",
			args{&Header{Opcode: 0x0f}},
			[]byte{0, 0, 0x78, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"AA set",
			args{&Header{AA: true}},
			[]byte{0, 0, 0x04, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"TC set",
			args{&Header{TC: true}},
			[]byte{0, 0, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"RD set",
			args{&Header{RD: true}},
			[]byte{0, 0, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"RA set",
			args{&Header{RA: true}},
			[]byte{0, 0, 0, 0x80, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"RCode set",
			args{&Header{RCode: 0xff}},
			[]byte{0, 0, 0, 0x0f, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"QDCount set",
			args{&Header{QDCount: 0x1234}},
			[]byte{0, 0, 0, 0, 0x12, 0x34, 0, 0, 0, 0, 0, 0},
			false,
		},
		{
			"ANCount set",
			args{&Header{ANCount: 0x1234}},
			[]byte{0, 0, 0, 0, 0, 0, 0x12, 0x34, 0, 0, 0, 0},
			false,
		},
		{
			"NSCount set",
			args{&Header{NSCount: 0x1234}},
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34, 0, 0},
			false,
		},
		{
			"ARCount set",
			args{&Header{ARCount: 0x1234}},
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HeaderMarshal(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("HeaderMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeaderMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderUnmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *Header
		want       int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HeaderUnmarshal(tt.args.h, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HeaderUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HeaderUnmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
