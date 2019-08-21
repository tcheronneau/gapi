# gapi

Inspired from https://github.com/nytm/go-grafana-api

Trying to add more stuff (like teams for now and will try to do best) 

Use example : 
```
import "github.com/tcheronneau/gapi"

g,err := gapi.New('user:pass',"https://yourgrafana.tld") // OR USE TOKEN
users,err := g.GetUsers()
user,err := g.GetUser("test")
```
