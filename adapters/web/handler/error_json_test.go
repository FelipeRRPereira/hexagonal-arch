package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "test error message"
	result := jsonError(msg)
	require.Equal(t, `{"message":"test error message"}`, string(result), "jsonError should return the correct JSON format")
}
