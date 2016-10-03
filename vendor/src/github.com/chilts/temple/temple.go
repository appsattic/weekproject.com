package temple

import (
	"errors"
	"html/template"
	"os"
	"sync"
)

var (
	ErrNotADirectory = errors.New("Not a directory")
)

// Type Temple allows you to read a directory of templates and either cache them (in production mode) or discard them
// to be re-read (in development mode).
type Temple struct {
	Dir   string
	Base  string
	Cache bool
	// protect the cache map
	mu    sync.RWMutex
	cache map[string]*template.Template
}

// NewTemple returns an initialised Temple. The directory should exist, the baseFilename should exist and cache should
// be true (for each template to be cached) or false (to be discarded and re-read on each invocation).
func NewTemple(dir, baseFilename string, cache bool) (*Temple, error) {
	// firstly, check that this directory exists
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, err
	}

	// check that this is a directory
	if !stat.IsDir() {
		return nil, ErrNotADirectory
	}

	// all okay
	tmpl := Temple{
		Dir:   dir,
		Base:  baseFilename,
		Cache: cache,
		cache: make(map[string]*template.Template),
	}

	return &tmpl, nil
}

// Get will return the html/template you asked for. Note: you should supply the full filename such as "index.html",
// rather than just "index". First time through, the template is read from disk. Second time through depends on whether
// you asked for the templates to be cached or not.
func (t Temple) Get(name string) (*template.Template, error) {
	// take a read lock to start off
	t.mu.RLock()
	tmpl, ok := t.cache[name]
	t.mu.RUnlock()
	if ok {
		return tmpl, nil
	}

	// doesn't yet exist, so let's read the templates in and store them
	tmpl, err := template.ParseFiles(t.Dir+"/"+t.Base, t.Dir+"/"+name)
	if err != nil {
		return nil, err
	}

	// if we don't want to cache, then just forget about the template immediately
	if !t.Cache {
		return tmpl, err
	}

	// now store this in the cache
	t.mu.Lock()
	t.cache[name] = tmpl
	t.mu.Unlock()

	return tmpl, err
}
