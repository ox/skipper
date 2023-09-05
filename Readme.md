# Skipper

An early prototype of a skiplist<sup>[1]</sup> data structure with persistence. It's effectively a key-value store where the keys are strings and the values are bytes. 

## Usage

```go

s := skiplist.New()
s.Set("hello", []byte("world"))

if val, ok := s.Get("hello"); ok {
  fmt.Printf("hello=%s", string(val))
}

```

# Roadmap

- [ ] Improve file format to add version, metadata, and data size prefixes
- [ ] Add a WAL<sup>[2]</sup>
- [ ] Add a text-based API like redis (ex: `SET foo bar`)
- [ ] Add a TCP/HTTP/JSON API server

Inspirations:

- [Skipdb](https://github.com/stevedekorte/skipdb)
- [Building a Log-Structured Merge Tree in Go](https://dev.to/justinethier/log-structured-merge-trees-1jha)

[1]: https://en.wikipedia.org/wiki/Skip_list
[2]: https://en.wikipedia.org/wiki/Write-ahead_logging
