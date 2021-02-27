package checker

import (
	"fmt"
	"net/http"
	"time"
)

type Checker interface {
	Check() error
}

type RestChecker struct {
	url    string
	client *http.Client
}

func (rc *RestChecker) Configure(url string, timeout int) {
	fmt.Println(url)
	rc.url = url
	rc.client = &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
}

func (rc *RestChecker) Check() error {
	resp, err := rc.client.Get(rc.url)

	if resp != nil && resp.StatusCode != 200 {
		return fmt.Errorf("Failed status code, expected [200] got [%d]", resp.StatusCode)
	}

	return err
}
