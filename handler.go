package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc { // runs once routes set up
	//	TODO: Implement this...
	// Where:
	// r: an interface that decides how responses get written
	// w: struct describing data of the INCOMING request, with fields to be read
	return func(w http.ResponseWriter, r *http.Request) { // parameters come from HandlerFunc type definition, run once per request
		path := r.URL.Path                     // pulls requested path out of the request
		if dest, ok := pathsToUrls[path]; ok { //if ok, dest will have a value
			http.Redirect(w, r, dest, http.StatusFound) // redirect response
			return
		}
		fallback.ServeHTTP(w, r) // runs if path is not in the map
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
