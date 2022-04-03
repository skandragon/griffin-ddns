/*
 * Copyright 2022 Michael Graff.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dns

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		wantHeader Header
		want       int
		wantErr    bool
	}{
		{
			"all zeros",
			args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{},
			12,
			false,
		},
		{
			"ID set",
			args{[]byte{0x12, 0x34, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{ID: 0x1234},
			12,
			false,
		},
		{
			"QR set",
			args{[]byte{0, 0, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{QR: true},
			12,
			false,
		},
		{
			"Opcode set",
			args{[]byte{0, 0, 0x78, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{Opcode: 0x0f},
			12,
			false,
		},
		{
			"AA set",
			args{[]byte{0, 0, 0x04, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{AA: true},
			12,
			false,
		},
		{
			"TC set",
			args{[]byte{0, 0, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{TC: true},
			12,
			false,
		},
		{
			"RD set",
			args{[]byte{0, 0, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{RD: true},
			12,
			false,
		},
		{
			"RA set",
			args{[]byte{0, 0, 0, 0x80, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{RA: true},
			12,
			false,
		},
		{
			"RCode set",
			args{[]byte{0, 0, 0, 0x0f, 0, 0, 0, 0, 0, 0, 0, 0}},
			Header{RCode: 0x0f},
			12,
			false,
		},
		{
			"QDCount set",
			args{[]byte{0, 0, 0, 0, 0x12, 0x34, 0, 0, 0, 0, 0, 0}},
			Header{QDCount: 0x1234},
			12,
			false,
		},
		{
			"ANCount set",
			args{[]byte{0, 0, 0, 0, 0, 0, 0x12, 0x34, 0, 0, 0, 0}},
			Header{ANCount: 0x1234},
			12,
			false,
		},
		{
			"NSCount set",
			args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34, 0, 0}},
			Header{NSCount: 0x1234},
			12,
			false,
		},
		{
			"ARCount set",
			args{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34}},
			Header{ARCount: 0x1234},
			12,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotHeader Header
			got, err := HeaderUnmarshal(&gotHeader, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HeaderUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantHeader, gotHeader)
			assert.Equal(t, tt.want, got)
		})
	}
}
