package main

import (
	"beandev.io/beancli/pkg/cmd"
	_ "beandev.io/beancli/pkg/cmd"
	"fmt"
	"io"
	_ "k8s.io/component-base/cli"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	rootCommand := cmd.NewDefaultBeanCommand()
	print("v0.0.2")
	println(rootCommand.Long)
	download_url("https://raw.githubusercontent.com/BeanProjects/beancli/master/scripts/echo.sh")
	chmod("echo.sh")
}

func download_url(fullURLFile string) {

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
}

func chmod(filename string) {
	stats, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File Before: %s\n", stats.Mode())
	err = os.Chmod(filename, 0700)
	if err != nil {
		log.Fatal(err)
	}

	stats, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After: %s\n", stats.Mode())
}
