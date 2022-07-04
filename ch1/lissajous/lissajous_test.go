package lissajous

import (
	"math/rand"
	"os"
	"testing"
	"time"
)



func Test(t *testing.T) {
	filePath := "./test"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		t.Fatalf("file %s create failed: %v", filePath, err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	Lissajous(file)

	if _, err := os.Stat(filePath); err != nil {
		t.Fatalf("no %s: %v", filePath, err)
	}
}