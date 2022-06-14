//Package httpReqTools will provide functions to work with httpRequests
package httpReqTools

import (
	"io/ioutil"
	"net/http"
)

//JsonBytesFromRequest is a function that makes a request and then returns the bytes from the body
func JsonBytesFromRequest(hR *http.Request) []byte {
	//res is the result of executing the request
	res, err := http.DefaultClient.Do(hR)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//body is the full bytes slice of the request body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}
