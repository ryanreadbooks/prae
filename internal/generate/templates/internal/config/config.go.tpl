package config

import (
  "github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
  DB struct {
		User   string `json:"user"`
		Pass   string `json:"pass"`
		Addr   string `json:"addr"`
		DbName string `json:"db_name"`
	} `json:"db"`

  {{ if .ServiceTypeHasHttp }}
  Http rest.RestConf `json:"http"`
  {{ end }}
  {{ if .ServiceTypeHasGrpc }}
  Grpc zrpc.RpcServerConf `json:"grpc"`
  {{ end }}
}
