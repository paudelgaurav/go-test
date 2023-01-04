package gotest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateStr(t *testing.T) {
	dupSlice := []string{"a", "a", "b", "c"}
	newSlice := RemoveDuplicateStr(dupSlice)
	assert.Equal(t, newSlice, []string{"a", "b", "c"})
	assert.NotEqual(t, dupSlice, newSlice)
}
