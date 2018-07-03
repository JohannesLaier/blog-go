package server

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/config"
	"de/blog/projekt/01/core/controller"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"testing"
	"strings"
	"time"
	"os"
)

func TestServer_AddController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	cfg := config.GetConfig()
	s := NewServer(cfg)
	ctrl := controller.NewController("/admin/login")

	s.AddController(ctrl)
	assert.NotEmpty(t, s.controller, "No controller added")
}

func TestServer_AddFileServer(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../")
	defer os.Chdir(backup_cwd)

	result := "func main()"

	go func() {
		ctrl := controller.NewController("Test")
		ctrl.AddHandler("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(result))
		})

		cfg := config.GetConfig()
		s := NewServer(cfg)
		s.AddController(ctrl)
		s.AddFileServer("/data/", ".")
		s.Run()
	}()

	time.Sleep(1 * time.Second)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8443/data/app.go")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	assert.True(t, strings.Contains(string(body), result))

	resp, err = client.Get("https://127.0.0.1:8443/test")
	if err != nil {
		t.Error(err)
	}

	body, err = ioutil.ReadAll(resp.Body)

	assert.Equal(t, result, string(body))
}
