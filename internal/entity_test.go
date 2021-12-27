package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberDto_IsValid(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		obj := MemberDto{
			Name:  "Jimmy",
			Email: "jimmy@mail.com",
		}
		Members.members = make(map[string]Member)
		err := obj.IsValid()
		assert.NoError(t, err)
		assert.Nil(t, err)
	})
	t.Run("Invalid name", func(t *testing.T) {
		obj := MemberDto{
			Name:  "Джимми =)",
			Email: "jimmy@mail.com",
		}
		Members.members = make(map[string]Member)
		err := obj.IsValid()
		assert.Error(t, err)
		assert.Equal(t, "invalid name", err.Error())
		assert.NotNil(t, err)
	})
	t.Run("Invalid email", func(t *testing.T) {
		obj := MemberDto{
			Name:  "Jimmy",
			Email: "jiailcm",
		}
		Members.members = make(map[string]Member)
		err := obj.IsValid()
		assert.Error(t, err)
		assert.Equal(t, "invalid email", err.Error())
		assert.NotNil(t, err)
	})
	t.Run("Invalid email exists", func(t *testing.T) {
		obj := MemberDto{
			Name:  "Jimmy",
			Email: "jimmy@mail.com",
		}
		Members.members = make(map[string]Member)
		Members.Add(*NewMember(obj))
		err := obj.IsValid()
		assert.Error(t, err)
		assert.Equal(t, "email already exists", err.Error())
		assert.NotNil(t, err)
	})
}
