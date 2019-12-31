package api

import (
    "encoding/json"
    "fmt"
    dto "github.com/sobitada/go-jormungandr/api/dto"
    "io/ioutil"
    "net/http"
    "net/url"
)

type JormungandrAPI interface {
    // gets the node statistics of the jormungandr node, the second
    // return value indicates whether the node is in bootstrap mode
    // or running.
    GetNodeStatistics() (*NodeStatistic, bool, error)
    // gets the leader logs of the jormunagndr node.
    GetLeadersLogs() ([]dto.LeaderAssignment, error)
    // sends a shutdown request to the jormungandr node.
    Shutdown() error
}

type invalidArgumentException struct {
    MethodName  string
    Expectation string
}

func (iArg invalidArgumentException) Error() string {
    return fmt.Sprintf("Invalid argument passed to %v. %v", iArg.MethodName, iArg.Expectation)
}

type apiCallFailedException struct {
    ApiPath    string
    StatusCode int
    Message    string
}

func (apiFail apiCallFailedException) Error() string {
    return fmt.Sprintf("Exploration request failed with status code %v. %v", apiFail.StatusCode, apiFail.StatusCode)
}

type jormungandrAPIImpl struct {
    url *url.URL
}

const DefaultAPIVersion string = "v0"

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
    return nil, invalidArgumentException{
        MethodName:  "GetAPIFromHost",
        Expectation: fmt.Sprintf("You must enter a valid host URL, but it was '%v'.", host),
    }
}

func (api jormungandrAPIImpl) GetNodeStatistics() (*NodeStatistic, bool, error) {
    apiURL, err := combine(api.url, "./node/stats")
    if err == nil {
        var response, err = http.Get(apiURL.String())
        if err == nil {
            if response.StatusCode == 200 {
                data, parseErr := ioutil.ReadAll(response.Body)
                if parseErr == nil {
                    var response dto.NodeStatistic
                    unmarshalError := json.Unmarshal(data, &response)
                    if unmarshalError == nil {
                        if response.State != "Running" {
                            return nil, true, nil
                        } else {
                            nodeStatistic, err := getNodeStatisticFromDto(response)
                            if err == nil {
                                return nodeStatistic, false, nil
                            }
                            return nil, false, err
                        }
                    }
                    return nil, false, unmarshalError
                }
                return nil, false, parseErr
            }
            return nil, false, apiCallFailedException{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
        }
        return nil, false, err
    }
    return nil, false, err
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
                return apiCallFailedException{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
            }
        }
    }
    return err
}
