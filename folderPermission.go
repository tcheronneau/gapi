package gapi

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "fmt"
  "bytes"
)

func (g *Grafana) GetFolderPermissions(uid string) ([]FolderPermission, error){
  perms := make([]FolderPermission,0)
  url := fmt.Sprintf("/api/folders/%s/permissions",uid)
  req, err := g.newRequest("GET",url,nil,nil)
  if err != nil {
    return perms,err
  }
  resp, err := g.Do(req)
  if err != nil {
    return perms, err
  }
  if resp.StatusCode != 200 {
    return perms, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return perms, err
  }
  err = json.Unmarshal(data, &perms)
  return perms, err
}

func (g *Grafana) UpdateFolderPermissions(uid string, items []map[string]string) (error){
  data, err := json.Marshal(items)
  if err != nil {
    return err
  }
  url := fmt.Sprintf("/api/folders/%s/permissions",uid)
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
