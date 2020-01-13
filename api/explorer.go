package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type JormungandrExplorer struct {
	host *url.URL
}

// gets an explorer instance, which is located at the given
// host URL. The host URL must not include the explorer path,
// but simply the blank host URL of the jormungandr node on
// which the explorer is enabled.
func GetExplorer(host string) (*JormungandrExplorer, error) {
	hostUrl, err := url.Parse(host)
	if err == nil && hostUrl != nil {
		return &JormungandrExplorer{host: hostUrl}, nil
	}
	return nil, invalidArgument{
		MethodName:  "GetExplorer",
		Expectation: fmt.Sprintf("You must enter a valid host URL, but it was '%v'.", host),
	}
}

// sends the given GraphQL query to the explorer and
// returns the responding JSON data as a byte array.
func (api *JormungandrExplorer) Explore(graphQL string) ([]byte, error) {
	var explorerUrl, err = combine(api.host, "explorer/graphql")
	if err == nil {
		var response, err = http.Post(explorerUrl.String(), "application/json", strings.NewReader(fmt.Sprintf("\"query\":\"%s\"", graphQL)))
		if err == nil {
			if response.StatusCode == 200 {
				return ioutil.ReadAll(response.Body)
			} else {
				return nil, explorerRequestFailed{
					StatusCode: response.StatusCode,
					Message:    response.Status,
				}
			}
		}
	}
	return nil, err
}
