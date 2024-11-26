/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"fmt"
	"strings"
	"testing"
)

func TestLinkList(t *testing.T) {
	paths := []string{"a/b/c/d/e", "a/c/d/e", "a/d/e/f"}
	link := &Link{Path: "root"}
	for _, path := range paths {
		link.AddSub(NewLink(path))
	}
	if !link.Has("a") {
		t.Fatal("Expected a to be in the link")
	}
	if link.Has("b") {
		t.Fatal("Expected b to be in the link")
	}
	if link.Has("c") {
		t.Fatal("Expected c to be in the link")
	}
	if link.Has("d") {
		t.Fatal("Expected d to be in the link")
	}

	for _, path := range paths {
		rules := strings.Split(path, "/")
		sub := link
		for i, rule := range rules {
			sub = getLinkFromPath(sub.Subs, rule)
			if sub == nil {
				t.Fatal("Expected to find path:", rules[i:])
			}
		}
	}

	list := link.StringList("")
	fmt.Println(list)
	//Output: [root/a/b/c/d/e root/a/c/d/e root/a/d/e/f]
}

func TestLinkAdd(t *testing.T) {
	paths := []string{"a/b/c/d/e", "a/c/d/e", "a/d/e/f"}
	link := &Link{Path: "root"}
	link.AddSub(NewLink("/"))

	list := link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("/*"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("*"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("*/a/b/c/d/e"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}
}
