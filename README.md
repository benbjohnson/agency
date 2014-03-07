# Agency

Agency is a fast user-agent parser in pure Go. It uses a simple tokenizer and weighted tokens to determine the correct device type, category and browser.

## Usage

To use Agency, simply import the package and call `Scan()`:

```go
import "github.com/benbjohnson/agency"

ua := agency.Scan("Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 blah blah...")
```


## User Agent Info

The `Scan()` function returns a `UserAgent` with the following string properties:

```
Type
----------
Desktop
Tablet
Mobile

Category
----------
Windows
Mac
Linux
iOS
Blackberry
Android

Browser
----------
Chrome
Safari
Internet Explorer
```


## Performance

On my laptop, I can execute almost 120,000 scans per second.
Your mileage may vary.

```sh
$ go test -test.bench=.
BenchmarkScan	  200000	      7931 ns/op
```
