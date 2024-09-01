version: {{ .Version }}
app: {{ .AppName }}
type: {{ .ServiceType }} # grpc | http | grpc+http
go: 
  version: {{ .Go.Version }}
  module: {{ .Go.Module }}
style: {{ .Style }}
