package encoding

import "encoding"

type ValidData interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
