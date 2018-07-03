package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type TestDummy struct {
	Id int
	Data string
}

func (d TestDummy) GetID() int {
	return d.Id
}

func (d TestDummy) SetID(id int) {
	d.Id = id
}

type TestDummy2 struct {
	Id int
	Data string
	Value int
}

func (d TestDummy2) GetID() int {
	return d.Id
}

func (d TestDummy2) SetID(id int) {
	d.Id = id
}

func TestNewDBCollection(t *testing.T) {
	collection := NewDBCollection("my-collection")
	assert.NotNil(t, collection)
}

func TestDBCollection_GetName(t *testing.T) {
	name := "my-collection"
	collection := NewDBCollection(name)
	assert.Equal(t, name, collection.GetName(), "Name of collection does not match")
}

func TestDBCollection_GetAll(t *testing.T) {
	collection := NewDBCollection("TestDummy")

	data := []TestDummy{
		TestDummy{1,"Huhu"},
		TestDummy{2, "Test"},
		TestDummy{3, "ABC"},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	results := collection.GetAll()

	for index, _ := range data {
		assert.Equal(t, data[index], results[index], "Elements of collections doesnt match")
	}
}

func TestDBCollection_GetFilter(t *testing.T) {
	collection := NewDBCollection("TestDummy2")

	data := []TestDummy2{
		TestDummy2{1, "Huhu", 1},
		TestDummy2{2, "Test", 5},
		TestDummy2{3, "ABC", 10},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	result := collection.GetFilter(func(entry DBCollectionEntry) bool {
		return ((entry).(TestDummy2)).Value == 5
	})

	assert.Equal(t, data[1], result, "Could not find the correct element.")
}

func TestDBCollection_GetFilter_Nothing_Found(t *testing.T) {
	collection := NewDBCollection("TestDummy2")

	data := []TestDummy2{
		TestDummy2{1, "Huhu", 1},
		TestDummy2{2, "Test", 5},
		TestDummy2{3, "ABC", 10},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	result := collection.GetFilter(func(entry DBCollectionEntry) bool {
		return ((entry).(TestDummy2)).Value == 100
	})

	assert.Nil(t, result, "There should be nothing")
}

func TestDBCollection_GetListFilter(t *testing.T) {
	collection := NewDBCollection("TestDummy2FilterList")

	data := []TestDummy2{
		TestDummy2{1, "Huhu", 1},
		TestDummy2{2, "Test", 5},
		TestDummy2{3, "ABC", 10},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	results := collection.GetListFilter(func(entry DBCollectionEntry) bool {
		return (entry).(TestDummy2).Value > 4
	})

	assert.Equal(t, data[1], results[0], "Could not find the correct element.")
	assert.Equal(t, data[2], results[1], "Could not find the correct element.")
}

func TestDBCollection_GetByID(t *testing.T) {
	collection := NewDBCollection("TestDummy2GetByID")

	data := []TestDummy2{
		TestDummy2{1, "Huhu", 1},
		TestDummy2{2, "Test", 5},
		TestDummy2{3, "ABC", 10},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	result0 := collection.GetByID(data[0].Id)
	result1 := collection.GetByID(data[1].Id)
	result2 := collection.GetByID(data[2].Id)


	assert.Equal(t, data[0], result0, "Could not find the correct element.")
	assert.Equal(t, data[1], result1, "Could not find the correct element.")
	assert.Equal(t, data[2], result2, "Could not find the correct element.")
}

func TestDBCollection_GetByID_NotFound(t *testing.T) {
	collection := NewDBCollection("TestDummy2GetByID")

	data := []TestDummy2{
		TestDummy2{1, "Huhu", 1},
		TestDummy2{2, "Test", 5},
		TestDummy2{3, "ABC", 10},
	}

	for _, entry := range data {
		collection.Add(entry)
	}

	result := collection.GetByID(150)

	assert.Nil(t, result)
}

func TestDBCollection_Add(t *testing.T) {
	dummy := TestDummy{1, "Test Value"}

	collection := NewDBCollection("dummy")

	assert.Equal(t, 0, len(collection.GetAll()))

	collection.Add(dummy)

	assert.Equal(t, 1, len(collection.GetAll()))
}

func TestDBCollection_Update(t *testing.T) {
	dummy := TestDummy{1, "Test Value"}
	dummy2 := TestDummy{2, "Test Value"}

	collection := NewDBCollection("dummy")
	collection.Add(dummy)
	collection.Add(dummy2)

	assert.Equal(t, dummy, collection.GetAll()[0], "Could not find element in Database")


	dummy.Data = "This is trash content"

	assert.NotEqual(t, dummy, collection.GetAll()[0], "Element must be diferent")
	assert.Equal(t, dummy2, collection.GetAll()[1], "Element must be equal")

	collection.Update(dummy)

	assert.Equal(t, dummy, collection.GetAll()[0], "Element must be equal")
	assert.Equal(t, dummy2, collection.GetAll()[1], "Element must be equal")
}

func TestDBCollection_Remove(t *testing.T) {
	dummy := TestDummy{1, "Test Value"}

	collection := NewDBCollection("dummy")
	collection.Add(dummy)

	assert.Equal(t, 1, len(collection.GetAll()), "Collection must contain ")

	collection.Remove(dummy.Id)

	assert.Equal(t, 0, len(collection.GetAll()), "Collection must be empty")
}