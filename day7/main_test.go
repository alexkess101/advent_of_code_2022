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
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

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
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	ProcessItems(item, scanner)

	if len(item.children) != 2 {
		t.Fatal("failed")
	}
}

func TestProcessItemsBreaks(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	setup([]byte("dir qwvfpgl\ndir rtmj\ndir shg\ndir tcdmgwp\n$ cd brhvclj"))
	item := NewItem(Dir)
	r, _ := os.Open(testFile)
	defer r.Close()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	ProcessItems(item, scanner)
	if len(item.children) != 4 {
		t.Fatal("failed")
	}

	log.Println(scanner.Text())
}
