package encrypt

import "bytes"

// PKCS5Padding
func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...)
}

func PKCS5UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length {
		return nil, ErrPaddingSize
	}
	return plainText[:length-number], nil
}

// TODO:
// PKCS7Padding
// PKCS7UnPadding
// ZeroPadding
