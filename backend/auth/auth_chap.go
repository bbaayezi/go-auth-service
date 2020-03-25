package auth

import (
	// import md5 hash funtion
	"crypto/md5"
	"fmt"
	"go-auth-service/util"
	"hash"
	"io"
	"strings"
)

// this file implements challenge handshake authentication protocol

type Challenge struct {
	ChallengeString string
	Hash            hash.Hash
}

// New generates new challenge string
func (c *Challenge) New(length int) {
	// using md5 as default hash function
	c.Hash = md5.New()
	// generate nonce
	nonce := util.GenerateNonce(length)
	c.ChallengeString = nonce
}

// NewHashFunc allows user to set custom hash function
func (c *Challenge) NewHashFunc(hash hash.Hash) {
	c.Hash = hash
}

// Validate is the function to validate challenge
func (c *Challenge) Validate(secret string, challengeResponse string) bool {
	// write secret to hash
	io.WriteString(c.Hash, secret+c.ChallengeString)
	// decode hex md5 to string
	decoded := fmt.Sprintf("%x", c.Hash.Sum(nil))
	// compares the challenge result
	if strings.Compare(decoded, challengeResponse) == 0 {
		return true
	}
	return false
}

func (c *Challenge) String() string {
	if c.ChallengeString == "" {
		// generate challenge of length 8 as default
		c.New(8)
	}
	return c.ChallengeString
}

// NewChallenge is a convient way to generate a challenge object
func NewChallenge(length int) *Challenge {
	challenge := &Challenge{}
	challenge.New(length)
	return challenge
}
