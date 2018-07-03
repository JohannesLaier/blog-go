package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
)

func TestConfig_Parse_NotFound(t *testing.T) {
	assert.Panics(t, func() {
		GetConfig()
	}, "The code did not panic")
}

func TestConfig_GetHttpsPort(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	port := 8443
	config := GetConfig()

	assert.Equal(t, port, config.GetHttpsPort(), "wrong port")
}

func TestConfig_Parse_Found(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	var config Config
	defer func() {
		if err := recover(); err != nil {
			assert.True(t, false, "Could not found the res folder: %s, ", config.GetResourceFolder())
		} else {
			assert.True(t, true)
		}
	}()
	config = GetConfig()
}

func TestConfig_Parse_DefaulValues(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	https_port := 8443
	resource_folder := "res/"
	session_expire := 15

	config := GetConfig()

	assert.Equal(t, config.GetHttpsPort(), https_port)
	assert.Equal(t, config.GetResourceFolder(), resource_folder)
	assert.Equal(t, config.GetSessionExpire(), session_expire)
}

func TestConfig_GetResourceFolder(t *testing.T) {
	folder := "res/"
	config := new(Config)
	config.resource_folder = folder

	assert.Equal(t, folder, config.GetResourceFolder(), "wrong resource folder path")
}

func TestGetConfig(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	config1 := GetConfig()
	config2 := GetConfig()

	assert.Equal(t, config1.GetHttpsPort(), config2.GetHttpsPort())
	assert.Equal(t, config1.GetSessionExpire(), config2.GetSessionExpire())
	assert.Equal(t, config1.GetResourceFolder(), config2.GetResourceFolder())
}

func TestConfig_GetSessionExpire(t *testing.T) {
	session_expire := 15
	config := new(Config)
	config.session_expire = session_expire

	assert.Equal(t, session_expire, config.GetSessionExpire(), "wrong session expire offset")
}