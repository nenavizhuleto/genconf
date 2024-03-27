package genconf_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nenavizhuleto/genconf"
)

type Config struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 struct {
		F4 int `json:"f4"`
	} `json:"field3"`
}

const (
	CNAME = "test.json"
	CPATH = "testdata"
)

func TestJSONLoad(t *testing.T) {

	content := []byte("{\"value\":\"hello world\"}")
	if err := os.WriteFile(filepath.Join(CPATH, CNAME), content, 0644); err != nil {
		t.Fatal(err)
	}

	var config struct {
		Value string `json:"value"`
	}

	if err := genconf.NewJSON(CNAME).Dir(CPATH).Load(&config); err != nil {
		t.Fatal(err)
	}

	if config.Value != "hello world" {
		t.Fatal("load failed: config.Value != 'hello world'")
	}

	t.Logf("Load: %+v\n", config)
}

func TestJSONUseCase(t *testing.T) {

	config := Config{
		Field1: "foo",
		Field2: "bar",
	}

	config.Field3.F4 = 42

	configuration := genconf.NewJSON(CNAME)
	if err := configuration.Load(&config); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(genconf.DefaultPath)

	t.Logf("config: %+v\n", config)

}

func TestJSONFileCreation(t *testing.T) {
	configuration := genconf.NewJSON(CNAME)
	defer os.RemoveAll(genconf.DefaultPath)
	if err := configuration.Load(&Config{}); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(genconf.DefaultPath); os.IsNotExist(err) {
		t.Fatal(err)
	}
	if _, err := os.Stat(configuration.FullPath()); os.IsNotExist(err) {
		t.Fatal(err)
	}
}
