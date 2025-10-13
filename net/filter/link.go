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

	segments := splitPath(path)
	if segments == nil {
		return nil
	}

	pathsWithTerminator := PathTerminator(segments)
	return newLink(pathsWithTerminator)
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

	// Handle leading slash: "a/b" and "/a/b" should result in ["a", "b"]
	// Handle trailing slash: "a/" should result in ["a", ""]
	// Handle consecutive slashes: "a//b" should be invalid
	// Handle root: "/" should result in [""]

	trimmedPath := path
	if strings.HasPrefix(path, "/") {
		trimmedPath = path[1:]
	}

	if trimmedPath == "" {
		if path == "/" {
			return []string{""}
		}
		return nil // Empty path or just leading slashes like "//"
	}

	segments := strings.Split(trimmedPath, "/")

	// Check for invalid empty segments (consecutive slashes)
	for i, segment := range segments {
		if segment == "" {
			// If it's not the last segment, it's an intermediate empty segment (e.g., "a//b")
			// If it's the last segment, it's a trailing slash (e.g., "a/") which is allowed for NewLink
			if i < len(segments)-1 {
				return nil // Invalid path due to consecutive slashes
			}
		}
	}

	return segments
}

func splitURI(path string, delimiter string) ([]string, string) {
	if path == "" || delimiter == "" {
		return nil, ""
	}

	var methods []string
	var uriPath string

	if delimiter == ":" {
		ss := strings.SplitN(path, delimiter, 2)
		if len(ss) == 1 {
			methods = append(methods, "*")
			uriPath = ss[0]
		} else { // len(ss) == 2
			if ss[0] == "" {
				methods = append(methods, "*")
			} else if strings.Contains(ss[0], ",") {
				methods = strings.Split(ss[0], ",")
			} else {
				methods = append(methods, ss[0])
			}
			uriPath = ss[1]
		}
		// Final validation of uriPath using NewLink for ':' delimiter
		if NewLink(uriPath) == nil && uriPath != "/" {
			return nil, "" // Invalid URI path
		}
	} else if delimiter == "/" {
		// For '/' delimiter, we need to be careful.
		// If the path starts or ends with '/', or contains '//', it's generally invalid for this context.
		// The only exception is the root path "/", which should be valid.

		if path == "/" {
			return []string{"*"}, "/"
		}

		// Check for leading/trailing slashes or consecutive slashes
		if strings.HasPrefix(path, "/") || strings.HasSuffix(path, "/") || strings.Contains(path, "//") {
			return nil, "" // Invalid URI format for '/' delimiter
		}

		// Now, try to parse methods if present (e.g., "get,post/resource")
		parts := strings.SplitN(path, "/", 2)

		if len(parts) == 1 {
			// No '/' in the path, so the whole path is the URI, and methods are "*"
			methods = []string{"*"}
			uriPath = parts[0]
		} else { // len(parts) == 2
			// Check if the first part contains commas (indicating explicit methods)
			if strings.Contains(parts[0], ",") {
				methods = strings.Split(parts[0], ",")
				uriPath = parts[1]
			} else {
				// No commas, so the first part is also part of the URI, methods are "*"
				methods = []string{"*"}
				uriPath = path // The entire original path is the URI
			}
		}
	} else {
		// Default case if delimiter is neither ":" nor "/"
		return nil, ""
	}

	return methods, uriPath
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
