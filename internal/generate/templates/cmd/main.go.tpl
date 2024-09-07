package main

import (
  "flag"

  "{{ .Go.Module }}/internal/config"
  "{{ .Go.Module }}/internal/svc"
  {{- if .HasHttp }}
  "{{ .Go.Module }}/internal/http"
  {{ end -}}

  {{- if .HasGrpc -}}
  {{ end }}

  {{ if .HasHttp -}}
  "github.com/zeromicro/go-zero/rest"
  {{ end -}}
  {{ if .HasGrpc -}}
  "github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
  {{ end -}}
  "github.com/zeromicro/go-zero/core/conf"
  "github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/{{ .AppName }}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

  var ctx = svc.New(&c)
  group := service.NewServiceGroup()
	defer group.Stop()

  {{ if .HasHttp }}
  restServer := rest.MustNewServer(c.Http)
  http.Register(restServer, ctx)
  group.Add(restServer)
  {{ end }}

  {{ if .HasGrpc }}
  grpcServer := zrpc.MustNewServer(c.Grpc, func(s *grpc.Server) {
		
	})
  group.Add(grpcServer)
  {{ end }}

  group.Start()
}
