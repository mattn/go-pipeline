# go-pipeline

unix command pipeline on golang

## Usage

```go
out, err := Output(
	[]string{"git", "log", "--oneline"},
	[]string{"grep", "first import"},
	[]string{"wc", "-l"},
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(out))
// Output:
// 1
```

## Installation

```
$ go get github.com/mattn/go-pipeline
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
