package filter

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/goexts/ggb/settings"
)

const (
	DefaultDelimiter = ":"
)

type Filter interface {
	Skip(req *http.Request) bool
	Denied(method, path string) bool
	Allowed(method, path string) bool
}

type filter struct {
	delimiter    string
	skipAllow    bool
	allows       [MethodTypeMax][]*Link
	skipDeny     bool
	denies       [MethodTypeMax][]*Link
	treeList     []Link
	skipAnyCheck bool
}

type Setting = settings.Setting[Option]

type Option struct {
	Delimiter    string
	SkipDeny     bool
	SkipAllow    bool
	SkipAnyCheck bool
	Allows       []string
	Denies       []string
}

func WithDelimiter(delimiter string) Setting {
	return func(o *Option) {
		o.Delimiter = delimiter
	}
}

func WithAllows(allows ...string) Setting {
	return func(o *Option) {
		o.Allows = allows
	}
}

func WithDenies(denies ...string) Setting {
	return func(o *Option) {
		o.Denies = denies
	}
}

func SkipAnyCheck() Setting {
	return func(o *Option) {
		o.SkipAnyCheck = true
	}
}

func DisableAllow() Setting {
	return func(o *Option) {
		o.SkipAllow = true
	}
}

func DisableDeny() Setting {
	return func(o *Option) {
		o.SkipDeny = true
	}
}

func contains(roots []*Link, path string) bool {
	var root *Link
	paths := strings.Split(strings.TrimPrefix(path, "/"), "/")
	for _, root = range roots {
		if root.Name == "" {
			return true
		}
		if root.Contains(paths) {
			return true
		}
	}
	return false
}

func containsString(roots []string, path string) bool {
	pathLen := len(path)
	if strings.HasSuffix(path, "/") {
		path += "/"
	}
	for _, rule := range roots {
		if rs := len(rule); pathLen >= rs && path[:rs] == rule {
			slog.Info("containsString", "path", path, "rule", rule, "equal", path[:rs] == rule)
			return true
		}
	}
	return false
}

func (f *filter) Skip(req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	return f.Denied(method, path) || !f.Allowed(method, path)
}

func (f *filter) Allowed(method, path string) bool {
	if f.skipAllow {
		return true
	}
	idx := MethodIndex(method)

	if f.skipAnyCheck && idx == MethodAny {
		return false
	}

	if roots := f.allows[idx]; len(roots) > 0 {
		if contains(roots, path) {
			return true
		}
	}

	if f.skipAnyCheck || idx == MethodAny {
		return false
	}

	if roots := f.allows[MethodAny]; len(roots) > 0 {
		if contains(roots, path) {
			return true
		}
	}
	return false
}

func (f *filter) Denied(method string, path string) bool {
	if f.skipDeny {
		return false
	}
	idx := MethodIndex(method)
	if f.skipAnyCheck && idx == MethodAny {
		return false
	}
	if roots := f.denies[idx]; len(roots) > 0 {
		if contains(roots, path) {
			return true
		}
	}

	if f.skipAnyCheck || idx == MethodAny {
		return false
	}

	if roots := f.denies[MethodAny]; len(roots) > 0 {
		if contains(roots, path) {
			// slog.Info("Deny Check on Contains Method", "method", method, "path", path)
			return true
		}
	}
	return false
}

func (f *filter) buildRoot(tree []*Link, path string) []*Link {
	path = strings.TrimPrefix(path, "/")
	if path == "" {
		return tree
	}
	paths := strings.Split(path, "/")
	return f.buildTree(tree, paths)
}

func (f *filter) buildTree(tree []*Link, paths []string) []*Link {
	if len(paths) == 0 {
		return tree
	}
	var cur *Link
	for i := range tree {
		if tree[i].Name == paths[0] {
			cur = tree[i]
			break
		}
	}
	if cur == nil {
		sz := len(f.treeList)
		f.treeList = append(f.treeList, Link{Name: paths[0]})
		tree = append(tree, &f.treeList[sz])
		cur = &f.treeList[sz]
	}

	if hasNext(paths) {
		paths = paths[1:]
		cur.Next = f.buildTree(cur.Next, paths)
	}
	return tree
}

