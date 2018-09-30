package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	//	"github.com/docopt/docopt-go"
)

func main() {
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

	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.Use(martini.Static(assets, martini.StaticOptions{SkipLogging: true}))
	m.Run()
}
