package gapi

import (
  "net/url"
  "fmt"
)


func (g *Grafana) GetUsers() ([]User, error) {
  users := make([]User,0)
  return users, g.getRequest("/api/users",nil,nil,&users)
}


func (g *Grafana) GetUserId(user string) (int, error) {
  id := int(0)
  u,err := g.GetUser(user)
  if err != nil {
    return id, err
  }
  id = u.Id
  return id, err
}

func (g *Grafana) GetUser(info string) (User, error) {
  user := User{}
  query := url.Values{}
  query.Add("loginOrEmail",info)
  return user, g.getRequest("/api/users/lookup",query,nil,&user)
}

func (g *Grafana) GetUserById(id int) (User, error) {
  user := User{}
  url := fmt.Sprintf("/api/users/%d",id)
  return user, g.getRequest(url,nil,nil,&user)
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


func (g *Grafana) DeleteUser(user string) error {
	id, err := g.GetUserId(user)
  if err != nil {
		return err
  }
	url := fmt.Sprintf("/api/admin/users/%d", id)
  return g.postRequest("DELETE",url,nil,nil)
}
