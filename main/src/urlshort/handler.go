package urlshort

import (
	"errors"
	"html"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		requestPath := html.EscapeString(r.URL.Path)
		if url, ok := pathsToUrls[requestPath]; ok {
			http.Redirect(w, r, url, 301)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	yamlInfo, err := parsYaml(yml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	mapPathsToUrls, err := buildMap(yamlInfo)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return MapHandler(mapPathsToUrls, fallback), nil
}

func buildMap(yamlInfo []map[string]string) (map[string]string, error) {
	mapPathsToUrls := make(map[string]string)
	for _, v := range yamlInfo {
		if v["url"] == "" {
			return nil, errors.New("item value nil")
		}
		mapPathsToUrls[v["path"]] = v["url"]
	}
	return mapPathsToUrls, nil
}

func parsYaml(yml []byte) ([]map[string]string, error) {
	m := make([]map[string]string, 0)
	// t := T{}
	err := yaml.Unmarshal(yml, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return m, nil
}
