package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5Str(t *testing.T) {
	md5, err := Md5Str("A")
	rx := "^[a-fA-F0-9]{32}$"

	assert.Regexp(t, rx, md5)
	assert.Nil(t, err)

	md5 = Sha1Str("A")
	assert.NotRegexp(t, rx, md5)
}

func TestSha1Str(t *testing.T) {
	sha1Str := Sha1Str("A")
	assert.Regexp(t, "^[a-fA-F0-9]{40}$", sha1Str)
}
