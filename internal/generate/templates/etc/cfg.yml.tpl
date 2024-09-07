{{- if .HasHttp }}
http:
  mode: dev
  log:
    level: debug
  name: {{ .AppName }}-api
  host: 0.0.0.0
  port: 8080
{{ end }}

{{ if .HasGrpc }}
grpc:
  mode: dev
  log:
    level: debug
  name: {{ .AppName }}-grpc
  listenon: 0.0.0.0:9090
  etcd:
    hosts:
      - 127.0.0.1:2379
    key: {{ .AppName }}.rpc
{{ end }}

db:
  user: ${ENV_DB_USER}
  pass: ${ENV_DB_PASS}
  addr: ${ENV_DB_ADDR}
  db_name: ${ENV_DB_NAME}
