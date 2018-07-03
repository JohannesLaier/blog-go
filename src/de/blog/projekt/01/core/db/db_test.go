package db

import (
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"io/ioutil"
	"testing"
	"os"
)

func TestNewDB(t *testing.T) {
	db := NewDB("my-db")
	assert.NotNil(t, db)
}

func TestGet(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	db_name := "my-db"
	db1 := Get(db_name)
	db2 := Get(db_name)
	assert.Equal(t, db1, db2)
}

func TestDB_AddCollection(t *testing.T) {
	db := NewDB("blog")

	collection_name := "posts"
	collection := NewDBCollection(collection_name)

	db.AddCollection(collection)

	assert.Equal(t, collection, db.GetCollection(collection_name), "Collection not found")
}

func TestDB_GetCollection(t *testing.T) {
	db := NewDB("blog")
	collection_name := "posts"

	collection := NewDBCollection(collection_name)

	db.AddCollection(collection)

	assert.Equal(t, collection, db.GetCollection(collection_name), "Collection not found")
}

func TestDB_Load(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	// Create database
	db := NewDB("blog")
	collection_name := "posts"

	// incl. collection
	collection := NewDBCollection(collection_name)

	// Fill collection
	collection.Add(TestDummy{1, "value1"});
	collection.Add(TestDummy{2, "value2"});
	collection.Add(TestDummy{3, "value3"});

	// Add collection to database
	db.AddCollection(collection)

	// Store database to disc
	db.Store()

	// Load db from disk
	db_load := NewDB("blog")
	db_load.Load()

	// Fetch collection
	collection_load := db_load.GetCollection(collection_name)

	// Fetch collection values
	dummies := collection_load.GetAll()

	// Check values
	assert.Equal(t, 3, len(dummies), "Collection content missmatchs")
}

func TestDB_Store(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	// DB and table names
	db_name := "blog"
	collection_name := "dummy"

	// Create dumy content to store
	dummy := TestDummy{1, "value"}

	// Add new collection
	collection := NewDBCollection(collection_name)
	collection.Add(dummy)

	// Create database
	db := NewDB(db_name)
	db.AddCollection(collection)

	// Store database to disk
	db.Store()

	// Load Database from disk
	db_load := NewDB(db_name)
	db_load.Load()

	// Read the value of the session store
	content_binary, _ := ioutil.ReadFile("res/db/"+db_name+".json")
	content_string := string(content_binary)

	// Check the session store value
	json_str, _ := json.Marshal(db.Collections)
	assert.Equal(t, string(json_str), content_string)
}

func TestDB_GetPath(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	db_name := "mydb"
	db_path := "res/db/mydb.json"

	db := NewDB(db_name)
	assert.Equal(t, db_path, db.getPath())
}