## Sample for Authontification


```golang

package main 
import ( 
    "fmt" "html/template" 
    "net/http" "strconv" 
    "renrenwaigua/model" 
    "renrenwaigua/model/status" 
    "renrenwaigua/model/user" 
    "renrenwaigua/oauth" ) 

var ( 
       _code = "" 
       _token *oauth.Token 
   ) 
   var oauthCfg = &oauth.Config{ClientId: "2a7e17af6b2e42e2b0682551735602df", 
                                ClientSecret: "7dbc044f03604aada969e095d7b8e670", 
                                AuthURL: "https://graph.renren.com/oauth/authorize", 
                                TokenURL: "https://graph.renren.com/oauth/token", 
                                RedirectURL: "http://localhost:8080/oauth2callback", 
                                Scope: "read_user_status", 
                                TokenCache: oauth.CacheFile("cache.json"), } 
 type run map[int64]chan bool 
 var running run 
 func (r run) IsHave(id int64) bool { 
      if _, ok := r[id]; ok { 
         return true 
       } 
  return false 
  } 
  
  func (r run) Add(id int64) (chan bool, bool) { 
       c := make(chan bool) if r.IsHave(id) { 
           return c, false } 
       else { 
           r[id] = c return c, true } 
   } 
   
  func (r run) Del(id int64) bool { 
     if r.IsHave(id) { 
        delete(r, id) 
        return true 
      } 
      return false 
 } 
 
 func init() { } 
 
 func getTransport(w http.ResponseWriter, r *http.Request) (*oauth.Transport, error) { 
 t := &oauth.Transport{Config: oauthCfg} 
 var err error 
 if _token == nil { 
    _token, err = oauthCfg.TokenCache.Token() 
    
     if err != nil { return nil, err } } 
     if _token.Expired() { 
         err = t.Refresh() 
     if err != nil { 
        return nil, err } 
      } 
     t.Token = _token return t, nil 
     
     } 
     
 func handleRoot(w http.ResponseWriter, r *http.Request) { 
     t, err := getTransport(w, r) 
     if err != nil { 
       fmt.Println(err.Error()) 
       fmt.Println("from / to /authorize") 
       http.Redirect(w, r, "/authorize", http.StatusFound) return } 
       info := user.UserGet(t, _token.User.Id).GetInfo(user.TYPE_HEAD) 
       userGetTemplate, err := template.ParseFiles("static/auto.html") 
       
       if err != nil { 
          fmt.Println(err.Error()) 
       } 
       err = userGetTemplate.Execute(w, info) 
       
       if err != nil { 
         fmt.Println(err.Error()) 
        } 
        } 
        
      func handleStatus(w http.ResponseWriter, r *http.Request) { 
          t, err := getTransport(w, r) 
          if err != nil { 
             fmt.Println(err.Error()) 
             fmt.Println("from / to /authorize") 
             http.Redirect(w, r, "/authorize", http.StatusFound) 
             return 
           } 
           
           sta, err := status.GetList(t, t.Token.User.Id, 20, 1) 
           if err != nil { 
              fmt.Println("handle status error:", err.Error()) 
            } 
            
            statusListTemplate, err := template.ParseFiles("static/status.html") 
            if err != nil { 
               fmt.Println(err.Error()) 
             } 
             
             err = statusListTemplate.Execute(w, sta) 
             if err != nil { 
                fmt.Println(err.Error()) } 
             } 
             
            func handleAutoSofa(w http.ResponseWriter, r *http.Request) { 
                 r.ParseForm() 
                 t, err := getTransport(w, r) 
                 if err != nil { 
                    fmt.Println(err.Error()) 
                    fmt.Println("from /autosofa to /authorize") 
                    http.Redirect(w, r, "/authorize", http.StatusFound) 
                    return 
                  } 
                  
                  str := r.Form.Get("id") 
                  if str != "" { 
                     id, err := strconv.ParseInt(str, 10, 64) 
                     if err != nil { 
                        fmt.Println("pares int:", err.Error()) 
                      } 
                  if running.IsHave(id) { 
                     fmt.Println("already running") 
                   } else if 
                     c, b := running.Add(id); b { 
                       go model.AutoSofa(*t, id, c) 
                      } 
                   } 
                   
                   var ids struct { Id []int64 } 
                   for k, _ := range running { 
                       ids.Id = append(ids.Id, k) 
                   } 
                   temp, err := template.ParseFiles("static/sofa.html") 
                   if err != nil { 
                      fmt.Println("parse error:", err.Error()) 
                      return 
                   } 
                   err = temp.Execute(w, ids) 
                   if err != nil { 
                      fmt.Println("execute error:", err.Error()) 
                      return 
                    } 
                   } 
                   
                   func handleDel(w http.ResponseWriter, r *http.Request) { 
                       r.ParseForm() 
                       str := r.Form.Get("id") 
                       if str != "" { 
                          id, err := strconv.ParseInt(str, 10, 64) 
                          if err != nil { 
                             fmt.Println("pares int:", err.Error()) 
                             return 
                           } 
                           
                           if running.IsHave(id) { 
                              running[id] <- false running.Del(id) 
                            } } 
                            
                            http.Redirect(w, r, "/autosofa", http.StatusFound) 
                 } 
                 
                 func handleAuthorize(w http.ResponseWriter, r *http.Request) { 
                      url := oauthCfg.AuthCodeURL("") 
                      fmt.Println("from /authorize to ", url) 
                      http.Redirect(w, r, url, http.StatusFound) 
                 } 
                 
                 func handleOAuth2CallBack(w http.ResponseWriter, r *http.Request) { 
                      code := r.FormValue("code") 
                      if code == "" { 
                         fmt.Println("from /oauthecallback to /authorize") 
                         http.Redirect(w, r, "/authorize", http.StatusFound) 
                         return 
                        } 
                        
                        t := &oauth.Transport{Config: oauthCfg} 
                        _, err := t.Exchange(code) 
                        
                        if err != nil { 
                           fmt.Print(err.Error()) 
                           } 
                        _token = t.Token fmt.Println("from /oauthcallback to ", r.URL.RequestURI()) 
                        
                        http.Redirect(w, r, r.URL.RequestURI(), http.StatusFound) 
                        } 
                        
 func main() { 
                      running = make(map[int64]chan bool) 
                      http.HandleFunc("/", handleRoot) 
                      http.HandleFunc("/authorize", handleAuthorize) 
                      http.HandleFunc("/oauth2callback", handleOAuth2CallBack) 
                      http.HandleFunc("/status", handleStatus) 
                      http.HandleFunc("/autosofa", handleAutoSofa) 
                      http.HandleFunc("/del", handleDel) 
                      http.Handle("/bootstrap/", http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("bootstrap")))) 
                      fmt.Println("try listen on localhost:8080 .....") 
                      err := http.ListenAndServe(":8080", nil) 
                      
                      if err != nil { 
                         fmt.Println(err.Error()) return 
                      } 
 }
                                         
                      
                      
                      
                      
