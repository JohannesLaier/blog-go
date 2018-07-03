package author

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthor(t *testing.T) {
	username := "admin"
	password := "s3cr3t"

	a := NewAuthor(username, password)

	assert.Equal(t, username, a.GetUsername(), "Incorrect username")
	assert.True(t, a.Verify(password), "Password didnt match")
	assert.False(t, a.Verify("incorrectpwd"), "Incorrect password passed the auth")
}

func TestAuthor_Verify(t *testing.T) {
	password := "s3cr3t"

	a := NewAuthor("admin", password)

	assert.True(t, a.Verify(password), "Password didn't match: " + password)
}

func TestAuthor_generateSalt(t *testing.T) {
	a := NewAuthor("admin", "s3cr3t")
	salt := a.generateSalt()

	// Testen ob String die Länge 20 hat?
	assert.Equal(t, len(salt), 20 , "Wrong length of Salt")
}

func TestAuthor_hashPassword(t *testing.T) {
	a := NewAuthor("admin", "s3cr3t")
	hash := a.hashPassword("s3cr3t")

	assert.NotEqual(t, "s3cr3t", hash, "Password wasn't hashed")
}

func TestAuthor_GetUsername(t *testing.T) {
	username := "admin"

	a := NewAuthor(username, "s3cr3t")

	assert.Equal(t, username, a.GetUsername(),"Could not set the username")
}

func TestAuthor_GetID(t *testing.T) {
	id := 12345

	a := NewAuthor("admin", "s3cr3t")
	a.SetID(12345)

	assert.Equal(t, id, a.GetID(), "ID doesn't match")
}

func TestAuthor_SetPassword(t *testing.T) {
	password := "s3cr3t"

	a := NewAuthor("admin", password)

	assert.True(t, a.Verify(password), "Password didn't match: " + password)

	newPassword := "123456"
	a.SetPassword(newPassword)

	assert.False(t, a.Verify(password), "Still possible to log in with old password")
	assert.True(t, a.Verify(newPassword), "Not possible to log in with new password")
}

func TestAuthor_SetID(t *testing.T) {

	a := NewAuthor("admin", "s3cr3t")

	newId := 54321
	a.SetID(newId)

	assert.Equal(t, newId, a.GetID(), "new ID wasn't set")
}

func TestAuthor_SetUsername(t *testing.T) {
	username := "peter"
	password := "mypasswd"

	a := NewAuthor(username, password)

	assert.Equal(t, a.GetUsername(), username, "Username doesnt match")

	newUsername := "max"

	a.SetUsername(newUsername)

	assert.Equal(t, a.GetUsername(), newUsername, "Username doesnt match")
}
