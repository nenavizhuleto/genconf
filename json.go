package genconf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var DefaultPath = "config"

type JSON struct {
	create   bool
	filename string
	path     string
}

func NewJSON(filename string) *JSON {
	return &JSON{
		create:   true,
		filename: filename,
		path:     DefaultPath,
	}
}

func (c *JSON) Dir(path string) *JSON {
	c.path = path
	return c
}

func (c *JSON) FullPath() string {
	return filepath.Join(c.path, c.filename)
}

func (c *JSON) Save(config any) error {
	fp := c.FullPath()
	os.Mkdir(c.path, 0777)
	content, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(fp, content, 0644)
}

func (c *JSON) Load(config any) error {
	fp := c.FullPath()
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if !c.create {
			return err
		}
		return c.Save(config)
	} else if err != nil {
		return err
	}

	content, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, config); err != nil {
		return err
	}

	return nil
}
