package gapi

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "fmt"
)


func (g *Grafana) GetFolders() ([]Folder, error) {
  folders := make([]Folder, 0)
  req, err := g.newRequest("GET", "/api/folders", nil, nil)
  if err != nil {
    return folders, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return folders, err
  }
  if resp.StatusCode != 200 {
    return folders, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return folders, err
  }
  err = json.Unmarshal(data, &folders)
  if err != nil {
    return folders, err
  }
  return folders, err
}


func (g *Grafana) GetFolderByUid(uid string) (Folder, error) {
  folder := Folder{}
  url := fmt.Sprintf("/api/folders/%s", uid)
  req, err := g.newRequest("GET",url,nil,nil)
  if err != nil {
    return folder, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return folder, err
  }
  if resp.StatusCode != 200 {
    return folder, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return folder, err
  }
  tmp := Folder{}
  err = json.Unmarshal(data, &tmp)
  if err != nil {
    return folder, err
  }
  folder = Folder(tmp)
  return folder,err
}

func (g *Grafana) CreateFolder(title string) error {
}

