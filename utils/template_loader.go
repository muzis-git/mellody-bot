package utils

import (
	"io/ioutil"
	"path"
	"strings"
)

const (
	TEMPLATE_DIRECTORY = "templates"
	WELCOME_TEMPLATE_NAME = "welcome.md"
)

func LoadTemplate(name string, variables map[string]string) (string, error) {
	bytes, err := ioutil.ReadFile(path.Join(TEMPLATE_DIRECTORY, name))
	if err != nil {
		return "", err
	}

	content := string(bytes)
	for key, val := range variables {
		content = strings.Replace(content, key, val, -1)
	}

	return content, nil
}