func hasNext(paths []string) bool {
	return len(paths) > 1
}

func NewFilter(ss ...Setting) Filter {
	o := settings.Apply(&Option{
		Delimiter:    DefaultDelimiter,
		SkipAnyCheck: false,
		SkipDeny:     false,
		SkipAllow:    false,
	}, ss)

	f := &filter{
		delimiter:    o.Delimiter,
		skipAnyCheck: o.SkipAnyCheck,
		treeList:     make([]Link, 0),
		skipAllow:    o.SkipAllow,
		allows:       [MethodTypeMax][]*Link{},
		skipDeny:     o.SkipDeny,
		denies:       [MethodTypeMax][]*Link{},
	}

	for _, s := range o.Allows {
		if s == "" {
			continue
		}
		methods, path := parsePath(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			f.allows[idx] = f.buildRoot(f.allows[idx], path)
		}
		f.skipAllow = false
	}

	for _, s := range o.Denies {
		if s == "" {
			continue
		}
		methods, path := parsePath(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			f.denies[idx] = f.buildRoot(f.denies[idx], path)
		}
		f.skipDeny = false
	}

	return f
}

func showTree(prefix string, list []*Link) {
	if len(list) == 0 {
		return
	}
	for i := range list {
		showTree(prefix+"-"+list[i].Name, list[i].Next)
	}
}

func parsePath(path string, delimiter string) ([]string, string) {
	if path == "" || delimiter == "" {
		return nil, ""
	}
	ss := strings.Split(path, delimiter)
	var (
		methods []string
	)
	switch len(ss) {
	case 1:
		methods = append(methods, "*")
		path = ss[0]
	case 2:
		if ss[0] == "" {
			ss[0] = "*"
		}
		if strings.Contains(ss[0], ",") {
			methods = strings.Split(ss[0], ",")
		} else {
			methods = append(methods, ss[0])
		}
		path = strings.TrimPrefix(ss[1], "/")
	default:
		return nil, ""
	}
	return methods, path
}

type stringFilter struct {
	skipAllow bool
	allows    [MethodTypeMax][]string
	skipDeny  bool
	denies    [MethodTypeMax][]string
}

func (s stringFilter) Skip(req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	return s.Denied(method, path) || !s.Allowed(method, path)
}

func (s stringFilter) Denied(method, path string) bool {
	if s.skipDeny {
		return false
	}

	idx := MethodIndex(method)
	if roots := s.denies[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains *", "method", method, "path", path)
			return true
		}
	}

	if idx == MethodAny {
		return false
	}

	if roots := s.denies[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains Method", "method", method, "path", path)
			return true
		}
	}
	return false
}

func (s stringFilter) Allowed(method, path string) bool {
	if s.skipAllow {
		return true
	}

	idx := MethodIndex(method)
	if roots := s.allows[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains *", "method", method, "path", path)
			return true
		}
	}

	if idx == MethodAny {
		return false
	}

	if roots := s.allows[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			return true
		}
	}
	return false

}

func NewStringFilter(ss ...Setting) Filter {
	opt := settings.Apply(&Option{
		Delimiter:    DefaultDelimiter,
		SkipAnyCheck: false,
		SkipDeny:     false,
		SkipAllow:    false,
	}, ss)

	f := &stringFilter{
		skipAllow: true,
		allows:    [MethodTypeMax][]string{},
		skipDeny:  true,
		denies:    [MethodTypeMax][]string{},
	}

	for _, s := range opt.Allows {
		if s == "" {
			continue
		}
		methods, path := parsePath(s, opt.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			f.allows[idx] = append(f.allows[idx], path)
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
		}
	}

	for _, s := range opt.Denies {
		if s == "" {
			continue
		}
		methods, path := parsePath(s, opt.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			f.denies[idx] = append(f.denies[idx], path)
		}
	}
	return f
}
