package api

import (
    "fmt"
)

type invalidArgument struct {
    MethodName  string
    Expectation string
}

func (iArg invalidArgument) Error() string {
    return fmt.Sprintf("Invalid argument passed to %v. %v", iArg.MethodName, iArg.Expectation)
}

type apiCallFailed struct {
    ApiPath    string
    StatusCode int
    Message    string
}

func (apiFail apiCallFailed) Error() string {
    return fmt.Sprintf("Exploration request failed with status code %v. %v", apiFail.StatusCode, apiFail.StatusCode)
}

type explorerRequestFailed struct {
    StatusCode int
    Message    string
}

func (apiFail explorerRequestFailed) Error() string {
    return fmt.Sprintf("Exploration request failed with status code %v. %v", apiFail.StatusCode, apiFail.Message)
}
