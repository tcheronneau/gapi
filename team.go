package gapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
  req, err := g.newRequest("GET", url, nil,nil)
  if err != nil {
    return team, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return team, err
  }
  if resp.StatusCode != 200 {
    return team, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return team, err
  }
  err = json.Unmarshal(data, &team)
  return team, err
}

func (g *Grafana) GetTeamByName(teamname string) ([]Team, error) {
	tmp := respTeam{}
  teams := make([]Team, 0)
  query := url.Values{}
  query.Add("name",teamname)
	req, err := g.newRequest("GET", "/api/teams/search", query, nil)
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
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return teams, err
	}
  teams = tmp.Teams
	return teams, err
}

func (g *Grafana) getTeamId(teamname string) (int, error) {
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
  req, err := g.newRequest("POST", "/api/teams", nil, bytes.NewBuffer(data))
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

func (g *Grafana) DeleteTeam(team interface{}) (error){
  id := int(0)
  switch team.(type) {
    case string:
      id, _ = g.getTeamId(team.(string))
    case int:
      id = int(team.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/teams/%d",id)
  req, err := g.newRequest("DELETE",url,nil,nil)
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