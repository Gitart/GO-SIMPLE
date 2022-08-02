# fcgi.nginx

location /fastcgi_hello {
    # host and port to fastcgi server
    include         fastcgi.conf;
    fastcgi_pass 172.17.0.89:5000;
}

Client Request ----> Nginx (Reverse-Proxy) ----> App. FastCGI Server I. 127.0.0.1:5000

либо с балансировкой на несколько серверов:


	

# Nginx
upstream myapp1 {
    server 127.0.0.1:5000;
    server 127.0.0.1:5001;
    server 127.0.0.1:5002;
}

server {
    listen 80;

    location /some/path {
        fastcgi_pass http://myapp1;
    }
}

Client Request ----> Nginx (Reverse-Proxy)
                        |
                       /|\
                      | | `-> App. FastCGI Server I.   127.0.0.1:5000
                      |  `--> App. FastCGI Server II.  127.0.0.1:5001
                       `----> App. FastCGI Server III. 127.0.0.1:5002

