package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

func PasswordEncrypt(plainText string) (string, error) {
	encryptString, err := gmd5.EncryptString(plainText)
	if err != nil {
		return "", err
	}
	return encryptString, nil
}
