package gomvc

import (
	"fmt"
	"github.com/gorilla/mux"
	//"log"
	"net/http"
	//"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<script>console.log('xd');</script>\n"+strings.Repeat("\n", 1024))
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

func ReqWrapper(handler interface{}) FutureReq {
	return func(res http.ResponseWriter, req *http.Request) {
		val := reflect.ValueOf(handler)
		typ := reflect.New(val.Type().In(0)).Elem().Interface()
		nw := NewWrapper(req, res)
		val.Call([]reflect.Value{reflect.ValueOf(typ), reflect.ValueOf(nw)})
	}
}

type App struct {
	Routing RouteMap
}

func (a App) Run() {
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
		r.HandleFunc(route, ReqWrapper(handler)).Methods(method)
	}

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
