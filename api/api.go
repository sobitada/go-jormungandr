package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

const DefaultAPIVersion string = "v0"

type JormungandrAPI struct {
    url *url.URL
}

// gets an API instance, which is located at the given
// host URL. The host URL must not include the API path,
// but simply the blank host URL of the jormungandr API
// endpoint e.g. "127.0.0.1:3101".
func GetAPIFromHost(host string) (*JormungandrAPI, error) {
    hostURL, err := url.Parse(host)
    if err == nil && hostURL != nil {
        hostURL, err = combine(hostURL, fmt.Sprintf("/api/%s/", DefaultAPIVersion))
        if err == nil {
            return &JormungandrAPI{url: hostURL}, nil
        }
    }
    return nil, invalidArgument{
        MethodName:  "GetAPIFromHost",
        Expectation: fmt.Sprintf("You must enter a valid host URL, but it was '%v'.", host),
    }
}

// gets an API instance, which can be found at the given URL.
func GetAPIFromURL(url *url.URL) *JormungandrAPI {
    return &JormungandrAPI{url: url}
}

// gets the node statistics of the jormungandr node, the second
// return value indicates whether the node is in bootstrap mode
// or running.
func (api *JormungandrAPI) GetNodeStatistics() (*NodeStatistic, bool, error) {
    apiURL, err := combine(api.url, "./node/stats")
    if err == nil {
        var response, err = http.Get(apiURL.String())
        if err == nil {
            if response.StatusCode == 200 {
                data, parseErr := ioutil.ReadAll(response.Body)
                if parseErr == nil {
                    var response nodeStatisticJSON
                    unmarshalError := json.Unmarshal(data, &response)
                    if unmarshalError == nil {
                        if response.State != "Running" {
                            return nil, true, nil
                        } else {
                            return transformJSONToNodeStatistic(response), false, nil
                        }
                    }
                    return nil, false, unmarshalError
                }
                return nil, false, parseErr
            }
            return nil, false, apiCallFailed{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
        }
        return nil, false, err
    }
    return nil, false, err
}

// gets the leader ids that are registered at this node.
func (api *JormungandrAPI) GetRegisteredLeaders() ([]uint64, error) {
    var apiURL, err = combine(api.url, "./leaders")
    if err == nil {
        var response, err = http.Get(apiURL.String())
        if err == nil {
            if response.StatusCode == 200 {
                var data, parseErr = ioutil.ReadAll(response.Body)
                if parseErr == nil {
                    var response []uint64
                    var parseErr = json.Unmarshal(data, &response)
                    if parseErr == nil {
                        return response, nil
                    }
                }
            }
        }
    }
    return nil, err
}

// removes the leader with the given from this node. returns true, if
// such a leader existed, otherwise false.
func (api *JormungandrAPI) RemoveRegisteredLeader(leaderID uint64) (bool, error) {
    var apiURL, err = combine(api.url, fmt.Sprintf("./leaders/%v", leaderID))
    if err == nil {
        client := &http.Client{}
        var request, err = http.NewRequest(http.MethodDelete, apiURL.String(), nil)
        if err == nil {
            var response, err = client.Do(request)
            if err == nil {
                if response.StatusCode == 200 {
                    return true, nil
                } else if response.StatusCode == 404 {
                    return false, nil
                } else {
                    return false, apiCallFailed{
                        ApiPath:    apiURL.String(),
                        StatusCode: response.StatusCode,
                        Message:    fmt.Sprintf("Failed to remove leader with ID='%v'.", leaderID)}
                }
            }
        }
    }
    return false, err
}

// sends the given leader certificate and registers this leader. if successful, then the
// id of the registered leader is returned.
func (api *JormungandrAPI) PostLeader(leaderCertificate LeaderCertificate) (uint64, error) {
    var apiURL, err = combine(api.url, fmt.Sprintf("./leaders"))
    if err == nil {
        certJSONData, err := CertToJSON(leaderCertificate)
        if err == nil {
            response, err := http.Post(apiURL.String(), "application/json", bytes.NewReader(certJSONData))
            if err == nil {
                responseData, err := ioutil.ReadAll(response.Body)
                if err == nil {
                    var leaderID uint64
                    err := json.Unmarshal(responseData, &leaderID)
                    if err == nil {
                        return leaderID, nil
                    }
                    return 0, err
                }
                return 0, err
            }
            return 0, err
        }
        return 0, err
    }
    return 0, err
}

// gets the leader schedule of the jormunagndr node.
func (api *JormungandrAPI) GetLeadersSchedule() ([]LeaderAssignment, error) {
    apiURL, err := combine(api.url, "./leaders/logs")
    if err == nil {
        var response, err = http.Get(apiURL.String())
        if err == nil {
            if response.StatusCode == 200 {
                var data, parseErr = ioutil.ReadAll(response.Body)
                if parseErr == nil {
                    var response []LeaderAssignment
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

// sends a shutdown request to the jormungandr node.
func (api *JormungandrAPI) Shutdown() error {
    apiURL, err := combine(api.url, "./shutdown")
    if err == nil {
        var response, err = http.Get(apiURL.String())
        if err == nil {
            if response.StatusCode == 200 {
                return nil //success
            } else {
                return apiCallFailed{ApiPath: apiURL.String(), StatusCode: response.StatusCode, Message: response.Status}
            }
        }
    }
    return err
}
