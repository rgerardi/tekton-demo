package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	a := assert.New(t)

	ts, cleanup := testServer(t)
	defer cleanup()

	hostname, err := os.Hostname()
	a.NoError(err)

	exp := fmt.Sprintf("This REQUEST is being served by sever %s\n", hostname)

	res, err := http.Get(ts)
	a.NoError(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	a.NoError(err)

	a.Equal(exp, string(body))
}

func testServer(t *testing.T) (string, func()) {
	t.Helper()

	s := httptest.NewServer(newMux())

	return s.URL, func() {
		s.Close()
	}
}
