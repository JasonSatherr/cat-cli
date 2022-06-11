package httpReqTools

import (
	"io/ioutil"
	"net/http"
)

func JsonBytesFromRequest(hR *http.Request) []byte {

	res, _ := http.DefaultClient.Do(hR)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
