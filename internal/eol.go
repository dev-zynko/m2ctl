package internal

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func CheckEOL(stdin io.Writer, stdout io.Reader) (string, bool) {
	stdin.Write([]byte("freebsd-version \n"))

	val, err := GetOutput(stdout)
	if err != nil {
		log.Fatal("Failed reading output EOF", err)
	}

	var wg1 sync.WaitGroup
	var isEol bool
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		isEol = checkFreeBSDSite(val)
	}()
	wg1.Wait()
	return val, isEol
}

func checkFreeBSDSite(version string) bool {
	resp, err := http.Get("https://www.freebsd.org/security/unsupported/")
	if err != nil {
		log.Fatal("Failed requesting freebsd site unsported", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed parsing body", err)
	}

	sb := string(body)

	return strings.Contains(sb, version)
}
