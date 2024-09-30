package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type urlBinding struct {
}

var bind = new(urlBinding)

func (u urlBinding) Name() string {
	return "uri"
}

func (u urlBinding) BindUri(m map[string][]string, obj any) error {
	return binding.MapFormWithTag(obj, m, "json")
}

// BindURI bind form parameters to target.
func BindURI(ctx *gin.Context, target interface{}) error {
	m := make(map[string][]string, len(ctx.Params))
	for _, v := range ctx.Params {
		m[v.Key] = []string{v.Value}
	}
	return bind.BindUri(m, target)
}
