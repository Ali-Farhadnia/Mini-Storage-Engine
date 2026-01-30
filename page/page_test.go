package page_test

import (
	"testing"

	"github.com/Ali-Farhadnia/Mini-Storage-Engine/page"
	"github.com/stretchr/testify/assert"
)

func TestCreation(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		p, err := page.New(1, page.PageTypeInternal)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), p.PageID())
		assert.Equal(t, page.PageTypeInternal, p.PageType())
		assert.Equal(t, uint16(0), p.KeyCount())
		assert.Equal(t, 4085, len(p.Data()))
	})

	t.Run("wrong page type", func(t *testing.T) {
		p, err := page.New(1, 5)
		assert.ErrorIs(t, err, page.ErrInvalidType)
		assert.Nil(t, p)
	})

}
