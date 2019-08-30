package gapi

import (
  "encoding/json"
  "net/url"
  "fmt"
  "bytes"
)

type respTeam struct {
  TotalCount int
  Teams []Team
  Page int
  PerPage int
}

func (g *Grafana) GetTeamById(teamid int) (Team, error) {
  team := Team{}
  url := fmt.Sprintf("/api/teams/%d", teamid)
  return team, g.getRequest(url,nil,nil,&team)
}

func (g *Grafana) GetTeamByName(teamname string) ([]Team, error) {
	tmp := respTeam{}
  teams := make([]Team, 0)
  query := url.Values{}
  query.Add("name",teamname)
  err := g.getRequest("/api/teams/search", query, nil, &tmp)
  teams = tmp.Teams
  return teams, err
}

func (g *Grafana) GetTeamMember(team interface{}) ([]User, error){
  users := make([]User, 0)
  id := int(0)
  switch team.(type) {
    case string:
      id, _ = g.GetTeamId(team.(string))
    case int:
      id = int(team.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/teams/%d/members",id)
  return users, g.getRequest(url,nil,nil,&users)
}


func (g *Grafana) GetTeamId(teamname string) (int, error) {
  id := int(0)
  teams, err := g.GetTeamByName(teamname)
  if err != nil {
    return id, err
  }
  for _,t := range(teams) {
    if t.Name == teamname {
      id = t.Id
    } else {
      id = 0
    }
  }
  return id, err
}


func (g *Grafana) AddTeam(teamname, teamemail string) (error) {
  dataMap := map[string]string{
    "name": teamname,
    "email": teamemail,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  return g.postRequest("POST", "/api/teams", nil, bytes.NewBuffer(data))
}

func (g *Grafana) DeleteTeam(team interface{}) (error){
  id := int(0)
  switch team.(type) {
    case string:
      id, _ = g.GetTeamId(team.(string))
    case int:
      id = int(team.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/teams/%d",id)
  return g.postRequest("DELETE",url,nil,nil)
}

func (g *Grafana) UpdateTeam(team interface{}, teamname string) (error){
  id := int(0)
  dataMap := map[string]string{
    "name": teamname,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  switch team.(type) {
    case string:
      id, _ = g.GetTeamId(team.(string))
    case int:
      id = int(team.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/teams/%d",id)
  return g.postRequest("PUT",url,nil,bytes.NewBuffer(data))
}

func (g *Grafana) AddTeamMember(user, team string) error {
  u,_ := g.GetUser(user)
  uid := u.Id
  tid, _ := g.GetTeamId(team)
  url := fmt.Sprintf("/api/teams/%d/members",tid)
  dataMap := map[string]int{
    "userId": uid,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  return g.postRequest("POST",url,nil,bytes.NewBuffer(data))
}

func (g *Grafana) RemoveTeamMember(user, team string) error {
  u,_ := g.GetUser(user)
  uid := u.Id
  tid, _ := g.GetTeamId(team)
  url := fmt.Sprintf("/api/teams/%d/members/%d",tid,uid)
  return g.postRequest("DELETE",url,nil,nil)
}
