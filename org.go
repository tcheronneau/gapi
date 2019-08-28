package gapi

import (
	"encoding/json"
  "fmt"
  "bytes"
)

type respOrg struct {
  TotalCount int
  Orgs []Org
  Page int
  PerPage int
}

func (g *Grafana) getOrgId(orgname string) int {
  fmt.Printf(orgname)
  org,_ := g.GetOrgByName(orgname)
  fmt.Println(org)
  oid := org.Id
  return oid
}

func (g *Grafana) GetOrgs() ([]Org, error) {
  orgs := make([]Org, 0)
  return orgs, g.getRequest("/api/orgs",nil,nil,&orgs)
}

func (g *Grafana) GetOrgById(orgid int) (Org, error) {
  org := Org{}
  url := fmt.Sprintf("/api/orgs/%d", orgid)
  return org, g.getRequest(url,nil,nil,&org)
}

func (g *Grafana) GetOrgByName(orgname string) (Org, error) {
  org := Org{}
  url := fmt.Sprintf("/api/orgs/name/%s", orgname)
  return org, g.getRequest(url,nil,nil,&org)
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
  return users, g.getRequest(url,nil,nil,&users)
}



func (g *Grafana) AddOrg(orgname string) (error) {
  dataMap := map[string]string{
    "name": orgname,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  return g.postRequest("POST","/api/orgs",nil,bytes.NewBuffer(data))
}

func (g *Grafana) DeleteOrg(org interface{}) (error){
  id := int(0)
  switch org.(type) {
    case string:
      id = g.getOrgId(org.(string))
    case int:
      id = int(org.(int))
  }
  url := fmt.Sprintf("/api/orgs/%d",id)
  return g.postRequest("DELETE",url,nil,nil)
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
  return g.postRequest("PUT",url,nil,bytes.NewBuffer(data))
}

func (g *Grafana) AddOrgMember(org, user, role string) error {
  tid := g.getOrgId(org)
  url := fmt.Sprintf("/api/orgs/%d/users",tid)
  dataMap := map[string]string{
    "loginOrEmail": user,
    "role": role,
  }
  data, err := json.Marshal(dataMap)
  if err != nil {
    return err
  }
  return g.postRequest("POST",url,nil,bytes.NewBuffer(data))
}

func (g *Grafana) RemoveOrgMember(org, user string) error {
  u,_ := g.GetUser(user)
  uid := u.Id
  tid := g.getOrgId(org)
  url := fmt.Sprintf("/api/orgs/%d/members/%d",tid,uid)
  return g.postRequest("DELETE",url,nil,nil)
}
