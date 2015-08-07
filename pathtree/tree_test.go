package pathtree

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestSingleFile(t *testing.T) {
	str := "test.ext"
	s := strings.Split(str, "/")

	items := make(Items)
	items.Add(s)

	if _, ok := items["test.ext"]; !ok {
		t.Fatal("missing element")
	}
}

func TestFileInDirectory(t *testing.T) {
	str := "dir1/test.ext"
	s := strings.Split(str, "/")

	items := make(Items)
	items.Add(s)

	if _, ok := items["dir1"]["test.ext"]; !ok {
		t.Fatal("missing element")
	}
}

func TestFileInManyDirectory(t *testing.T) {
	str := "dir1/dir2/test.ext"
	s := strings.Split(str, "/")

	items := make(Items)
	items.Add(s)

	if _, ok := items["dir1"]["dir2"]["test.ext"]; !ok {
		t.Fatal("missing element")
	}
}

func TestManyFiles(t *testing.T) {
	items := make(Items)

	str1 := "dir1/dir2/test.ext"
	items.Add(strings.Split(str1, "/"))

	str2 := "dir1/other.ext"
	items.Add(strings.Split(str2, "/"))

	if _, ok := items["dir1"]["dir2"]["test.ext"]; !ok {
		t.Fatal("missing element")
	}

	if _, ok := items["dir1"]["other.ext"]; !ok {
		t.Fatal("missing element")
	}

	b, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		t.FailNow()
	}

	log.Println(string(b))
}
