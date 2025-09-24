/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"fmt"
	"math/rand"
	"runtime/debug"
	"strconv"
	"testing"
)

// Properly parses allowed paths into Link structures
// Ahoy! Let's check if the allowed paths be parsed into Link structures!
func TestAllowed(t *testing.T) {
	settings := []Setting{
		func(o *Option) {
			o.Allows = []string{"GET:/home", "POST:/login", "PUT:/update", "DELETE:/delete", "GET:/admin/dashboard"}
			o.Denies = []string{"GET:/admin", "POST:/logout", "DELETE:/delete/protected"}
		},
	}
	filter := NewFilter(settings...)
	bf := NewBloomFilter(settings...)

	// Test allowed paths
	if !filter.Allowed("GET", "/home") {
		t.Error("Blimey! The linkFilter should allow GET:/home!")
	}
	if !bf.Allowed("GET", "/home") {
		t.Error("Blimey! The bloomFilter should allow GET:/home!")
	}

	if !filter.Allowed("POST", "/login") {
		t.Error("Blimey! The linkFilter should allow POST:/login!")
	}
	if !bf.Allowed("POST", "/login") {
		t.Error("Blimey! The bloomFilter should allow POST:/login!")
	}

	if !filter.Allowed("GET", "/admin/dashboard") {
		t.Error("Blimey! The linkFilter should allow GET:/admin/dashboard!")
	}
	if !bf.Allowed("GET", "/admin/dashboard") {
		t.Error("Blimey! The bloomFilter should allow GET:/admin/dashboard!")
	}

	// Additional test cases
	if filter.Allowed("GET", "/profile") {
		t.Error("Blimey! The linkFilter should not allow GET:/profile!")
	}
	if bf.Allowed("GET", "/profile") {
		t.Error("Blimey! The bloomFilter should not allow GET:/profile!")
	}

	if filter.Allowed("POST", "/register") {
		t.Error("Blimey! The linkFilter should not allow POST:/register!")
	}
	if bf.Allowed("POST", "/register") {
		t.Error("Blimey! The bloomFilter should not allow POST:/register!")
	}

	if filter.Allowed("PUT", "/update/profile") {
		t.Error("Blimey! The linkFilter should not allow PUT:/update/profile!")
	}
	if bf.Allowed("PUT", "/update/profile") {
		t.Error("Blimey! The bloomFilter should not allow PUT:/update/profile!")
	}

	if filter.Allowed("DELETE", "/delete/user") {
		t.Error("Blimey! The linkFilter should not allow DELETE:/delete/user!")
	}
	if bf.Allowed("DELETE", "/delete/user") {
		t.Error("Blimey! The bloomFilter should not allow DELETE:/delete/user!")
	}

	if filter.Allowed("GET", "/admin/settings") {
		t.Error("Blimey! The linkFilter should not allow GET:/admin/settings!")
	}
	if bf.Allowed("GET", "/admin/settings") {
		t.Error("Blimey! The bloomFilter should not allow GET:/admin/settings!")
	}
}
func TestDenied(t *testing.T) {
	settings := []Setting{
		func(o *Option) {
			o.Allows = []string{"GET:/home", "POST:/login", "PUT:/update", "DELETE:/delete", "GET:/admin/dashboard"}
			o.Denies = []string{"GET:/admin", "POST:/logout", "DELETE:/delete/protected"}
		},
	}
	filter := NewFilter(settings...)
	bf := NewBloomFilter(settings...)

	// Test denied paths
	if !filter.Denied("GET", "/admin") {
		t.Error("Blimey! The linkFilter should deny GET:/admin!")
	}
	if !bf.Denied("GET", "/admin") {
		t.Error("Blimey! The bloomFilter should deny GET:/admin!")
	}

	if !filter.Denied("POST", "/logout") {
		t.Error("Blimey! The linkFilter should deny POST:/logout!")
	}
	if !bf.Denied("POST", "/logout") {
		t.Error("Blimey! The bloomFilter should deny POST:/logout!")
	}

	if !filter.Denied("DELETE", "/delete/protected") {
		t.Error("Blimey! The linkFilter should deny DELETE:/delete/protected!")
	}
	if !bf.Denied("DELETE", "/delete/protected") {
		t.Error("Blimey! The bloomFilter should deny DELETE:/delete/protected!")
	}

	// Test allowed paths should not affect Denied
	if filter.Denied("GET", "/home") {
		t.Error("Blimey! The linkFilter should not deny GET:/home!")
	}
	if bf.Denied("GET", "/home") {
		t.Error("Blimey! The bloomFilter should not deny GET:/home!")
	}

	if filter.Denied("POST", "/login") {
		t.Error("Blimey! The linkFilter should not deny POST:/login!")
	}
	if bf.Denied("POST", "/login") {
		t.Error("Blimey! The bloomFilter should not deny POST:/login!")
	}

	if filter.Denied("GET", "/admin/dashboard") {
		t.Error("Blimey! The linkFilter should not deny GET:/admin/dashboard!")
	}
	if bf.Denied("GET", "/admin/dashboard") {
		t.Error("Blimey! The bloomFilter should not deny GET:/admin/dashboard!")
	}

	// Additional test cases
	if filter.Denied("GET", "/profile") {
		t.Error("Blimey! The linkFilter should not deny GET:/profile!")
	}
	if bf.Denied("GET", "/profile") {
		t.Error("Blimey! The bloomFilter should not deny GET:/profile!")
	}

	if filter.Denied("POST", "/register") {
		t.Error("Blimey! The linkFilter should not deny POST:/register!")
	}
	if bf.Denied("POST", "/register") {
		t.Error("Blimey! The bloomFilter should not deny POST:/register!")
	}

	if filter.Denied("PUT", "/update/profile") {
		t.Error("Blimey! The linkFilter should not deny PUT:/update/profile!")
	}
	if bf.Denied("PUT", "/update/profile") {
		t.Error("Blimey! The bloomFilter should not deny PUT:/update/profile!")
	}

	if filter.Denied("DELETE", "/delete/user") {
		t.Error("Blimey! The linkFilter should not deny DELETE:/delete/user!")
	}
	if bf.Denied("DELETE", "/delete/user") {
		t.Error("Blimey! The bloomFilter should not deny DELETE:/delete/user!")
	}

	if filter.Denied("GET", "/admin/settings") {
		t.Error("Blimey! The linkFilter should not deny GET:/admin/settings!")
	}
	if bf.Denied("GET", "/admin/settings") {
		t.Error("Blimey! The bloomFilter should not deny GET:/admin/settings!")
	}
}
func TestCombinedAllowedAndDenied(t *testing.T) {
	settings := []Setting{
		func(o *Option) {
			o.Allows = []string{"GET:/home", "POST:/login", "PUT:/update", "DELETE:/delete", "GET:/admin/dashboard"}
			o.Denies = []string{"GET:/admin", "POST:/logout", "DELETE:/delete/protected", "GET:/admin/settings"}
		},
	}
	filter := NewFilter(settings...)
	bf := NewBloomFilter(settings...)

	// Test cases where the request should be skipped
	testCases := []struct {
		method string
		path   string
		skip   bool
	}{
		{"GET", "/admin", true},
		{"POST", "/logout", true},
		{"DELETE", "/delete/protected", true},
		{"GET", "/admin/settings", true},
		{"GET", "/profile", true},
		{"POST", "/register", true},
		{"PUT", "/update/profile", true},
		{"DELETE", "/delete/user", true},
	}

	for _, tc := range testCases {
		if (filter.Denied(tc.method, tc.path) || !filter.Allowed(tc.method, tc.path)) != tc.skip {
			t.Errorf("Blimey! The linkFilter should %s skip %s %s!", boolToAction(tc.skip), tc.method, tc.path)
		}
		if (bf.Denied(tc.method, tc.path) || !bf.Allowed(tc.method, tc.path)) != tc.skip {
			t.Errorf("Blimey! The bloomFilter should %s skip %s %s!", boolToAction(tc.skip), tc.method, tc.path)
		}
	}

	// Test cases where the request should not be skipped
	testCases = []struct {
		method string
		path   string
		skip   bool
	}{
		{"GET", "/home", false},
		{"POST", "/login", false},
		{"GET", "/admin/dashboard", false},
	}

	for _, tc := range testCases {
		if (filter.Denied(tc.method, tc.path) || !filter.Allowed(tc.method, tc.path)) != tc.skip {
			t.Errorf("Blimey! The linkFilter should %s skip %s %s!", boolToAction(tc.skip), tc.method, tc.path)
		}
		if (bf.Denied(tc.method, tc.path) || !bf.Allowed(tc.method, tc.path)) != tc.skip {
			t.Errorf("Blimey! The bloomFilter should %s skip %s %s!", boolToAction(tc.skip), tc.method, tc.path)
		}
	}
}

