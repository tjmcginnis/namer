# namer

Namer implements a Heroku-like random name generator, using logic similar to what is described in this [StackOverflow comment](https://stackoverflow.com/questions/7621341/how-can-i-programmatically-generate-heroku-like-subdomain-names).

## Example usage

```go
n := namer.New()
name := n.Name()
```

## Testing

This package exports a Namer interface, not a struct type. As such, you can easily mock it in your tests by implementing a type with a `Name() string` method.

For example:

```go
type mockNamer struct {
    want string
}

func (m *mockNamer) Name() string {
    return m.want
}
```
