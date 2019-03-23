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

```go
out, _ = CombinedOutput(
	[]string{"rmdir", "not_exist_dir"},
)
fmt.Println(string(out))
// Output:
// rmdir: failed to remove 'not_exist_dir': No such file or directory
```

## Installation

```
$ go get github.com/mattn/go-pipeline
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
