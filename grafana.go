package gapi

import (
  "net/url"
  "strings"
  "fmt"
  "net/http"
	"io"
	"path"
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
  OrgId int
  Name string
  Role string
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
