# Randomizer

## Getting Started
### Download

```sh
go get -u github.com/walkersumida/randomizer
```

### Usage

```golang
import "github.com/walkersumida/randomizer"

fmt.Println(randomizer.Run("a,b,c,d"))
// => "b"

fmt.Println(randomizer.Run("10,20,30,40,50"))
// => "40"
```