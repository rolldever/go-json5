# About

[go-json5][] is a [JSON5][] **parser**/**unmarshaler** written in [Go][].


# Installing

## Using `go get`

	$ go get github.com/rolldever/go-json5


# Example

JSON5 example (see [json5.org][JSON5]):

```js
{
    foo: 'bar',
    while: true,

    this: 'is a \
multi-line string',

    // this is an inline comment
    here: 'is another', // inline comment

    /* this is a block comment
       that continues on another line */

    hex: 0xDEADbeef,
    half: .5,
    delta: +10,
    // to: Infinity,   // and beyond!

    finally: 'a trailing comma',
    oh: [
        "we shouldn't forget",
        'arrays can have',
        'trailing commas too',
    ],
}
```

Parsing using [go-json5][]:

```go
package main

import (
	"fmt"
	"github.com/rolldever/go-json5"
)

func main() {
	var json5Bytes []byte = ...

	// or simply `interface{}` if you know nothing about the JSON5 source
	var obj map[string]interface{}

	if err := json5.Unmarshal(json5Bytes, &obj); err != nil {
		panic(err)
	}

	fmt.Println(obj)
}
```

Unmarshaling to a `struct`:

```go
package main

import (
	"fmt"
	"github.com/rolldever/go-json5"
)

type SomeStruct struct {
	Foo         string   `json:"foo"`
	This        string   `json:"this"`
	Hex         int      `json:"hex"`
	StringArray []string `json:"oh"`
}

func main() {
	var json5Bytes []byte = ...

	var obj SomeStruct
	if err := json5.Unmarshal(json5Bytes, &obj); err != nil {
		panic(err)
	}

	fmt.Println(obj)
}
```


# Full docs

See [https://godoc.org/github.com/rolldever/go-json5][go-json5-doc].

Or run:

	$ godoc github.com/rolldever/go-json5


# License

[go-json5][] is distributed under the [MIT License][].



[Go]: https://golang.org/
[JSON5]: http://json5.org/
[go-json5]: https://github.com/rolldever/go-json5
[go-json5-doc]: https://godoc.org/github.com/rolldever/go-json5
[MIT License]: ./LICENSE
