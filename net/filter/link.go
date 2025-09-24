/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"cmp"
	"strings"
)

type Linker interface {
	Contains(paths []string) bool
	Equal(path string) bool
	Has(name string) bool
	String() string
}

type Link struct {
	Path string
	Subs []*Link
}

var terminateLink = &Link{Path: Terminator}

func (l *Link) Equal(path string) bool {
	return l.Path == path
}

func (l *Link) Has(name string) bool {
	var t *Link
	for _, t = range l.Subs {
		if t.Equal(name) {
			return true
		}
	}
	return false
}

func (l *Link) HasSubs() bool {
	return len(l.Subs) > 0
}

func (l *Link) Contains(paths []string) bool {
	if l.IsEnd() {
		return true
	}

	if len(paths) == 0 || !l.Equal(paths[0]) {
		return false
	}

	next := paths[1:]
	if len(next) == 0 {
		return true
	}
	var sub *Link
	for _, sub = range l.Subs {
		if sub.Contains(next) {
			return true
		}
	}
	return false
}

func (l *Link) AddSub(s *Link) {
	var t *Link // target
	if l.IsEnd() || s == nil {
		return
	}
	for _, t = range l.Subs {
		if !t.Equal(s.Path) {
			break
		}
		t.AddSubs(s.Subs...)
		return
	}
	l.Subs = append(l.Subs, s)
}

func (l *Link) AddSubs(src ...*Link) {
	var s *Link
	if l.IsEnd() || len(src) == 0 {
		return
	}
	for _, s = range src {
		l.AddSub(s)
	}
}

func (l *Link) IsEnd() bool {
	return !l.HasSubs() && l.Equal(Terminator)
}

func (l *Link) String() string {
	return l.Path
}

func (l *Link) StringList(parent string) []string {
	var ret []string
	path := l.Path
	if parent != "" {
		path = parent + "/" + l.Path
	}
	for i := range l.Subs {
		result := l.Subs[i].StringList(path)
		if result != nil {
			ret = append(ret, result...)
		}
	}
	if len(ret) == 0 {
		return []string{path}
	}
	return ret
}

func newLink(paths []string) *Link {
	links := make([]*Link, 0, len(paths))
	for i := range paths {
		if i != 0 && paths[i] == "" {
			return nil
		}
		links = append(links, &Link{Path: paths[i]})
	}

	if len(links) == 0 {
		return nil
	}
	iter := links[0]
	for i := 1; i < len(links); i++ {
		iter.AddSubs(links[i])
		iter = links[i]
	}
	return links[0]
}

func NewLink(path string) *Link {
	if path == "" {
		return nil
	}
	paths := PathTerminator(splitPath(path))
	return newLink(paths)
}

func PathTerminator(paths []string) []string {
	if len(paths) == 0 {
		return []string{Terminator}
	}

	end := len(paths) - 1
	paths[end] = cmp.Or(paths[end], Terminator)

	// find the end if there is a Terminator
	if IsTerminated(paths) {
		return paths
	}

	return paths
}

func LinkTerminator(l *Link) {
	if l == nil {
		return
	}
	if l.HasSubs() {
		return
	}

	if l.IsEnd() {
		return
	}
	if l.Path == "" {
		l.Path = Terminator
		return
	}

	l.Subs = append(l.Subs, terminateLink)
}

func IsTerminated(paths []string) bool {
	return len(paths) > 0 && paths[len(paths)-1] == Terminator
}

func OnlyTerminator(paths []string) bool {
	switch len(paths) {
	case 1:
		return paths[0] == Terminator
	}
	return false
}

func splitPath(path string) []string {
	if path == "" {
		return nil
	}
	if stringContainsCTLByte(path) {
		return nil
	}
	return strings.Split(strings.TrimPrefix(path, "/"), "/")
}

func splitURI(path string, delimiter string) ([]string, string) {
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
		path = ss[1]
	default:
		return nil, ""
	}
	return methods, path
}

func buildLinkTree(tree []*Link, paths []string) []*Link {
	if len(paths) == 0 {
		return tree
	}

	cur := getLinkFromPath(tree, paths[0])
	if cur == nil {
		cur = &Link{Path: paths[0]}
		tree = append(tree, cur)
	}

	var ok bool
	if paths, ok = hasNext(paths); ok {
		cur.Subs = buildLinkTree(cur.Subs, paths)
	}
	return tree
}

func buildLinkRoot(tree []*Link, path string) []*Link {
	if path == "" {
		return tree
	}
	if path == "/" {
		return []*Link{terminateLink}
	}

	if tree == nil {
		return append(tree, NewLink(path))
	}

	return buildLinkTree(tree, splitPath(path))
}

func getLinkFromPath(tree []*Link, path string) *Link {
	for i := range tree {
		if tree[i].Equal(path) {
			return tree[i]
		}
	}
	return nil
}

// stringContainsCTLByte reports whether s contains any ASCII control character.
func stringContainsCTLByte(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < ' ' || b == 0x7f {
			return true
		}
	}
	return false
}
