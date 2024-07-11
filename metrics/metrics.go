package metrics

import (
	"net/http"

	"github.com/origadmin/toolkits/context"
)

type Metrics interface {
	ObserverHTTP(ctx context.Context, request *http.Request, response *http.Response, err error)
}
