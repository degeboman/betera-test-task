package models

import "io"

type ImageUnitCore struct {
	Payload     io.Reader
	PayloadName string
	PayloadSize int64
}
