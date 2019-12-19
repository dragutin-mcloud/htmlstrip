# HTML strip for Go

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
