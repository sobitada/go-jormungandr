package wrapper

import "net/url"

func combine(base url.URL, path string) (string, error) {
    var ref, err = url.Parse(path)
    if err == nil {
        return base.ResolveReference(ref).String(), nil
    }
    return "", err
}
