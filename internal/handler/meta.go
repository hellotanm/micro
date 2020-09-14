package handler

import (
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"

	"github.com/micro/go-micro/v3/api/handler"
	"github.com/micro/go-micro/v3/api/handler/event"
	"github.com/micro/go-micro/v3/api/router"
	"github.com/micro/go-micro/v3/client"
	"github.com/micro/micro/v3/service"
	// TODO: only import handler package
	aapi "github.com/micro/go-micro/v3/api/handler/api"
	ahttp "github.com/micro/go-micro/v3/api/handler/http"
	arpc "github.com/micro/go-micro/v3/api/handler/rpc"
	aweb "github.com/micro/go-micro/v3/api/handler/web"
)

type metaHandler struct {
	c  client.Client
	r  router.Router
	ns string
}

func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service, err := m.r.Route(r)
	if err != nil {
		//er := errors.InternalServerError(m.ns, err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)

		v := struct {
			Code   int         `json:"code"`
			ErrMsg string      `json:"err_msg"`
			Data   interface{} `json:"data"`
		}{
			Code:   404,
			ErrMsg: err.Error(),
			Data:   "",
		}

		j := jsoniter.ConfigCompatibleWithStandardLibrary
		out, e := j.Marshal(&v)

		log.Fatalf(e.Error())

		w.Write(out)

		return
	}

	// TODO: don't do this ffs
	switch service.Endpoint.Handler {
	// web socket handler
	case aweb.Handler:
		aweb.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// proxy handler
	case ahttp.Handler:
		ahttp.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// rpcx handler
	case arpc.Handler:
		arpc.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// event handler
	case event.Handler:
		ev := event.NewHandler(
			handler.WithNamespace(m.ns),
			handler.WithClient(m.c),
		)
		ev.ServeHTTP(w, r)
	// api handler
	case aapi.Handler:
		aapi.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// default handler: rpc
	default:
		arpc.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	}
}

// Meta is a http.Handler that routes based on endpoint metadata
func Meta(s *service.Service, r router.Router, ns string) http.Handler {
	return &metaHandler{
		c:  s.Client(),
		r:  r,
		ns: ns,
	}
}
