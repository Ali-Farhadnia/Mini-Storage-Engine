package page

// Data returns a slice to the data region (after the header).
func (p *Page) Data() []byte {
	return p.data[HeaderSize:]
}
