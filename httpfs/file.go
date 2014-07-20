package httpfs

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// A https File implements a subset of os.File's methods.
type File struct {
	name   string // local file name passed to Open
	client *Client
	u      url.URL // url to access file on remote machine
	fd     uintptr // file descriptor on server
}

func (f *File) Read(p []byte) (n int, err error) {
	// send READ request
	u := f.u // (a copy)
	q := u.Query()
	q.Set("n", fmt.Sprint(len(p))) // number of bytes to read
	u.RawQuery = q.Encode()
	req, eReq := http.NewRequest("READ", u.String(), nil)
	panicOn(eReq)
	resp, eResp := f.client.client.Do(req)
	if eResp != nil {
		return 0, fmt.Errorf(`httpfs read "%v": %v`, f.name, eResp)
	}

	// read response
	defer resp.Body.Close()
	nRead, eRead := resp.Body.Read(p)
	errStr := resp.Header.Get(X_READ_ERROR)
	if errStr != "" {
		return nRead, errors.New(errStr) // other than EOF error goes to header
	}
	return nRead, eRead // passes on EOF or http errors
}

func (f *File) Write(p []byte) (n int, err error) {
	// send WRITE request
	req, eReq := http.NewRequest("WRITE", f.u.String(), bytes.NewBuffer(p))
	panicOn(eReq)
	resp, eResp := f.client.client.Do(req)
	if eResp != nil {
		return 0, fmt.Errorf(`httpfs write "%v": %v`, f.name, eResp)
	}

	defer resp.Body.Close()
	body, eBody := ioutil.ReadAll(resp.Body)
	if eBody != nil {
		return 0, fmt.Errorf("httpfs write: %v", eBody) // actually no idea how many bytes written!
	}
	nRead, eNRead := strconv.Atoi(string(body))
	if eNRead != nil {
		return 0, fmt.Errorf("httpfs write: bad response: %v", eNRead)
	}

	return nRead, nil
}

func (f *File) Close() error {
	return nil
}
