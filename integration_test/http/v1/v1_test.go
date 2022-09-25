package v1_test

import (
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..") // change to suit test file location
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	err := healthCheck(_attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", _host, err)
	}

	log.Printf("Integration tests: host %s is available", _host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(_healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf(
			"Integration tests: url %s is not available, attempts left: %d",
			_healthPath,
			attempts,
		)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}
