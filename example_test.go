package dropy_test

import (
	"io"
	"os"
	"strings"

	"github.com/segmentio/go-env"
	"github.com/NightMan-1/go-dropbox"
	"github.com/NightMan-1/go-dropy"
)

// Upload and read a file.
func Example() {
	token := env.MustGet("DROPBOX_ACCESS_TOKEN")
	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))

	client.Upload("/demo.txt", strings.NewReader("Hello World"))
	io.Copy(os.Stdout, client.Open("/demo.txt"))
	// Output: Hello World
}
