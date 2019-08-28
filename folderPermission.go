package gapi

import (
  "encoding/json"
  "fmt"
  "bytes"
)

func (g *Grafana) GetFolderPermissions(uid string) ([]FolderPermission, error){
  perms := make([]FolderPermission,0)
  url := fmt.Sprintf("/api/folders/%s/permissions",uid)
  return perms, g.getRequest(url,nil,nil,&perms)
}

func (g *Grafana) UpdateFolderPermissions(uid string, items []map[string]string) (error){
  data, err := json.Marshal(items)
  if err != nil {
    return err
  }
  url := fmt.Sprintf("/api/folders/%s/permissions",uid)
  return g.postRequest("PUT", url, nil, bytes.NewBuffer(data))
}
