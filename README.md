# Kwiscale addon to use Pongo2 template engine

That addon can be used with [Kwiscale](https://github.com/kwiscale/framework) (please use gopkg.in url [http://gopkg.in/kwiscale/framework.v0] instead) to use [Pongo2](https://github.com/flosch/pongo2) template engine.

# How to

At this time, we're in alpha stage so please use "v0" that is the "master" version

Install addon with

```
go get gopkg.in/kwiscale/framework.v0
go get gopkg.in/kwiscale/template-pongo2.v0
```

It's important to use the same version for framework and template.

Then in you application:

```go

import (
    "gopkg.in/kwiscale/framework.v0"
    _ "gopkg.in/kwiscale/template-pongo2.v0"
)

```

Note the required "_" that will only run "init".

Then, in configuration:

```go
app := kwiscale.NewApp(&kwiscale.Config{
    TemplateEngine : "pongo2",
    TeplateDir : "./templates",
})
//...
``` 

Render() method in handlers will use Pongo2.
