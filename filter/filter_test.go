package filter

import (
	"strconv"
	"testing"
)

// Correctly builds filter with default delimiter
func TestNewFilterWithDefaultDelimiter(t *testing.T) {
	// Arrr! Let's see if the filter be built with the default delimiter!
	settings := []Setting{}
	filter := NewFilter(settings...)
	if filter == nil {
		t.Error("Shiver me timbers! The filter be nil!")
	}
}

// Properly parses allowed paths into Link structures
func TestParseAllowedPaths(t *testing.T) {
	// Ahoy! Let's check if the allowed paths be parsed into Link structures!
	settings := []Setting{
		func(o *Option) {
			o.Allows = []string{"GET:/home", "POST:/login"}
		},
	}
	filter := NewFilter(settings...)

	if !filter.Allowed("GET", "/home") {
		t.Error("Blimey! The filter should not allow GET:/home!")
	}

	if !filter.Allowed("POST", "/login") {
		t.Error("Blimey! The filter should not allow POST:/login!")
	}

	if filter.Allowed("*", "/admin") {
		t.Error("Blimey! The filter should not allow *:/admin!")
	}

	if filter.Allowed("GET", "/admin") {
		t.Error("Blimey! The filter should not allow GET:/admin!")
	}
}

// Properly parses denied paths into Link structures
func TestParseDeniedPaths(t *testing.T) {
	// Avast! Let's see if the denied paths be parsed into Link structures!
	settings := []Setting{
		// {denies: []string{"GET:/admin", "DELETE:/user"}},
		WithDenies("GET:/admin", "DELETE:/user"),
	}
	filter := NewFilter(settings...)

	if !filter.Denied("GET", "/admin") {
		t.Error("Blimey! The filter should not deny GET:/admin!")
	}

	if !filter.Denied("DELETE", "/user") {
		t.Error("Blimey! The filter should not deny DELETE:/user!")
	}

	if filter.Denied("POST", "/user") {
		t.Error("Blimey! The filter should deny POST:/user!")
	}
}

// Handles paths with multiple delimiters
func TestHandlePathsMultipleDelimiters(t *testing.T) {
	// Ahoy matey! Let's see if the filter handles paths with multiple delimiters!
	settings := []Setting{
		WithAllows("GET:/home/user/profile", "GET:/home/user/info", "*:/home/user/data"),
	}
	filter := NewFilter(settings...)
	if !filter.Allowed("GET", "/home/user/profile") {
		t.Error("Blimey! The filter should allow GET:/home/user/profile!")
	}
	if filter.Allowed("POST", "/home/user/profile") {
		t.Error("Blimey! The filter should not allow POST:/home/user/profile!")
	}

	if !filter.Allowed("GET", "/home/user/info") {
		t.Error("Blimey! The filter should allow GET:/home/user/info!")
	}

	if !filter.Allowed("GET", "/home/user") {
		t.Error("Blimey! The filter should allow GET:/home/user!")
	}

	if filter.Allowed("GET", "/home/data") {
		t.Error("Blimey! The filter should not allow GET:/home/data!")
	}

	if filter.Allowed("POST", "/home/user/info") {
		t.Error("Blimey! The filter should not allow POST:/home/user/info!")
	}

	if !filter.Allowed("GET", "/home/user/data") {
		t.Error("Blimey! The filter should allow GET:/home/user/data!")
	}
	if !filter.Allowed("POST", "/home/user/data") {
		t.Error("Blimey! The filter should allow POST:/home/user/data!")
	}
}

// Manages empty string inputs in allows and denies
func TestEmptyStringInputsAllowed(t *testing.T) {
	// Arrr! Let's see if the filter manages empty string inputs!
	settings := []Setting{
		WithAllows(""),
	}
	filter := NewFilter(settings...)
	allowed := filter.Allowed("*", "/home")
	denied := filter.Denied("*", "/home")
	if !allowed {
		t.Error("Blimey! The filter should allow *:/home!", allowed)
	}
	if denied {
		t.Error("Blimey! The filter should not deny *:/home!", denied)
	}
}

func TestEmptyStringInputsDenied(t *testing.T) {
	// Arrr! Let's see if the filter manages empty string inputs!
	settings := []Setting{
		WithDenies(""),
	}
	filter := NewFilter(settings...)
	allowed := filter.Allowed("*", "/home")
	denied := filter.Denied("*", "/home")
	if !allowed {
		t.Error("Blimey! The filter should allow *:/home!", !allowed)
	}
	if denied {
		t.Error("Blimey! The filter should not deny *:/home!", denied)
	}
}

// Handles nil settings input
func TestNilSettingsInput(t *testing.T) {
	// Yo-ho-ho! Let's see if the filter handles nil settings input!
	var settings []Setting = nil
	filter := NewFilter(settings...)
	if filter == nil {
		t.Error("Shiver me timbers! The filter should not be nil with nil settings input!")
	}
}

// Correctly parses allows and denies with default delimiter
func TestParsesAllowsAndDeniesWithDefaultDelimiter(t *testing.T) {
	// Arrr! Let's set sail with some default delimiter!
	settings := []Setting{
		WithAllows("GET:/home", "POST:/login"),
		WithDenies("DELETE:/admin"),
	}
	filter := NewFilter(settings...)
	if !filter.Allowed("GET", "/home") || !filter.Allowed("POST", "/login") {
		t.Error("Blimey! The filter should allow GET:/home and POST:/login!")
	}
	if filter.Allowed("DELETE", "/admin") {
		t.Error("Blimey! The filter should not allow DELETE:/admin!")
	}
}

// Deals with nil settings input
func TestDealsWithNilSettingsInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Blimey! What if there be no settings at all?
	filter := NewFilter(nil)
	if filter == nil {
		t.Errorf("Expected a filter instance, got nil")
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
			b.Error("Blimey! The filter should allow GET:/home/" + strconv.Itoa(i) + "/id!")
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
			b.Error("Blimey! The filter should allow GET:/home/" + strconv.Itoa(i) + "/id!")
		}
	}
}
