package gapi

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "fmt"
  "bytes"
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

func (g *Grafana) GetFolderById(id int) (Folder, error) {
  folder := Folder{}
  url := fmt.Sprintf("/api/folders/%d", id)
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

func (g *Grafana) CreateFolder(title string) (error){
  dataMap := map[string]string{
    "title": title,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  req, err := g.newRequest("POST", "/api/folders", nil, bytes.NewBuffer(data))
  if err != nil {
    return err
  }
  resp, err := g.Do(req)
  if err != nil {
    return err
  }
  if resp.StatusCode != 200 {
    return errors.New(resp.Status)
  }
  return err
}

func (g *Grafana) UpdateFolder(uid string, dataMap map[string]string) (error){
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  url := fmt.Sprintf("/api/folders/%s",uid)
  req, err := g.newRequest("PUT", url, nil, bytes.NewBuffer(data))
  if err != nil {
    return err
  }
  resp, err := g.Do(req)
  if err != nil {
    return err
  }
  if resp.StatusCode != 200 {
    return errors.New(resp.Status)
  }
  return err
}

func (g *Grafana) DeleteFolder(uid string) (error){
  url := fmt.Sprintf("/api/folders/%s",uid)
  req, err := g.newRequest("DELETE", url, nil, nil)
  if err != nil {
    return err
  }
  resp, err := g.Do(req)
  if err != nil {
    return err
  }
  if resp.StatusCode != 200 {
    return errors.New(resp.Status)
  }
  return err
}
