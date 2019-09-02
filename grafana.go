package gapi

import (
  "net/url"
  "strings"
  "fmt"
  "net/http"
	"io"
	"path"
  "errors"
  "encoding/json"
  "io/ioutil"
)
type Grafana struct {
  key string
  baseURL url.URL
  *http.Client
}

type User struct {
  Id int
  Name string
  Login string
  Email string
  Password string
  IsAdmin bool
}

type Team struct {
  Id int
  OrgId int
  Name string
  Email string
  AvatarURL string
  MemberCount int
}

type Org struct {
  Id int
  Name string
}

type Folder struct {
  Id int
  Uid string
  Title string
  Url string
  HasACL bool
  CanSave bool
  CanEdit bool
  CanAdmin bool
  CreatedBy string
  Created string
  UpdatedBy string
  Updated string
  Version int
}

type FolderPermission struct {
  Id int
  Oid int
  Created string
  Updated string
  UserId int
  UserLogin string
  UserEmail string
  TeamId int
  Team string
  Role string
  Permission int
  PermissionName string
  Uid string
  Title string
  Slug string
  IsFolder bool
  Url string
}

func New(auth string, baseURL string) (*Grafana, error) {
  u, err := url.Parse(baseURL)
  if err != nil {
    return nil, err
  }
  key := ""
  if strings.Contains(auth,":") {
    split := strings.Split(auth,":")
    u.User = url.UserPassword(split[0],split[1])
  } else {
    key = fmt.Sprintf("Bearer %s", auth)
  }
  return &Grafana{key,*u,&http.Client{}}, nil
}

func (g *Grafana) newRequest(method, requestPath string, query url.Values, body io.Reader) (*http.Request, error) {
  url := g.baseURL
  url.Path = path.Join(url.Path, requestPath)
  url.RawQuery = query.Encode()
  req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return req, err
	}
	if g.key != "" {
		req.Header.Add("Authorization", g.key)
	}

	req.Header.Add("Content-Type", "application/json")
	return req, err
}

func (g *Grafana) getRequest(requestPath string, query url.Values, body io.Reader, obj interface{}) error {
  req, err := g.newRequest("GET",requestPath, query, body)
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
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }
  return json.Unmarshal(data,obj)
}

func (g *Grafana) postRequest(method ,requestPath string, query url.Values, body io.Reader) error {
  req, err := g.newRequest(method,requestPath,query,body)
  if err != nil {
    return err
  }
  resp, err := g.Do(req)
  if err != nil {
    return err
  }
  if resp.StatusCode != 200 {
    return errors.New(resp.StatusCode)
  }
  return err
}
