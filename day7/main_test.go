package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

var testFile = "test.txt"

func setup(fileContent []byte) {
	if err := os.WriteFile(testFile, fileContent, 777); err != nil {
		log.Fatal(err)
	}
}

func clear() {
	if err := os.Remove(testFile); err != nil {
		log.Fatal(err)
	}
}

func TestProcessItemsDir(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("dir brhvclj\ndir clnvqg\ndir dtqtvvrn\ndir lcz\ndir pcqjncwl\ndir qwvfpgl\ndir rtmj\ndir shg\ndir tcdmgwp"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)

	ProcessItems(item, scanner)
	if len(item.children) != 9 {
		t.Fatal("Failed")
	}
}

func TestProcessItemsFiles(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("93724 brhvclj\n169467 cwqwcjc.lgd"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)

	ProcessItems(item, scanner)

	if len(item.children) != 2 {
		t.Fatal("failed")
	}
}

func TestFirstCd(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir brhvclj\ndir clnvqg\n$ cd brhvclj\n$ ls\n1 file.txt"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)

	Analyzer(item, scanner)

	if len(item.children["brhvclj"].children) != 1 {
		t.Fatal("failed")
	}
}

func TestCdBack(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir brhvclj\ndir clnvqg\n$ cd brhvclj\n$ ls\n1 file.txt\n$ cd ..\n$ ls\n2 file2.txt"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)

	Analyzer(item, scanner)

	if len(item.children) != 3 && len(item.children["brhvclj"].children) != 1 {
		t.Fatal("failed")
	}
}

func TestFileStructureCount(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)
	Analyzer(item, scanner)
	var counter int64
	counter = 0
	Counter(item, &counter)

	if item.size != 48381165 {
		t.Fatal("failed")
	}
}

func TestTotalSumOfSizeLessThanOrEqualTo(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"))
	item := NewItem(Dir)
	item.name = "main"
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewReader(r)
	Analyzer(item, scanner)
	var counter int64
	var totalSize int64
	counter = 0
	totalSize = 0
	Counter(item, &counter)
	TotalSumOfSizeLessThanOrEqualTo(100000, item, &totalSize)

	if totalSize != 95437 {
		t.Fatal("failed - expected '95437' but got: ", totalSize)
	}
}
