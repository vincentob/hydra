package encrypt

import "errors"

var (
	ErrNot8BytesKey  = errors.New("8 bytes secret key is required")
	ErrNot8BytesIv   = errors.New("8 bytes iv is required")
	ErrNot16BytesKey = errors.New("16, 24 or 32 bytes secret key is required")
	ErrNot16BytesIv  = errors.New("16 bytes iv is required")
	ErrPaddingSize   = errors.New("padding size error please check the secret key or iv")

	iv16 = "vincent612345678"
	iv8  = "vincent6"
)
