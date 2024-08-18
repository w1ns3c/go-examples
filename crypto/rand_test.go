package crypto

import (
	"testing"
)

func TestGenRandStr(t *testing.T) {
	tests := []struct {
		name string
		n    int // string length
		//wantRandStr string
		wantErr bool
	}{
		{
			name: "n=16",
			n:    16,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
		{
			name: "n=27",
			n:    27,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
		{
			name: "n=0",
			n:    0,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRandStr, err := GenRandStr(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRandStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotRandStr) != tt.n {
				t.Errorf("GenRandStr() len = %v, want %v", len(gotRandStr), tt.n)
				return
			}
		})
	}
}

func TestGenRandSlice(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		n    int // string length
		//want    []byte
		wantErr bool
	}{
		{
			name: "n=16",
			n:    16,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
		{
			name: "n=27",
			n:    27,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
		{
			name: "n=0",
			n:    0,
			//wantRandStr: "0123456789",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenRandSlice(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenRandSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != tt.n {
				t.Errorf("GenRandSlice() len = %v, want %v", len(got), tt.n)
				return
			}
		})
	}
}
