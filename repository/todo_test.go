package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursorRoundTrip(t *testing.T) {
	// Given
	item := &TodoItem{ID: &NullableID{Value: "7"}}
	const index = 3

	// When
	encoded := item.EncodeCursor(index)
	decoded, err := DecodeCursor(encoded)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, index, decoded.I)
	assert.Equal(t, item.ID.Value, decoded.TodoID)
}
