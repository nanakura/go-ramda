
# Install
```shell
go get github.com/nanakura/go-ramda@latest
```



# UseCase

```go
package main

import (
	"fmt"
	"github.com/nanakura/go-ramda"
)

func main() {
	fmt.Println(ramda.Map(ramda.Multiply(2))([]int{2, 5, 10, 100})) // [4 10 20 200]
}
```



# Document

[go-ramda Wiki](https://github.com/nanakura/go-ramda/wiki)



# License

MIT