func boolToAction(b bool) string {
	if b {
		return "not"
	}
	return ""
}

// Deals with nil settings input
func TestDealsWithNilSettingsInput(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic: %s", debug.Stack())
			t.Errorf("The code did not panic")
		}
	}()
	// Blimey! What if there be no settings at all?
	filter := NewFilter()
	if filter == nil {
		t.Errorf("Expected a linkFilter instance, got nil")
	}
}

func BenchmarkNewFilter(b *testing.B) {
	var allows []string
	for i := 0; i < 1000; i++ {
		allows = append(allows, "GET:/home/"+strconv.Itoa(i)+"/id")
	}
	settings := []Setting{
		WithAllows(allows...),
	}
	f := NewFilter(settings...)
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		if !f.Allowed("GET", "/home/"+strconv.Itoa(i)+"/id") {
			b.Error("Blimey! The linkFilter should allow GET:/home/" + strconv.Itoa(i) + "/id!")
		}
	}
}

func skipOldFilter(path string, prefixes ...string) bool {
	if len(prefixes) == 0 {
		return false
	}

	pathLen := len(path)
	for _, p := range prefixes {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}
	return false
}

func BenchmarkOldFilter(b *testing.B) {
	var allows []string
	for i := 0; i < 1000; i++ {
		allows = append(allows, "/home/"+strconv.Itoa(i)+"/id")
	}
	// settings := []Setting{
	// 	WithAllows(allows...),
	// }
	// f := NewFilter(settings...)
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		skipOldFilter("/home/"+strconv.Itoa(i)+"/id", allows...)
	}
}

