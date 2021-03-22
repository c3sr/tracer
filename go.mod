module github.com/c3sr/tracer

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

replace github.com/uber/jaeger => github.com/jaegertracing/jaeger v1.22.0

replace github.com/jaegertracing/jaeger => github.com/uber/jaeger v1.22.0

replace github.com/openzipkin-contrib/zipkin-go-opentracing => github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5

require (
	github.com/GeertJohan/go-sourcepath v0.0.0-20150925135350-83e8b8723a9b
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/Workiva/go-datastructures v1.0.52
	github.com/apex/log v1.9.0
	github.com/c3sr/config v1.0.1
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/machine v1.0.0
	github.com/c3sr/nvml-go v1.0.0
	github.com/c3sr/utils v1.0.0
	github.com/c3sr/uuid v1.0.1
	github.com/c3sr/vipertags v1.0.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/imdario/mergo v0.3.12
	github.com/jaegertracing/jaeger v1.22.0 // indirect
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/labstack/echo/v4 v4.2.1
	github.com/nicolai86/instruments v0.0.0-20170630130909-a667d8f6e278
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492
	github.com/opentracing-contrib/perfevents v0.0.0-20171011010702-a7a7e747782c
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil v3.21.2+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger v1.22.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/ulule/deepcopier v0.0.0-20200430083143-45decc6639b6
	github.com/unknwon/com v1.0.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.36.0
)
