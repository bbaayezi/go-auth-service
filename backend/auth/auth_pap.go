package auth

import (
	"crypto/md5"
	"fmt"
	"go-auth-service/entity"
	"go-auth-service/util"
	"io"
)

// Salting generates salt and append it to the text, it returns the hashed text and salt
func Salting(text string) (result string, salt string) {
	// encode the text with random generated salt
	// generate a 8 bit nonce
	s := util.GenerateNonce(8)
	// append the salt to the end of text
	newText := text + s
	// using md5 algorithm to hash the text
	hash := md5.New()
	io.WriteString(hash, newText)
	hashText := fmt.Sprintf("%x", hash.Sum(nil))
	// return the result
	return hashText, s
}

func PasswordAuth(user entity.User) bool {
	return false
}
