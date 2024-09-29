{{$svrType :=.ServiceType}}
{{$svrName :=.ServiceName}}

{{- range.MethodSets}}
	const {{$svrType}}_{{.OriginalName}}_OperationName = "/{{$svrName}}/{{.OriginalName}}"
{{- end}}

type {{.ServiceType}}GINServer interface {
{{- range.MethodSets}}
    {{- if ne .Comment ""}}
        {{.Comment}}
    {{- end}}
    {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{.ServiceType}}GINServer(engine *gin.Engine, srv {{.ServiceType}}GINServer) {
{{- range.Methods}}
	engine.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_GIN_Handler(srv))
{{- end}}
}

{{range.Methods}}
	func _{{$svrType}}_{{.Name}}{{.Num}}_GIN_Handler(srv {{$svrType}}GINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
	var in {{.Request}}
  {{- if.HasBody}}
		if err := ctx.ShouldBind(&in{{.Body}}); err != nil {
		ctx.Error(err)
		return
		}
  {{- end}}
	if err := ctx.BindQuery(&in); err != nil {
	ctx.Error(err)
	return
	}
  {{- if.HasVars}}
		if err := ctx.BindUri(&in); err != nil {
		ctx.Error(err)
		return
		}
  {{- end}}
	http.SetOperation(ctx, {{$svrType}}_{{.OriginalName}}_OperationName)
	newCtx := ginctx.WithContext(ctx)
	reply, err := srv.{{.Name}}(newCtx, &in)
	if err != nil {
	ctx.Error(err)
	return
	}
	ctx.JSON(200, reply{{.ResponseBody}})
	return
	}
	}
{{end}}

type {{.ServiceType}}GINClient interface {
{{- range.MethodSets}}
    {{.Name}}(ctx context.Context, req *{{.Request}}, opts ...http.CallOption) (rsp *{{.Reply}}, err error)
{{- end}}
}

type {{.ServiceType}}GINClientImpl struct{
cc *http.Client
}

func New{{.ServiceType}}GINClient (client *http.Client) {{.ServiceType}}GINClient {
return &{{.ServiceType}}GINClientImpl{client}
}

{{range.MethodSets}}
	func (c *{{$svrType}}GINClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...http.CallOption) (*{{.Reply}}, error) {
	var out {{.Reply}}
	pattern := "{{.ClientPath}}"
	path := binding.EncodeURL(pattern, in, {{not .HasBody}})
	opts = append(opts, http.Operation({{$svrType}}_{{.OriginalName}}_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
  {{if.HasBody -}}
		err := c.cc.Invoke(ctx, "{{.Method}}", path, in{{.Body}}, &out{{.ResponseBody}}, opts...)
  {{else -}}
		err := c.cc.Invoke(ctx, "{{.Method}}", path, nil, &out{{.ResponseBody}}, opts...)
  {{end -}}
	if err != nil {
	return nil, err
	}
	return &out, nil
	}
{{end}}
