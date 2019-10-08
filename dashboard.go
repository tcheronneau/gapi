package gapi

import (
  "fmt"
  "encoding/json"
  "bytes"
  "net/url"
)

type TagResponse struct {
	ID          int      `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URI         string   `json:"uri"`
	URL         string   `json:"url"`
	Slug        string   `json:"slug"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderID    int      `json:"folderId"`
	FolderUID   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderURL   string   `json:"folderUrl"`
}

type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Slug      string `json:"slug"`
	Folder    int64  `json:"folderId"`
}

type DashboardSaveResponse struct {
	Slug    string `json:"slug"`
	Id      int64  `json:"id"`
	Uid     string `json:"uid"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}
type Dashboard struct {
	Meta      DashboardMeta          `json:"meta"`
	Model     map[string]interface{} `json:"dashboard"`
	Folder    int64                  `json:"folderId"`
	Overwrite bool                   `json:overwrite`
}

func (g *Grafana) SearchTag(tag string) ([]string, error) {
  tmp := make([]TagResponse,0)
  uid := make([]string,0)
  query := url.Values{}
  query.Add("tag",tag)
  err := g.getRequest("/api/search",query,nil,&tmp)
  for i := 0 ; i < len(tmp); i++ {
    uid = append(uid, tmp[i].UID)
  }
  return uid, err
}

func (g *Grafana) GetDashboard(uid string) (Dashboard, error) {
  dashboard := Dashboard{}
  url := fmt.Sprintf("/api/dashboards/uid/%s",uid)
  err := g.getRequest(url,nil,nil,&dashboard)
  if err != nil {
    return dashboard, err
  }
  dashboard.Folder = dashboard.Meta.Folder
  return dashboard, nil
}

func (g *Grafana) UpdateDashboard(dashboard Dashboard) (error) {
	data, err := json.Marshal(dashboard)
	if err != nil {
		return err
	}
  return g.postRequest("POST","/api/dashboards/db",nil, bytes.NewBuffer(data))
}
