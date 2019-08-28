package gapi

import (
  "encoding/json"
  "fmt"
  "bytes"
)


func (g *Grafana) GetFolders() ([]Folder, error) {
  folders := make([]Folder, 0)
  return folders, g.getRequest("/api/folders",nil,nil,&folders)
}


func (g *Grafana) GetFolderByUid(uid string) (Folder, error) {
  folder := Folder{}
  url := fmt.Sprintf("/api/folders/%s", uid)
  tmp := Folder{}
  err := g.getRequest(url,nil,nil,&tmp)
  folder = Folder(tmp)
  return folder, err
}

func (g *Grafana) GetFolderById(id int) (Folder, error) {
  folder := Folder{}
  url := fmt.Sprintf("/api/folders/%d", id)
  tmp := Folder{}
  err := g.getRequest(url,nil,nil,&tmp)
  folder = Folder(tmp)
  return folder, err
}

func (g *Grafana) CreateFolder(title string) (error){
  dataMap := map[string]string{
    "title": title,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  return g.postRequest("POST", "/api/folders", nil, bytes.NewBuffer(data))
}

func (g *Grafana) UpdateFolder(uid string, dataMap map[string]string) (error){
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  url := fmt.Sprintf("/api/folders/%s",uid)
  return g.postRequest("PUT", url, nil, bytes.NewBuffer(data))
}

func (g *Grafana) DeleteFolder(uid string) (error){
  url := fmt.Sprintf("/api/folders/%s",uid)
  return g.postRequest("DELETE", url, nil, nil)
}
