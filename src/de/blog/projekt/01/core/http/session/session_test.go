package session

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
	"time"
	"os"
)

func TestSession_Get(t *testing.T) {
	// Create a new session
	session := NewSession("my-id")

	// Insert values
	session.Put("firstName", "Joe")
	session.Put("lastName", "Doe")
	session.Put("age", 30)

	// Check the values
	assert.Equal(t, "Joe", session.Get("firstName"), "Firstname does not match.")
	assert.Equal(t, "Doe", session.Get("lastName"), "Lastname does not match.")
	assert.Equal(t, 30, session.Get("age"), "Age does not match.")

	// Invalid keys
	assert.Nil(t, session.Get("invalid-key"), "Invalid key has to return nil.")
}

func TestSession_Get_Expired(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	session_key := "test"
	session_value := 1

	// Create a new session
	session := NewSession("my-id")

	// Insert values
	session.Put(session_key, session_value)

	// Not expired date
	time_future := time.Now()
	time_future = time_future.Add(-14 * time.Minute)

	// Set new session date
	session.Date = time_future

	// Session is not expired
	assert.Equal(t, session_value, session.Get(session_key))

	// Expired date
	time_future = time_future.Add(-20 * time.Minute)

	// Set new session date
	session.Date = time_future

	// Session should be expired
	assert.Nil(t, session.Get(session_key))
}

func TestSession_Put(t *testing.T) {
	// Create a new session
	session := NewSession("my-id")

	// Insert values
	session.Put("salary", 150000.00)
	session.Put("cto", true)

	// Check Values
	assert.Equal(t, 150000.00, session.Get("salary"), "Salary does not match.")
	assert.Equal(t, true, session.Get("cto"), "CTO-Flag does not match.")
}

func TestSession_Remove(t *testing.T) {
	// Create a new session
	session := NewSession("my-id")

	// Insert values
	session.Put("salary", 150000.00)
	session.Put("cto", true)

	// Check Values
	assert.Equal(t, 150000.00, session.Get("salary"), "Salary does not match.")
	assert.Equal(t, true, session.Get("cto"), "CTO-Flag does not match.")

	// Delete sallary value
	session.Remove("salary")

	// Sallary should be deleted
	assert.Nil(t, session.Get("salary"), "Salary has been removed.")
	assert.Equal(t, true, session.Get("cto"), "CTO-Flag does not match.")

	// Delete cto flag
	session.Remove("cto")

	// Both should be deleted
	assert.Nil(t, session.Get("salary"), "Salary has been removed.")
	assert.Nil(t, session.Get("cto"), "CTO-Flag has been removed.")

}

func TestSession_CreateCookie(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create responses
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	// Create session and add cookie to response
	session := NewSession("my-id")
	session.CreateCookie(w1)

	// Manually generate cookie
	expiration := time.Now().Add(15 * time.Minute)
	cookie := http.Cookie{Name: "SESSIONID", Value: session.Id, Expires: expiration}

	// Set cookie to second response
	http.SetCookie(w2, &cookie)

	// Response-Headers should be equal
	assert.Equal(t, w1.Header().Get("Set-Cookie"), w2.Header().Get("Set-Cookie"),"empty cookie")
}

func TestSession_DestroyCookie(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create responses
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	// Create request
	r1 := httptest.NewRequest("POST", "/", nil)

	// Create new session
	session := NewSession("my-id")

	// And cookie of existing session
	expiration := time.Now()
	cookie := http.Cookie{Name: "SESSIONID", Value: session.Id, Expires: expiration}

	// Add cookie to request to fake existing session
	r1.AddCookie(&cookie)

	// Destory existing session
	session.DestroyCookie(w1, r1)

	// Do the same on the dummy session
	expiration = time.Unix(0, 0)
	cookie2 := http.Cookie{Name: "SESSIONID", Value: session.Id, Expires: expiration}
	http.SetCookie(w2, &cookie2)

	// Both should be equal
	assert.Equal(t, w1.Header().Get("Set-Cookie"), w2.Header().Get("Set-Cookie"),"empty cookie")
}