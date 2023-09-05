# Skipper

An early prototype of a skiplist[1] data structure with persistence. Next steps here would be:

  [] Add a WAL[2]
  [] Add a text-based API like redis (ex: `SET foo bar`)
  [] Add a TCP/HTTP/JSON API server

Inspirations:

- [Skipdb](https://github.com/stevedekorte/skipdb)
- [Building a Log-Structured Merge Tree in Go](https://dev.to/justinethier/log-structured-merge-trees-1jha)

[1]: https://en.wikipedia.org/wiki/Skip_list
[2]: https://en.wikipedia.org/wiki/Write-ahead_logging
