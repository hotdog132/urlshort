package urlshort

import (
	"fmt"
	"testing"
)

var mockYaml1 = `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
`
var mockYaml2 = `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: 
`

func TestParseYaml(t *testing.T) {
	yamlInfo, err := parsYaml([]byte(mockYaml1))
	if err != nil {
		t.Errorf("get error when exec parsYaml: %v", err)
	}
	if len(yamlInfo) == 0 {
		t.Errorf("Expected yamlInfo length != 0 but get 0")
	}
}

func TestBuildMap(t *testing.T) {
	yamlInfo, _ := parsYaml([]byte(mockYaml2))
	_, err := buildMap(yamlInfo)
	fmt.Println(err)
	if err == nil {
		t.Errorf("Expected get value error, but nothing happend")
	}
}
