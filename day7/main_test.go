package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

var testFile = "test.txt"

func setup(fileContent []byte) {
	if err := os.WriteFile(testFile, fileContent, 777); err != nil {
		log.Fatal(err)
	}
}

func TestProcessInput(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir brhvclj\ndir clnvqg\ndir dtqtvvrn\ndir lcz\ndir pcqjncwl\ndir qwvfpgl\ndir rtmj\ndir shg\ndir tcdmgwp\n$ cd brhvclj\n$ ls\n40016 mtlscfrd.gdr\ndir mvslzl\n112449 npp.vjv\n46674 pbgjwb\ndir qdtls\ndir tfns\n$ cd mvslzl\n$ ls\ndir dngldfww\ndir dzplphqw\n$ cd dngldfww\n$ ls\n102218 wcrvztrh.mzb\n$ cd ..\n$ cd dzplphqw\n$ ls\n93724 brhvclj\n169467 cwqwcjc.lgd\n$ cd ..\n$ cd ..\n$ cd qdtls\n$ ls\ndir nnlzrwgh"))
	content, _ := os.ReadFile(testFile)
	input := strings.Split(string(content), "\n")

	key := ProcessInput(input)

	if len(key) != 6 {
		t.Fatal("failed - length should be 6 but got: ", len(key))
	}
	if key[len(key)-1] != "/brhvclj/mvslzl/dzplphqw 169467" {
		t.Fatal("failed")
	}
}

func TestProcessInput2(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir a\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"))
	content, _ := os.ReadFile(testFile)
	input := strings.Split(string(content), "\n")

	key := ProcessInput(input)
	dirTotals := MapValuesToDirectory(key)

	if GetTotalCount(dirTotals) != 95437 {
		t.Fatal("failed - expected '95437' but got: ", GetTotalCount(dirTotals))
	}
}

func TestProcessInput3(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ cd a\n$ cd b\n$ cd c\n$ ls\n100 blah\n$ cd ..\n$ ls\n100 test\n$ cd ..\n$ ls\n100 another"))
	content, _ := os.ReadFile(testFile)
	input := strings.Split(string(content), "\n")

	key := ProcessInput(input)
	dirTotals := MapValuesToDirectory(key)
	log.Println(key)
	log.Println(dirTotals)

	if GetTotalCount(dirTotals) != 600 {
		t.Fatal("failed - expected '600' but got: ", GetTotalCount(dirTotals))
	}
}

func TestRemoveDir(t *testing.T) {
	dir := "/dir1/dir2"
	if RemoveDir(dir) != "/dir1" {
		t.Fatal("failed - expected '/dir1' but got: ", RemoveDir(dir))
	}
}
