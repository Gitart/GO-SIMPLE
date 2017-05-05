
## RUN server



### Install and Configure Nginx
Nginx is a high performant web server, load balancer, and well suits to deploy the high traffic websites. Even though this decision is opinionated, Python and Node developers usually use this. On Ubuntu 16.04, use these
```
sudo apt-get update
sudo apt-get install nginx
```

On Mac, the default Nginx listening port will be 8000. So modify the port to 80 by editing this file.
```
sudo vi /usr/local/etc/nginx/nginx.conf
```
Then find this code section and modify 8000 port to 80.
```
server {
        listen       8080; # Change this to 80 
        server_name  localhost;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            root   html;
            index  index.html index.htm;
        }
        
        ...       
}
```

Now everything is ready. This basic server serves static files from a directory called html. root can be modified to any directory that is comfortable to use.
You can check the status of Nginx with command
```
systemctl status nginx
```


### Monitoring Install supervisord
Monitoring your Go web server with Supervisord   
https://serverfault.com/questions/96499/how-to-automatically-start-supervisord-on-linux-ubuntu


#### Installing supervisord

We can easily install supervisord using Python’s pip command. On Ubuntu 16.04, just use the apt-get command.
sudo apt-get install -y supervisor
This installs two tools, supervisor and supervisorctl. Supervisorctl is intended to control the supervisord and add tasks, restart tasks etc.
Now create a configuration file at
```
/etc/supervisor/conf.d/goproject.conf
```

You can add any number of configuration files and supervisor treats them as separate processes to run. Add this content to the above file.
```
[supervisord]
logfile = /tmp/supervisord.log
[program:myserver]
command=/home/naren/golab/bin/myserver
autostart=true
autorestart=true
redirect_stderr=true
By default, we have a file called supervisord.conf at /etc/supervisor/. Look at it for more reference.
[supervisord] section tells the location of logfile for supervisord.
[program:myserver] is the task block which traverses to given directory and executes the command given.
Now we can ask our supervisorctl to re read the configuration and restart the tasks(process). For that just say
supervisorctl reread
supervisorctl update
Then launch our supervisorctl with
supervisorctl
```




### GO Srever

```golang
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		log.Printf("%q", r.UserAgent())
	})
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
```

Now if we need to run this project, use these commands.
go install github.com/narenaryan/myserver

This copies the main executable in the folder golab/bin directory. If you have this path included in the $PATH variable, you can run project executable as a normal command. For now, we can run our project as
```
sudo $GOPATH/bin/myserver
```

Our Go application server is running on port 8000. We now need to proxy it through Nginx.

As we previously discussed, we need to edit the default sites-available server block called as default
vi /etc/nginx/sites-available/default
and modify the location section to this.

```
server {
        listen 80 default_server;
        listen [::]:80 default_server;
        root /var/www/html;
        # Add index.php to the list if you are using PHP
        index index.html index.htm index.nginx-debian.html;
       server_name _;
       location / {
           # First attempt to serve request as file, then
           # as directory, then fall back to displaying a 404.
           proxy_pass http://127.0.0.1:8000;
           try_files $uri $uri/ =404;
       }
}
```


Now if you have a microservice based architecture and need to forward various URL requests to different service, we can do that using server blocks like this.
```
server {
    listen       ...;
    ...
    location / {
        proxy_pass http://127.0.0.1:8000;
    }
    
    location /api {
        proxy_pass http://127.0.0.1:8001;
    }

    location /mail {
        proxy_pass http://127.0.0.1:8002;
    }
    ...
}
```



