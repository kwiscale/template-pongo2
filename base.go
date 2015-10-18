package kwpongo2

import (
	"encoding/json"
	"io"
	"log"
	"path/filepath"

	"gopkg.in/kwiscale/framework.v1"

	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"
)

// Init is used to register Pongo2 template engine to Kwiscale as name "pongo2".
func init() {
	// there call a register function
	kwiscale.RegisterTemplateEngine("pongo2", &PongoTemplateEngine{})
}

type PongoTemplateEngine struct {
	tpldir string
}

// Render writes given template path with context to the Writer w.
func (p *PongoTemplateEngine) Render(w io.Writer, path string, ctx interface{}) error {
	path = filepath.Join(p.tpldir, path)
	tpl := pongo2.Must(pongo2.FromFile(path))
	pctx := pongo2.Context{}

	// We try to keep the correct context structure
	switch ctx := ctx.(type) {
	case pongo2.Context:
		pctx = ctx
	case map[string]interface{}:
		for k, v := range ctx {
			pctx[k] = v
		}
	default:
		// but if context is not regular...
		// encode context...
		// TODO: fix the data type lose
		jsonin, err := json.Marshal(ctx)
		if err != nil {
			log.Fatal(err)
		}

		// decode to map to pongo2.Context
		json.Unmarshal(jsonin, &pctx)
	}

	return tpl.ExecuteWriter(pctx, w)
}

// SetTemplateDir set the root path of templates.
func (p *PongoTemplateEngine) SetTemplateDir(path string) {
	t, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	p.tpldir = t
}

// ne need to set options
func (p *PongoTemplateEngine) SetTemplateOptions(ops kwiscale.TplOptions) {}
