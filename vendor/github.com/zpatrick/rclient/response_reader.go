package rclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zpatrick/go-series"
)

// A ResponseReader attempts to read a *http.Response into v.
type ResponseReader func(resp *http.Response, v interface{}) error

// ReadJSONResponse attempts to marshal the response body into v
// if and only if the response StatusCode is in the 200 range.
// Otherwise, an error is thrown.
// It assumes the response body is in JSON format.
func ReadJSONResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	switch {
	case !series.Ints(200, 299).Contains(resp.StatusCode):
		return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	case v == nil:
		return nil
	default:
		return json.NewDecoder(resp.Body).Decode(v)
	}
}
