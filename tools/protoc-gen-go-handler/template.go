package main

import (
	"bytes"
	"strings"
	"text/template"
)

// TODO 模版最后的c.JSON(200, reply)加一个判断如果是 http.body.
var httpTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

type {{.ServiceType}} interface {
{{- range .MethodSets}}
	{{- if not (eq .HeadComment "")}}
	{{.HeadComment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func {{.ServiceType}}APIRouters(ctr {{.ServiceType}}) []httpserver.Router {
	return []httpserver.Router{
		{{- range .Methods}}
		{
			HTTPMethod:   "{{.Method}}",
			Path:         "{{.Path}}",
			HandlerFuncs: []gin.HandlerFunc{ _{{.Name}}Handler(ctr)},
		},
		{{- end}}
	}
}

{{range .Methods}}
func _{{.Name}}Handler(ctr {{$svrType}}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &{{.Request}}{}
		{{- if .HasBody}}
		if err := c.Bind(in); err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}
		
		{{- if not (eq .Body "")}}
		if err := c.BindQuery(in); err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}
		{{- end}}
		{{- else}}
		if err := c.BindQuery(in); err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}
		{{- end}}
		{{- if .HasVars}}
		if err := c.BindUri(in); err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}
		{{- end}}

		if v, ok := interface{}(in).(interface{Validate() error}); ok {
			if err := v.Validate(); err != nil {
				c.JSON(400, err.Error())
				c.Abort()
				return
			}
		}

		reply, err := ctr.{{.Name}}(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			}else {
				c.JSON(200, coder)
			}
			return
		}

		c.JSON(200, reply)
	}
}
{{end}}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	HeadComment  string
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	// http_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
