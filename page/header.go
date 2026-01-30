package page

import "encoding/binary"

func (p *Page) PageID() uint64 {
	return binary.LittleEndian.Uint64(p.data[0:8])
}

func (p *Page) PageType() PageType {
	return PageType(p.data[8])
}

func (p *Page) KeyCount() uint16 {
	return binary.LittleEndian.Uint16(p.data[9:11])
}

func (p *Page) setPageID(id uint64) {
	binary.LittleEndian.PutUint64(p.data[0:8], id)
}

func (p *Page) setPageType(typ PageType) {
	p.data[8] = byte(typ)
}

func (p *Page) setKeyCount(count uint16) {
	binary.LittleEndian.PutUint16(p.data[9:11], count)
}