func BenchmarkNewStringFilter(b *testing.B) {
	var allows []string
	for i := 0; i < 1000; i++ {
		allows = append(allows, "GET:/home/"+strconv.Itoa(i)+"/id")
	}
	settings := []Setting{
		WithAllows(allows...),
	}
	f := NewStringFilter(settings...)
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		if !f.Allowed("GET", "/home/"+strconv.Itoa(i)+"/id") {
			b.Error("Blimey! The linkFilter should allow GET:/home/" + strconv.Itoa(i) + "/id!")
		}
	}
}

// BenchmarkFilterVsStringFilter benchmarks the performance of the filter and string filter.
// goos: windows
// goarch: amd64
// pkg: github.com/origadmin/toolkits/net/filter
// cpu: 12th Gen Intel(R) Core(TM) i7-12700H
// BenchmarkFilterVsStringFilter
// BenchmarkFilterVsStringFilter/Filter
// BenchmarkFilterVsStringFilter/Filter-20         	  500352	      2234 ns/op
// BenchmarkFilterVsStringFilter/StringFilter
// BenchmarkFilterVsStringFilter/StringFilter-20   	11157154	       106.8 ns/op
// BenchmarkFilterVsStringFilter/BloomFilter
// BenchmarkFilterVsStringFilter/BloomFilter-20    	 9271764	       134.2 ns/op
func BenchmarkFilterVsStringFilter(b *testing.B) {
	// Create a slice of allows and denies
	allows := make([]string, 1000)
	denies := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		allows[i] = fmt.Sprintf("GET:/home/%d/id", i)
		denies[i] = fmt.Sprintf("/home/%d/id", i+100)
	}

	// Create a filter and a string filter
	f := NewFilter(WithAllows(allows...), WithDenies(denies...))
	sf := NewStringFilter(WithAllows(allows...), WithDenies(denies...))
	bf := NewBloomFilter(WithAllows(allows...), WithDenies(denies...))
	b.ResetTimer()
	// Benchmark the filter
	b.Run("Filter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f.Allowed("GET", fmt.Sprintf("/home/%d/id", rand.Intn(1000)))
		}
	})

	// Benchmark the string filter
	b.Run("StringFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sf.Allowed("GET", fmt.Sprintf("/home/%d/id", rand.Intn(1000)))
		}
	})

	b.Run("BloomFilter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bf.Allowed("GET", fmt.Sprintf("/home/%d/id", rand.Intn(1000)))
		}
	})

}
