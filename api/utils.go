package api

import "net/url"

func combine(base *url.URL, path string) (*url.URL, error) {
	var ref, err = url.Parse(path)
	if err == nil {
		return base.ResolveReference(ref), nil
	}
	return nil, err
}
