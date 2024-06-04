package crypto

import "testing"

func TestGenRandStr(t *testing.T) {
	tests := []struct {
		name string
		n    int // len of wanted string
	}{
		// TODO: Add test cases.
		{
			name: "Compare string len",
			n:    10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := GenRandStr(tt.n)
			if err != nil {
				t.Errorf("GenRandStr() error = %v", err)
			}

			if len(gotS) != tt.n {
				t.Errorf("GenRandStr() len(got_string) = %v, want %v", len(gotS), tt.n)
			}
		})
	}
}
