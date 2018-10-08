package main

import (
	"path/filepath"
	//	"net/http"
	"log"
	"time"

	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/render"
	//	"github.com/mloves0824/antalk_web/pkg/test"
	//	"github.com/docopt/docopt-go"
	pb "github.com/mloves0824/antalk_web/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
import "net/http"
import "github.com/gin-gonic/gin"

const (
	address     = "localhost:60000"
	defaultName = "world"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/chenb/workspace/github/mnt/gopath/src/github.com/mloves0824/antalk_web/cmd/fe/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/test/pegasus/set", func(c *gin.Context) {
		key := c.DefaultQuery("key", "Guest")
		value := c.Query("value") // shortcut for c.Request.URL.Query().Get("lastname")

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Printf("did not connect: %v", err)
			c.String(http.StatusInternalServerError, "error: set %s %s", key, value)
			return
		}
		defer conn.Close()
		client := pb.NewTestPegasusClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.TestSet(ctx, &pb.SetReq{Key: key, Value: value})
		if err != nil {
			log.Printf("could not set: %v", err)
			c.String(http.StatusInternalServerError, "error %s %s", key, value)
			return
		}
		log.Printf("Status: %s", r.Status)
		c.String(http.StatusOK, "Hello %s %s", key, value)
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

	println(assets)

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
