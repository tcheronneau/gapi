package gapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
  "fmt"
  "bytes"
)

type respOrg struct {
  TotalCount int
  Orgs []Org
  Page int
  PerPage int
}


func (g *Grafana) getOrgId(orgname string) (id int) {
  org, _ := g.GetOrgByName(orgname)
  id = org.Id
  return id
}

func (g *Grafana) GetOrgs() ([]Org, error) {
  orgs := make([]Org, 0)
  req, err := g.newRequest("GET", "/api/orgs", nil,nil)
  if err != nil {
    return orgs,err
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
  return orgs, err
}

func (g *Grafana) GetOrgById(orgid int) (Org, error) {
  org := Org{}
  url := fmt.Sprintf("/api/orgs/%d", orgid)
  req, err := g.newRequest("GET", url, nil,nil)
  if err != nil {
    return org, err
  }
  resp, err := g.Do(req)
  if err != nil {
    return org, err
  }
  if resp.StatusCode != 200 {
    return org, errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return org, err
  }
  err = json.Unmarshal(data, &org)
  return org, err
}

func (g *Grafana) GetOrgByName(orgname string) (Org, error) {
  org := Org{}
  url := fmt.Sprintf("/api/orgs/%s", orgname)
	req, err := g.newRequest("GET", url, nil, nil)
	if err != nil {
		return org, err
	}
	resp, err := g.Do(req)
	if err != nil {
		return org, err
	}
	if resp.StatusCode != 200 {
		return org, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	if err != nil {
		return org, err
	}
	return org, err
}


func (g *Grafana) GetOrgMember(org interface{}) ([]User, error){
  users := make([]User, 0)
  id := int(0)
  switch org.(type) {
    case string:
      id = g.getOrgId(org.(string))
    case int:
      id = int(org.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/orgs/%d/users",id)
  req, err := g.newRequest("GET",url,nil,nil)
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



func (g *Grafana) AddOrg(orgname string) (error) {
  dataMap := map[string]string{
    "name": orgname,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  req, err := g.newRequest("POST", "/api/orgs", nil, bytes.NewBuffer(data))
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

func (g *Grafana) DeleteOrg(org interface{}) (error){
  id := int(0)
  switch org.(type) {
    case string:
      id = g.getOrgId(org.(string))
    case int:
      id = int(org.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/orgs/%d",id)
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

func (g *Grafana) UpdateOrg(org interface{}, orgname string) (error){
  id := int(0)
  dataMap := map[string]string{
    "name": orgname,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  switch org.(type) {
    case string:
      id = g.getOrgId(org.(string))
    case int:
      id = int(org.(int))
    default:
      id = 0
  }
  url := fmt.Sprintf("/api/orgs/%d",id)
  req, err := g.newRequest("PUT",url,nil,bytes.NewBuffer(data))
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

func (g *Grafana) AddOrgMember(org, user, role string) error {
  tid := g.getOrgId(org)
  url := fmt.Sprintf("/api/orgs/%d/users",tid)
  dataMap := map[string]string{
    "loginOrEmail": user,
    "role": role,
  }
  data, err := json.Marshal(dataMap)
  req, err := g.newRequest("POST",url,nil,bytes.NewBuffer(data))
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

func (g *Grafana) RemoveOrgMember(org, user string) error {
  u,_ := g.GetUser(user)
  uid := u.Id
  tid := g.getOrgId(org)
  url := fmt.Sprintf("/api/orgs/%d/members/%d",tid,uid)
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
