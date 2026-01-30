package page

import (
	"errors"
)

type PageType uint8

const (
	PageTypeUnknown  PageType = 0
	PageTypeInternal PageType = 1
	PageTypeLeaf     PageType = 2
)

const (
	pageSize = 4096

	headerPageIDSize   = 8 // uint64
	headerPageTypeSize = 1 // uint8
	headerKeyCountSize = 2 // uint16

	HeaderSize = headerPageIDSize +
		headerPageTypeSize +
		headerKeyCountSize
)

var (
	ErrInvalidType = errors.New("invalid page type")
)

// Page represents a fixed-size storage unit.
type Page struct {
	data []byte
}
