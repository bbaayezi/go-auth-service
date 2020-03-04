package util

import "testing"

func TestGenerateNonce(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test 8 bytes",
			args{
				8,
			},
		},
		{
			"test 16 bytes",
			args{
				16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateNonce(tt.args.length)
			t.Log("Nonce: ", got)
		})
	}
}
