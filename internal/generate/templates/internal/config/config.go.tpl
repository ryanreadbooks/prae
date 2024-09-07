package config

import (
  {{ if .HasHttp }}
  "github.com/zeromicro/go-zero/rest"
  {{ end }}
  {{ if .HasGrpc }}
	"github.com/zeromicro/go-zero/zrpc"
  {{ end }}
)

type Config struct {
  DB struct {
		User   string `json:"user"`
		Pass   string `json:"pass"`
		Addr   string `json:"addr"`
		DbName string `json:"db_name"`
	} `json:"db"`

  {{ if .HasHttp }}
  Http rest.RestConf `json:"http"`
  {{ end }}
  {{ if .HasGrpc }}
  Grpc zrpc.RpcServerConf `json:"grpc"`
  {{ end }}
}
