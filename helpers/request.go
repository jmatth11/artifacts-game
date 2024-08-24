package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jmatth11/artifacts-game/types"
)

const baseURL = "https://api.artifactsmmo.com"

func addHeaders(client types.Client, req *http.Request) {
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token))
}

func executeRequest(req *http.Request) (types.Response, error) {
  var result types.Response
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return result, err
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return result, err
  }
  result.Data = body
  result.Type = types.ServerCode(resp.StatusCode)
  return result, nil
}

func Post(client types.Client, path string, body interface{}) (types.Response, error) {
  var bodyData = []byte{}
  var err error
  if body != nil {
    bodyData, err = json.Marshal(body)
    if err != nil {
      return types.Response{}, err
    }
    fmt.Println(string(bodyData))
  }
  responseBody := bytes.NewBuffer(bodyData)
  fullPath,err := url.JoinPath(baseURL, path)
  if err != nil {
    return types.Response{}, nil
  }
  req, err := http.NewRequest(http.MethodPost, fullPath, responseBody)
  if err != err {
    return types.Response{}, err
  }
  addHeaders(client, req)
  req.Header.Add("Content-Type", "application/json")
  return executeRequest(req)
}

func Get(client types.Client, path string, params interface{}) (types.Response, error) {
  queryParams := url.Values{}
  objMap := make(map[string]string)
  if params != nil {
    objMap = MapFromStruct(params)
  }
  for k, v := range objMap {
    queryParams.Add(k, v)
  }
  fullPath,err := url.JoinPath(baseURL, path)
  if err != nil {
    return types.Response{}, err
  }
  if len(queryParams) > 0 {
    fullPath = fmt.Sprintf("%s?%s", fullPath, queryParams.Encode())
  }
  req, err := http.NewRequest(http.MethodGet, fullPath, nil)
  if err != err {
    return types.Response{}, err
  }
  addHeaders(client, req)
  return executeRequest(req)
}
