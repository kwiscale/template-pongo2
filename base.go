package kwpongo2

import (
	"encoding/json"
	"io"
	"log"
	"path/filepath"

	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"
	"github.com/metal3d/kwiscale"
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

	// encode context
	jsonin, err := json.Marshal(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// decode to map to pongo2.Context
	var pctx pongo2.Context
	json.Unmarshal(jsonin, &pctx)

	return tpl.ExecuteWriter(pctx, w)
}

// SetTemplateDir set the root path of templates.
func (p *PongoTemplateEngine) SetTemplateDir(path string) {
	t, err := filepath.Abs(path)

	if err != nil {
		log.Fatalln("Path not ok", err)
	}
	p.tpldir = t
	log.Println("Pango2: template dir set to ", p.tpldir)
}
