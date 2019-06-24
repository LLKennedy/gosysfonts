package fontapi

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	err := Connect()
	assert.NoError(t, err)
}