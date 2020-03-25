package auth

import "testing"

func TestSalting(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test salting",
			args: args{
				"Sample text",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotSalt := Salting(tt.args.text)
			t.Logf("\nPlain text: %s\n", tt.args.text)
			t.Logf("\nGot hashed text: %s.\nGot salt: %s.\n", gotResult, gotSalt)
			// if gotResult != tt.wantResult {
			// 	t.Errorf("Salting() gotResult = %v, want %v", gotResult, tt.wantResult)
			// }
		})
	}
}
