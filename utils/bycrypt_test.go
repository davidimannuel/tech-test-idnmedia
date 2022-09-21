package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBycrypt(t *testing.T) {
	secret := "secret123"
	enc := BcryptEncrypt(secret)
	t.Log("enc", enc)
	assert.NotEmpty(t, enc)
	dec := BcryptCompare(enc, secret)
	assert.True(t, dec)
}
