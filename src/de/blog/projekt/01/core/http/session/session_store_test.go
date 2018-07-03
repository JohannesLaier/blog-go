package session

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"os"
	"encoding/json"
	"time"
	"net/http"
	"net/http/httptest"
)

func TestSessionStore_New(t *testing.T) {
	// Create a new session store containing one session
	store := NewSessionStore()
	id, session := store.New()

	// Session and the id of it should not be null
	assert.NotNil(t, id, "Did not return a valid session id")
	assert.NotNil(t, session, "Did not return a valid session store object")
}

func TestGetSessionStore(t *testing.T) {
	// Take twice a session store object
	store1 := GetSessionStore()
	store2 := GetSessionStore()

	// They should be equal
	assert.Equal(t, store1, store2)
}

func TestSessionStore_Get(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create a new session store containing one session
	store := NewSessionStore()
	id, session := store.New()

	// Select the session by the given id
	session2 := store.Get(id)

	// Selected session and the pointer  should be the same
	assert.Equal(t, session, session2, "Session must be equal")
	assert.Equal(t, *session, *session2, "Session must be equal")

	// Invalid session it have to return nil
	session3 := store.Get("invalid-id")
	assert.NotEqual(t, session, session3, "Session must not be equal")
	assert.NotEqual(t, session2, session3, "Session must not be equal")
}

func TestSessionStore_GetCurrent(t *testing.T) {

	session_key := "CURRENT"
	session_value := "john@mail.com"

	// Create a new session store containing one session
	store := NewSessionStore()
	id, session1 := store.New()

	// Put Current into session value
	session1.Put(session_key, session_value)

	// Manually generate cookie
	expiration := time.Now().Add(15 * time.Minute)
	cookie := http.Cookie{Name: "SESSIONID", Value: id, Expires: expiration}

	// Create request
	r1 := httptest.NewRequest("POST", "/", nil)

	// Get current session out of request
	session2 := store.GetCurrent(r1)

	// Session should not be found
	assert.Nil(t, session2)

	// Add the session cookie to the request
	r1.AddCookie(&cookie)

	// Get current session out of request
	session3 := store.GetCurrent(r1)

	// They should be equal
	assert.Equal(t, session1.Id, session3.Id)
	assert.Equal(t, session1.Date, session3.Date)
	assert.Equal(t, session1.Data, session3.Data)
}

func TestSessionStore_Discard(t *testing.T) {
	// Create a session store
	store := NewSessionStore()

	// Create two new sessions
	id1, session1 := store.New()
	id2, session2 := store.New()

	// Both sessions are in the store?
	assert.Equal(t, session1, store.Get(id1), "Session must be equal")
	assert.Equal(t, session2, store.Get(id2), "Session must be equal")

	// Delete first session
	store.Discard(id1)

	// First session was deleted second is existing anyway
	assert.Nil(t, store.Get(id1), "Session must not existing anymore")
	assert.NotNil(t, store.Get(id2), "Deleted the wrone session")

	// Discard the second session
	store.Discard(id2)

	// Both sessions should be discarded
	assert.Nil(t, store.Get(id1), "Session must not existing anymore")
	assert.Nil(t, store.Get(id2), "Deleted the wrone session")
}

func TestSessionStore_Load(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create a session store
	store1 := NewSessionStore()

	// Create a new session
	id, session1 := store1.New()

	// Add Values to session
	session1.Put("my-key-int", 12345);
	session1.Put("my-key-string", "value");
	session1.Put("my-key-bool", true);

	// Save store to disk
	store1.Store()

	// Create a second session store
	store2 := NewSessionStore()

	// Load store content from disk
	store2.Load()

	// Fetch session from store
	session2 := store2.Get(id)

	assert.Equal(t,12345.0, session2.Get("my-key-int"),"Session content missmatchs")
	assert.Equal(t,"value", session2.Get("my-key-string"),"Session content missmatchs")
	assert.Equal(t,true, session2.Get("my-key-bool"),"Session content missmatchs")
}

func TestSessionStore_Store(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create a session store
	store := NewSessionStore()

	// Create a new session
	_, session := store.New()

	// Add Values to session
	session.Put("my-key", 12345);

	// Store the session store to disk
	store.Store()

	// Read the value of the session store
	content_binary, _ := ioutil.ReadFile("res/sessions/sessions.json")
	content_string := string(content_binary)

	// Check the session store value
	json_str, _ := json.Marshal(store.Sessions)
	assert.Equal(t, string(json_str), content_string)
}

func TestSessionStore_getPath(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	// Create a new session store
	store := NewSessionStore()

	// Checks the path
	assert.Equal(t, store.getPath(), "res/sessions/sessions.json", "wrong path")
}