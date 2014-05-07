package gomvc

import "net/http"
import "fmt"

type Wrapper struct {
	Req *http.Request
	Res http.ResponseWriter
}

func (w Wrapper) Write(str string) {
	fmt.Fprintf(w.Res, str)
}

func (w Wrapper) Render(path string) {

}

func NewWrapper(req *http.Request, res http.ResponseWriter) Wrapper {
	return Wrapper{req, res}
}
