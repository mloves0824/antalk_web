package test

import (
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/binding"
	//"github.com/martini-contrib/gzip"
	//"github.com/martini-contrib/render"
)

type ApiServer struct {
}

func (s *ApiServer) Set(params martini.Params) (int, string) {
	//k := params["key"]
	//v := params["value"]
	//	if k == "" {
	//		return rpc.ApiResponseError(errors.New("missing key"))
	//	}
	//	if v == "" {
	//		return rpc.ApiResponseError(errors.New("missing value"))
	//	}
	//	return rpc.ApiResponseJson("OK")
	return 200, string("OK")
}
