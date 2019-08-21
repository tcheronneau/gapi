package gapi

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/url"
  "fmt"
)


func (g *Grafana) GetUsers() ([]User, error) {
  users := make([]User, 0)
  req, err := g.newRequest("GET", "/api/users", nil, nil)
  if err != nil {
    return users, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return users, err
  }
  if resp.StatusCode != 200 {
    return users, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return users, err
  }
  err = json.Unmarshal(data, &users)
  if err != nil {
    return users, err
  }
  return users, err
}

func (g *Grafana) GetUser(info string) (User, error) {
  user := User{}
  query := url.Values{}
  query.Add("loginOrEmail",info)
  req, err := g.newRequest("GET","/api/users/lookup",query,nil)
  if err != nil {
    return user, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return user, err
  }
  if resp.StatusCode != 200 {
    return user, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return user, err
  }
  tmp := User{}
  err = json.Unmarshal(data, &tmp)
  if err != nil {
    return user, err
  }
  user = User(tmp)
  return user,err
}

func (g *Grafana) GetTeamOf(info string) ([]Team, error) {
  teams := make([]Team, 0)
  user, err := g.GetUser(info)
  if err != nil {
    return teams, err
  }
  id := user.Id
  url := fmt.Sprintf("/api/users/%d/teams", id)
  req, err := g.newRequest("GET", url, nil,nil)
    if err != nil {
    return teams, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return teams, err
  }
  if resp.StatusCode != 200 {
    return teams, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return teams, err
  }
  err = json.Unmarshal(data, &teams)
  if err != nil {
    return teams, err
  }
  return teams, err
}

func (g *Grafana) GetOrgOf(info string) ([]Org, error) {
  orgs := make([]Org, 0)
  user, err := g.GetUser(info)
  if err != nil {
    return orgs, err
  }
  id := user.Id
  url := fmt.Sprintf("/api/users/%d/orgs", id)
  req, err := g.newRequest("GET", url, nil,nil)
    if err != nil {
    return orgs, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return orgs, err
  }
  if resp.StatusCode != 200 {
    return orgs, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return orgs, err
  }
  err = json.Unmarshal(data, &orgs)
  if err != nil {
    return orgs, err
  }
  return orgs, err
}
