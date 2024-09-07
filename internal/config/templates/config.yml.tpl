version: {{ .Version }}
app: {{ .AppName }}
type: {{ .ServiceType }} # grpc | http | grpc+http
style: {{ .Style }}
go: 
  version: {{ .Go.Version }}
  module: {{ .Go.Module }}
  tidy: {{ .Go.Tidy }}
