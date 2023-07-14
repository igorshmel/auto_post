package errs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Init(t *testing.T) {
	err := New().SetCode(Internal).SetMsg("Test")
	assert.NotNil(t, err.Error())
	assert.Equal(t, err.String(), "code: ERR_INTERNAL, message: Test")
	assert.Equal(t, string(err.Marshal()), "{\"error\":{\"code\":\"ERR_INTERNAL\",\"message\":\"Test\"}}")
}

func TestFromBytes(t *testing.T) {
	responseWithError := []byte(`{"error":{"code":"TEST_CODE","message":"test message"}}`)
	responseOk := []byte(`{"item": {"foo": 1, "bar": 2}}`)

	err := FromBytes(responseWithError)
	assert.Error(t, err.Error())
	assert.Equal(t, "TEST_CODE", err.Code)
	assert.Equal(t, "test message", err.Message)

	err = FromBytes(responseOk)
	assert.Nil(t, err)
}

func TestFromBody(t *testing.T) {
	responseWithError := bytes.NewBufferString(`{"error":{"code":"TEST_CODE","message":"test message"}}`)
	responseOk := bytes.NewBufferString(`{"item": {"foo": 1, "bar": 2}}`)

	err := FromBody(responseWithError)
	assert.Error(t, err.Error())
	assert.Equal(t, "TEST_CODE", err.Code)
	assert.Equal(t, "test message", err.Message)

	err = FromBody(responseOk)
	assert.Nil(t, err)
}
