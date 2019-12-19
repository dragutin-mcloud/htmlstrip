# HTML strip for Go

[![Used By][used-by-svg]][used-by-link]
[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

Strips HTML tags and return pure text in a efficient way

## Usage

```
package main

import "fmt"
import "github.com/dragutin-mcloud/htmlstrip"

func main() {
	original := []byte("TEXT1<script>console.log('Here i am')</script>TEXT2<script>console.log('Here i am again')</script>TEXT3")
	stripped:=htmlstrip.StripHTML(original)
	fmt.Println(string(stripped))
}
```


### OUTPUT:
`text1text2text3`

 [used-by-svg]: https://sourcegraph.com/github.com/dragutin-mcloud/htmlstrip/-/badge.svg
 [used-by-link]: https://sourcegraph.com/github.com/dragutin-mcloud/htmlstrip?badge
 [goreport-svg]: https://goreportcard.com/badge/github.com/dragutin-mcloud/htmlstrip
 [goreport-link]: https://goreportcard.com/report/github.com/dragutin-mcloud/htmlstrip
 [build-status-svg]: https://api.travis-ci.org/dragutin-mcloud/htmlstrip.svg?branch=master
 [build-status-link]: https://travis-ci.org/dragutin-mcloud/htmlstrip
 [coverage-status-svg]: https://coveralls.io/repos/dragutin-mcloud/htmlstrip/badge.svg?branch=master
 [coverage-status-link]: https://coveralls.io/r/dragutin-mcloud/htmlstrip?branch=master
 [codeclimate-status-svg]: https://codeclimate.com/github/dragutin-mcloud/htmlstrip/badges/gpa.svg
 [codeclimate-status-link]: https://codeclimate.com/github/dragutin-mcloud/htmlstrip
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/dragutin-mcloud/htmlstrip
 [license-svg]: https://img.shields.io/badge/license-BSD--style+patent--grant-blue.svg
 [license-link]: https://github.com/dragutin-mcloud/htmlstrip/blob/master/LICENSE
