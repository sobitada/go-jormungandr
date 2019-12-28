package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/sobitada/go-jormungandr/dto"
	"io/ioutil"
	"net/http"
	"net/url"
)

const DefaultAPIVersion string = "v0"

type JormungandrAPI interface {
	// gets the node statistics of the jormungandr node.
	GetNodeStatistics() (*dto.NodeStatistic, error)
	// gets the leader logs of the jormunagndr node.
	GetLeadersLogs() ([]dto.LeaderAssignment, error)
	// sends a shutdown request to the jormungandr node.
	Shutdown() error
}

type InvalidArgumentException struct {
	MethodName  string
	Expectation string
}

func (iArg InvalidArgumentException) Error() string {
	return fmt.Sprintf("Invalid argument passed to %v. %v", iArg.MethodName, iArg.Expectation)
}

type APICallFailedException struct {
	ApiPath    string
	StatusCode int
	Message    string
}

func (apiFail APICallFailedException) Error() string {
	return fmt.Sprintf("Exploration request failed with status code %v. %v", apiFail.StatusCode, apiFail.StatusCode)
}

type jormungandrAPIImpl struct {
	url *url.URL
}

// gets an API instance, which is located at the given
// host URL. The host URL must not include the API path,
// but simply the blank host URL of the jormungandr API
// endpoint.
func GetAPIFromHost(host string) (JormungandrAPI, error) {
	hostURL, err := url.Parse(host)
	if err == nil && hostURL != nil {
		hostURL, err = combine(hostURL, fmt.Sprintf("/api/%s/", DefaultAPIVersion))
		if err == nil {
			return jormungandrAPIImpl{url: hostURL}, nil
		}
	}
	return nil, InvalidArgumentException{
		MethodName:  "GetAPIFromHost",
		Expectation: fmt.Sprintf("You must enter a valid host URL, but it was '%v'.", host),
	}
}

func (api jormungandrAPIImpl) GetNodeStatistics() (*dto.NodeStatistic, error) {
	apiURL, err := combine(api.url, "./node/stats")
	if err == nil {
		var response, err = http.Get(apiURL.String())
		if err == nil {
			if response.StatusCode == 200 {
				var data, parseErr = ioutil.ReadAll(response.Body)
				if parseErr == nil {
					var response dto.NodeStatistic
					var parseErr = json.Unmarshal(data, &response)
					if parseErr == nil {
						return &response, nil
					}
				}
				return nil, parseErr
			} else {
				return nil, APICallFailedException{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
			}
		}
	}
	return nil, err
}

func (api jormungandrAPIImpl) GetLeadersLogs() ([]dto.LeaderAssignment, error) {
	apiURL, err := combine(api.url, "./leaders/logs")
	if err == nil {
		var response, err = http.Get(apiURL.String())
		if err == nil {
			if response.StatusCode == 200 {
				var data, parseErr = ioutil.ReadAll(response.Body)
				if parseErr == nil {
					var response []dto.LeaderAssignment
					var parseErr = json.Unmarshal(data, &response)
					if parseErr == nil {
						return response, nil
					}
				}
				return nil, parseErr
			}
		}
	}
	return nil, err
}

func (api jormungandrAPIImpl) Shutdown() error {
	apiURL, err := combine(api.url, "./shutdown")
	if err == nil {
		var response, err = http.Get(apiURL.String())
		if err == nil {
			if response.StatusCode == 200 {
				return nil //success
			} else {
				return APICallFailedException{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
			}
		}
	}
	return err
}
