package internal

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	t.Run("Valid request/response, no template", func(t *testing.T) {
		dto := MemberDto{
			Name:  "Jimmy",
			Email: "jimmy@mail.com",
		}
		b, err := json.Marshal(dto)
		reader := bytes.NewReader(b)
		request, err := http.NewRequest(http.MethodPost, "", reader)
		writer := httptest.NewRecorder()
		assert.NoError(t, err)
		indexHandler(writer, request)
		assert.NotNil(t, writer.Code)
		assert.Equal(t, writer.Code, http.StatusInternalServerError)
	})
}

func TestLoggerResponse(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(""))
		code := http.StatusOK
		err := loggerResponse(request, code)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}

func TestParseMemberRequest(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		dto := MemberDto{
			Name:  "Jimmy",
			Email: "jimmy@mail.com",
		}
		b, err := json.Marshal(dto)
		reader := bytes.NewReader(b)
		request, err := http.NewRequest(http.MethodPost, "", reader)
		assert.NoError(t, err)
		obj, err := parseMemberRequest(request.Body)
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.NotNil(t, obj)
	})
}

func TestResponse(t *testing.T) {
	t.Run("No template", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(""))
		writer := httptest.NewRecorder()
		err := response(writer, request)
		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}
