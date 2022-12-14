# argparser

An opinionated argument parser for Go.

## Table Of Contents

- [Usage](#usage)
- [License](#license)

## Usage

This package allows for understanding the following command line arguments format.
`command [arguments] [parameters]` i.e.: `command do --something`

Import the package into your project.

```go
import "github.com/eloonstra/argparser"
```

Grab the arguments from the command line.

```go
args := argparser.Parse()
```
Now you can check whether an argument is present.

```go
if args.HasArg("foo") {
    // do something
}
```

You can also check for parameters (including finding out their value) as follows.

```go
if args.HasParam("foo") {
    value, err := args.GetParamValue("foo")
    // do something
}
```

That's it! All the other stuff is handled for you.

## License

[MIT](LICENSE)
