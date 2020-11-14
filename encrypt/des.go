package encrypt

import (
	"crypto/cipher"
	"crypto/des"
	"log"
	"runtime"
)

func DesCbcEncrypt(plainText, key []byte, ivDes ...byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, ErrNot8BytesKey
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != 8 {
			return nil, ErrNot8BytesIv
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(iv8)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)

	return cipherText, nil
}

func DesCbcDecrypt(cipherText, key []byte, ivDes ...byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, ErrNot8BytesKey
	}

	block, err := des.NewCipher(key)
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
	if len(ivDes) != 0 {
		if len(ivDes) != 8 {
			return nil, ErrNot8BytesIv
		} else {
			iv = ivDes
		}
	} else {
		iv = []byte(iv8)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)

	unPaddingText, err := PKCS5UnPadding(plainText)
	if err != nil {
		return nil, err
	}

	return unPaddingText, nil
}
