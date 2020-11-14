package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"runtime"
)

// AesCbcEncrypt
func AesCbcEncrypt(plainText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrNot16BytesKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddingText := PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrNot16BytesIv
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(iv16)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// decrypt
func AesCbcDecrypt(cipherText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrNot16BytesKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key or text is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrNot16BytesIv
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(iv16)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText, err := PKCS5UnPadding(paddingText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// AesCtrEncrypt
func AesCtrEncrypt(plainText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrNot16BytesKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrNot16BytesIv
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(iv16)
	}

	stream := cipher.NewCTR(block, iv)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

func AesCtrDecrypt(cipherText, key []byte, ivAes ...byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrNot16BytesKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != 16 {
			return nil, ErrNot16BytesIv
		} else {
			iv = ivAes
		}
	} else {
		iv = []byte(iv16)
	}

	stream := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
