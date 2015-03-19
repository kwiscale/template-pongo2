# Kwiscale addon to use Pongo2 template engine

That addon can be used with [Kwiscale](https://github.com/kwiscale/framework) to use [Pongo2](https://github.com/flosch/pongo2) template engine.

# How to

Install addon with

```
go get github.com/kwiscale/template-pongo2
```

Then in you application:

```go

import (
    "github.com/kwiscale/framework"
    _ "github.com/kwiscale/template-pongo2"
)

```

Note the required "_" that will only run "init".

Then, in configuration:

```go
app := kwiscale.NewApp(&kwiscale.Config{
    TemplateEngine : "pongo2"
})
//...
``` 

Render() method in handlers will use Pongo2.
