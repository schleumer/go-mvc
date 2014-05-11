package gomvc

import (
	"fmt"
	"github.com/flosch/pongo"
	"github.com/gorilla/sessions"
	"net/http"
	"path"
	"strings"
)

type Wrapper struct {
	Req     *http.Request
	Res     http.ResponseWriter
	App     App
	Session *sessions.Session
}

var (
	dirtyBits int = 256
)

func (w Wrapper) Write(str string) {
	fmt.Fprintf(w.Res, str)
}

var tplCache = map[string]*pongo.Template{}

func (w Wrapper) Render(p string) {
	tplFile := path.Join(w.App.ViewsRoot, p)
	var tpl *pongo.Template
	if val, ok := tplCache[tplFile]; ok {
		tpl = val
	} else {
		tpl = pongo.Must(pongo.FromFile(tplFile, nil))
	}
	out, err := tpl.Execute(&pongo.Context{})
	if err != nil {
		panic(err)
	}
	w.Write(*out)
}

func (w Wrapper) Push(c string) {
	fmt.Fprintf(w.Res, "<script>console.log('xd');</script>\n"+strings.Repeat("\n", dirtyBits))
	f := w.Res.(http.Flusher)
	f.Flush()
}

func (w Wrapper) End() {
	fmt.Fprintf(w.Res, strings.Repeat("\n", dirtyBits))
	f := w.Res.(http.Flusher)
	f.Flush()
}

func (w Wrapper) SaveSession() {
	w.Session.Save(w.Req, w.Res)
}

/*func NewWrapper(req *http.Request, res http.ResponseWriter) Wrapper {
	return Wrapper{req, res}
}*/
