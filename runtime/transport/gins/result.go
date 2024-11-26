/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package gins implements the functions, types, and interfaces for the module.
package gins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"
)

type Result struct {
	Success bool           `json:"success"`
	Total   int64          `json:"total,omitempty"`
	Data    any            `json:"data,omitempty"`
	Extra   any            `json:"extra,omitempty"`
	Error   *httperr.Error `json:"error,omitempty"`
}

// RetJSON Response json data with status code
func RetJSON(c *gin.Context, status int, v any) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	c.Set(ContextResponseBodBytesKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// RetSuccess Response success data with status code
func RetSuccess(c *gin.Context, v any) {
	RetJSON(c, http.StatusOK, Result{
		Success: true,
		Data:    v,
	})
}

// RetOK Response success data with status code
func RetOK(c *gin.Context) {
	RetJSON(c, http.StatusOK, Result{
		Success: true,
	})
}

// RetPage Response page data with status code
func RetPage(c *gin.Context, v any, total int64, args ...map[string]any) {
	if v == nil {
		v = make([]any, 0)
	}
	var extra any
	if len(args) > 0 {
		if v, ok := args[0]["extra"]; ok {
			extra = v
		}
	}
	RetJSON(c, http.StatusOK, Result{
		Success: true,
		Data:    v,
		Extra:   extra,
		Total:   total,
	})
}

// RetError Response error data with status code
func RetError(c *Context, err error, status ...int) {
	var ierr *httperr.Error
	if ok := errors.As(err, &ierr); !ok {
		ierr = httperr.FromError(httperr.InternalServerError(err.Error())) // default error
	}

	code := int(ierr.Code)
	if len(status) > 0 {
		code = status[0]
	}

	if code >= 500 {
		ctx := c.Request.Context()
		//ctx = context.NewTag(ctx, logging.TagKeySystem)
		ctx = context.NewStack(ctx, fmt.Sprintf("%+v", err))
		//logging.Context(ctx).LogAttrs(ctx, slog.LevelError, "Internal server error", slog.Any("error", err))
		ierr.Detail = http.StatusText(http.StatusInternalServerError)
	}

	ierr.Code = int32(code)
	_ = c.Error(gin.Error{
		Err:  ierr,
		Type: gin.ErrorTypeAny,
		Meta: c.Request.URL.RawQuery,
	})
	RetJSON(c, code, Result{Error: ierr})
}
