package wrapper

import (
    "encoding/json"
    "fmt"
    "github.com/sobitada/go-jormungandr/dto"
    "io/ioutil"
    "net/http"
    "net/url"
)

const API_VERSION string = "v0"

type JormungandrAPI interface {
    // gets the node statistics of the jormungandr node.
    GetNodeStatistics() (*dto.NodeStatistic, error)
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
    host url.URL
}

// gets an API instance, which is located at the given
// host URL. The host URL must not include the API path,
// but simply the blank host URL of the jormungandr API
// endpoint.
func GetAPI(host string) (JormungandrAPI, error) {
    hostUrl, err := url.Parse(host)
    if err == nil && hostUrl != nil {
        return jormungandrAPIImpl{host: *hostUrl}, nil
    }
    return nil, InvalidArgumentException{
        MethodName:  "GetAPI",
        Expectation: fmt.Sprintf("You must enter a valid host URL, but it was '%v'.", host),
    }
}

func (api jormungandrAPIImpl) GetNodeStatistics() (*dto.NodeStatistic, error) {
    apiPath := fmt.Sprintf("/api/%s/node/stats", API_VERSION)
    apiUrl, err := combine(api.host, apiPath)
    if err == nil {
        var response, err = http.Get(apiUrl)
        if err == nil {
            if response.StatusCode == 200 {
                data, parseErr := ioutil.ReadAll(response.Body)
                if parseErr == nil {
                    var response dto.NodeStatistic
                    var parseErr = json.Unmarshal(data, &response)
                    if parseErr == nil {
                        return &response, nil
                    }
                }
                return nil, parseErr
            } else {
                return nil, APICallFailedException{ApiPath: apiPath, StatusCode: response.StatusCode, Message: response.Status}
            }
        }
    }
    return nil, err
}

func (api jormungandrAPIImpl) Shutdown() error {
    apiPath := fmt.Sprintf("/api/%s/shutdown", API_VERSION)
    apiUrl, err := combine(api.host, apiPath)
    if err == nil {
        var response, err = http.Get(apiUrl)
        if err == nil {
            if response.StatusCode == 200 {
                return nil //success
            } else {
                return APICallFailedException{ApiPath: apiPath, StatusCode: response.StatusCode, Message: response.Status}
            }
        }
    }
    return err
}
