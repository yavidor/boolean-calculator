package main

import "testing"

type testCase struct {
	name string
	A    bool
	B    bool
	want bool
}

func generateTruthTable(TT, TF, FT, FF bool) []testCase {
	return []testCase{
		{
			name: "TT",
			A:    true,
			B:    true,
			want: TT,
		},
		{
			name: "TF",
			A:    true,
			B:    false,
			want: TF,
		},
		{
			name: "FT",
			A:    false,
			B:    true,
			want: FT,
		},
		{
			name: "FF",
			A:    false,
			B:    false,
			want: FF,
		},
	}
}

func Test_and(t *testing.T) {
	tests := generateTruthTable(true, false, false, false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := and(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("and() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_or(t *testing.T) {
	tests := generateTruthTable(true, true, true, false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := or(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xor(t *testing.T) {
	tests := generateTruthTable(false, true, true, false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xor(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("xor(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}

func Test_nand(t *testing.T) {
	tests := generateTruthTable(false, true, true, true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nand(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("nand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nor(t *testing.T) {
	tests := generateTruthTable(false, false, false, true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nor(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("nor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xnor(t *testing.T) {
	tests := generateTruthTable(true, false, false, true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xnor(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("xnor(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}
