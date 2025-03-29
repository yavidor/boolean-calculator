package main

import "testing"

func Test_xor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		A    bool
		B    bool
		want bool
	}{
		struct {
			name string
			A    bool
			B    bool
			want bool
		}{
			name: "TT",
			A:    true,
			B:    true,
			want: false,
		},
		struct {
			name string
			A    bool
			B    bool
			want bool
		}{
			name: "TF",
			A:    true,
			B:    false,
			want: true,
		},
		struct {
			name string
			A    bool
			B    bool
			want bool
		}{
			name: "FT",
			A:    false,
			B:    true,
			want: true,
		},
		struct {
			name string
			A    bool
			B    bool
			want bool
		}{
			name: "FF",
			A:    false,
			B:    false,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xor(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("xor(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}
