package main

import (
	"path/filepath"
//	"net/http"


	"github.com/go-martini/martini"
//	"github.com/martini-contrib/render"
//	"github.com/mloves0824/antalk_web/pkg/test"
	//	"github.com/docopt/docopt-go"
)
import "net/http"
import "github.com/gin-gonic/gin"

func main() {
        router := gin.Default()
        router.LoadHTMLGlob("/home/messi/dev/go/src/github.com/mloves0824/antalk_web/cmd/fe/templates/*")
        //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
        router.GET("/index", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.tmpl", gin.H{
                        "title": "Main website",
                })
        })
        // Query string parameters are parsed using the existing underlying request object.
        // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
        router.GET("/test/pegasus/set", func(c *gin.Context) {
                firstname := c.DefaultQuery("key", "Guest")
                lastname := c.Query("value") // shortcut for c.Request.URL.Query().Get("lastname")
        
                c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
        })
        router.Run(":3000")
        
}  


func main_martini() {
	//	const usage = `
	//Usage:
	//	antalk-web-fe [--ncpu=N] [--log=FILE] [--log-level=LEVEL] [--assets-dir=PATH] [--pidfile=FILE]  --listen=ADDR
	//	antalk-web-fe  --version
	//
	//Options:
	//	--ncpu=N                        set runtime.GOMAXPROCS to N, default is runtime.NumCPU().
	//	-l FILE, --log=FILE             set path/name of daliy rotated log file.
	//	--log-level=LEVEL               set the log-level, should be INFO,WARN,DEBUG or ERROR, default is INFO.
	//	--listen=ADDR                   set the listen address.
	//`
	//	d, err := docopt.Parse(usage, nil, true, "", false)
	//	if err != nil {
	//		//log.PanicError(err, "parse arguments failed")
	//		return
	//	}
	//	m := martini.Classic()
	//	m.Get("/", func() string {
	//		return "Hello world!"
	//	})
	//	m.Run()

	var assets string
	abspath, err := filepath.Abs("/home/messi/dev/go/src/github.com/mloves0824/antalk_web/cmd/fe/assets")
	if err != nil {
		//log.PanicErrorf(err, "get absolute path of %s failed", s)
		return
	}
	assets = abspath

	println(assets);

//	api := &test.ApiServer{}
//
//	m := martini.New()
//	m.Use(martini.Recovery())
//	m.Use(render.Renderer())
//	m.Use(martini.Static(assets, martini.StaticOptions{SkipLogging: true}))
//	m.Use(func(c martini.Context, w http.ResponseWriter) {
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	})
//	//r := martini.NewRouter()
//	//m.Group("/test/pegasus", func(r martini.Router) {
////		r.Get("/set/:key", api.Set)
//	//})
////	m.Get("/list", api.List)
//
//	//m.MapTo(r, (*martini.Routes)(nil))
//	//m.Action(r.Handle)
//
//	//m.Get("/hello/:name", func(params martini.Params) string {
//  	//	return "Hello " + params["name"]
//	//})
//
//	m.Get("/hello", func() string {
//  		return "hello world" // HTTP 200 : "hello world"
//	})
//
//	m.Run()
	

	m := martini.Classic()
      //m.Use(martini.Recovery())
      //m.Use(render.Renderer())
      //m.Use(martini.Static(assets, martini.StaticOptions{SkipLogging: true}))
      //m.Use(func(c martini.Context, w http.ResponseWriter) {
      //        w.Header().Set("Content-Type", "application/json; charset=utf-8")
      //})

  	// ... middleware and routing goes here
	//m.Get("/hello", func() string {
	//	return "hello world"
	//})	
	m.Get("/hello1/:name", func(params martini.Params) string {
 		return "Hello " + "name"
	})
	 m.Run()	
}
