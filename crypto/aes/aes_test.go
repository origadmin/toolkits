package aes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// SecretKey Define aes secret key 2^5
var SecretKey = []byte("2985BCFDB5FE43129843DB59825F8647")

func TestAESEncrypt(t *testing.T) {
	assert := assert.New(t)

	data := []byte("hello world")

	bs64, err := EncryptToBase64(data, SecretKey)
	assert.Nil(err)
	assert.NotEmpty(bs64)

	t.Log(bs64)

	result, err := DecryptFromBase64(bs64, SecretKey)
	assert.Nil(err)
	assert.Equal(data, result)
}
