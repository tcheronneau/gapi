package gapi

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/url"
  "fmt"
)


func (g *Grafana) GetUsers() ([]User, error) {
  users := make([]User,0)
  return users, g.getRequest("/api/users",nil,nil,&users)
}


func (g *Grafana) GetUser(info string) (User, error) {
  user := User{}
  query := url.Values{}
  query.Add("loginOrEmail",info)
  return user, g.getRequest("/api/users/lookup",query,nil,&user)
}

func (g *Grafana) GetTeamOf(info string) ([]Team, error) {
  teams := make([]Team, 0)
  user, err := g.GetUser(info)
  if err != nil {
    return teams, err
  }
  id := user.Id
  url := fmt.Sprintf("/api/users/%d/teams", id)
  return teams, g.getRequest(url,nil,nil,&teams)
}

func (g *Grafana) GetOrgOf(info string) ([]Org, error) {
  orgs := make([]Org, 0)
  user, err := g.GetUser(info)
  if err != nil {
    return orgs, err
  }
  id := user.Id
  url := fmt.Sprintf("/api/users/%d/orgs", id)
  return orgs, g.getRequest(url,nil,nil,&orgs)
}
