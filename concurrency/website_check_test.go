package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furtherwe.geds" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"waat://furtherwe.geds",
	}

	want := map[string]bool{
		"http://www.google.com":      true,
		"http://blog.gypsydave5.com": true,
		"waat://furtherwe.geds":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

}
