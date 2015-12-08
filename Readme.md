
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### urlx

* Some utils for `url`.

### APIs

* Resolve: `Resolve(from, to string) (string, error)`

```go
result, err := Resolve("http://example.com/a", "b/c")
// result: http://example.com/b/c
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/urlx.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/urlx
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: http://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/urlx
