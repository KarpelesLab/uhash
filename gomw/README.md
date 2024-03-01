# gomw

MultiWriter that can use goroutines for parallelized writes.

Using `gomw.New` returns an io.Writer that implements `io.ReaderFrom` and will
instanciate one goroutine per target when doing operations like `io.Copy()`.

The `Write()` method itself will not however instanciate goroutines as this
would be costly to perform multiple times. This may change in the future.

## Usage

```go
mw := gomw.New(targets...)
_, err := io.Copy(mw, source)
```

