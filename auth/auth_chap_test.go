package auth

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"testing"
)

func TestChallenge(t *testing.T) {
	type fields struct {
		ChallengeString string
		Hash            hash.Hash
	}
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"default test",
			args{
				0,
			},
		},
		{
			"challenge with 16 length",
			args{
				16,
			},
		},
	}
	// test nonce generator
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Challenge{}
			if tt.args.length != 0 {
				c.New(tt.args.length)
			}
			t.Log(c)
		})
	}
	// test validation
	t.Run("validation test", func(t *testing.T) {
		secret := "123456"
		h := md5.New()
		c := &Challenge{}
		c.New(8)
		t.Log("Challenge: ", c.ChallengeString)
		io.WriteString(h, secret+c.String())
		res := h.Sum(nil)
		resStr := string(res)
		t.Log("Correct response: ", resStr, "\n")
		encodeStr := fmt.Sprintf("%x", res)
		t.Log("Encoded md5 string: ", encodeStr)
		// compare result
		if result := c.Validate(secret, encodeStr); result != true {
			t.Errorf("Validation test failed. challenge string: %v, accpeted response: %v", c.ChallengeString, encodeStr)
		}
	})
}
