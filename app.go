package gomvc

import (
	"fmt"
	"path"
	//"github.com/flosch/pongo"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	//"log"
	"net/http"
	//"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<script>console.log('xd');</script>\n"+strings.Repeat("\n", 256))
	f := res.(http.Flusher)
	f.Flush()
	amt := time.Duration(10000)
	time.Sleep(time.Millisecond * amt)
	fmt.Fprintf(res, "<script>console.log('cu');</script>\n")
}

func handle2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `
        <html>
            <head>
            </head>  
            <body>
                <iframe id="transport"></iframe>
                <script type="text/javascript">
                    document.onreadystatechange = function(){
                        if(document.readyState == "complete"){
                            //document.getElementById("transport").src = "/"
                        }
                    }
                </script>
            </body>
        </html>
    `)
}

type FutureReq func(res http.ResponseWriter, req *http.Request)
type RouteMap map[string]interface{}

type App struct {
	StaticRoot  string
	ViewsRoot   string
	ProjectPath string
	Routing     RouteMap
	Store       *sessions.FilesystemStore
}

func (a App) ReqWrapper(handler interface{}) FutureReq {
	return func(res http.ResponseWriter, req *http.Request) {
		val := reflect.ValueOf(handler)
		attrType := val.Type().In(0)
		typ := reflect.New(attrType)
		el := typ.Elem()
		inst := el.Interface()
		session, sesserr := a.Store.Get(req, "topfriends-session")

		if sesserr != nil {
			fmt.Println(sesserr)
		}

		nw := Wrapper{req, res, a, session}
		val.Call([]reflect.Value{reflect.ValueOf(inst), reflect.ValueOf(nw)})
	}
}

func (a App) Run() {
	fmt.Println(path.Join(a.ProjectPath, "tmp"))
	a.Store = sessions.NewFilesystemStore(path.Join(a.ProjectPath, "tmp"), []byte("something-very-secret"))
	r := mux.NewRouter()

	/*dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}*/

	for k, v := range a.Routing {
		reg := regexp.MustCompile(`[^\s\t]+`)
		routeParams := reg.FindAllString(k, -1)
		method := routeParams[0]
		route := routeParams[1]
		handler := v
		r.HandleFunc(route, a.ReqWrapper(handler)).Methods(method)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(a.StaticRoot)))
	http.Handle("/", r)
	fmt.Println("Started on " + "8080")
	http.ListenAndServe(":8080", nil)
}

func GotServed() App {
	// http.HandleFunc("/", handle)
	// http.HandleFunc("/test", handle2)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }\
	return App{}
}
