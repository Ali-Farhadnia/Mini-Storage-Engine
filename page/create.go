package page

func New(pID uint64, pType PageType) (*Page, error) {
	if pType != PageTypeInternal && pType != PageTypeLeaf {
		return nil, ErrInvalidType
	}

	data := make([]byte, pageSize)

	p := &Page{
		data: data,
	}

	p.setPageID(pID)
	p.setPageType(pType)
	p.setKeyCount(0)

	return p, nil
}
