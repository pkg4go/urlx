package urlx

import "net/url"

func Resolve(from, to string) (string, error) {
	base, err := url.Parse(from)
	if err != nil {
		return "", err
	}

	ref, err := url.Parse(to)
	if err != nil {
		return "", err
	}

	base = base.ResolveReference(ref)
	return base.String(), nil
}
