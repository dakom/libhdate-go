# libhdate-go
A pure go implementation of libhdate

## Testing

`go test ./tests`

## Command-line example

`go build -o ./example/run-example ./example && ./example/run-example`

## TODO
1. check for null dates in HdateSetGdate() in julien.go (use local time from example.go)
2. locale in strings.go hdate_is_hebrew_locale, right now doesn't do anything really
3. overall - move functions to receiver methods instead of passing hDate everywhere!
4. make "diaspora" a field in hebDate struct
5. add time formatting from example.go to strings.go
