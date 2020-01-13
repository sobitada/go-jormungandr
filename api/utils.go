package api

import "net/url"

// combines the base URL with the given path. if the given
// path is not valid, then an error will be returned.
func combine(base *url.URL, path string) (*url.URL, error) {
    var ref, err = url.Parse(path)
    if err == nil {
        return base.ResolveReference(ref), nil
    }
    return nil, err
}
