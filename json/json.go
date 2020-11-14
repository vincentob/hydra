package json

import "github.com/json-iterator/go"

var (
	JJ     = jsoniter.ConfigCompatibleWithStandardLibrary // 100% compatible with standard library behavior
	FastJJ = jsoniter.ConfigFastest                       // ConfigFastest marshals float with only 6 digits precision
)
