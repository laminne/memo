package memo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDir(t *testing.T) {
	expected := []File{
		{
			Path:     "test/test1.txt",
			Contents: []byte("Hello world, this is test 1."),
			Size:     28,
		},
		{
			Path:     "test/test2/test2.txt",
			Contents: []byte("Hello world, this is test 2."),
			Size:     28,
		},
	}

	res, _ := ReadDir("test", []string{"ignore"})
	assert.Equal(t, res, expected, "")
}
