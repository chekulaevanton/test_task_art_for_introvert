package utils

import (
    "fmt"
    "strings"
)

func GetErrorJSONResponse(errText string) string {
    return fmt.Sprintf("{\"error\": \"%s\"}", errText)
}

func GetPartsFromURL(url string) []string {
    return strings.Split(strings.Trim(url, "/"), "/")
}
