# greeting

Go library that builds the greeting `greeter` returns.

`Greet(name)` trims the name and returns `Hello, <name>!`. It returns
`ErrEmptyName` for a blank name and `ErrNameTooLong` for a name over 64
characters.
`Farewell(name)` returns `Goodbye, <name>!` and accepts the same names.

```bash
putnami test --projects greeting
```
