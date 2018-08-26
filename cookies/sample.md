# Cookies

Set cookies in Go
Go uses	the SetCookie function	in the net/http	package to set cookies:

http.SetCookie(w ResponseWriter, cookie	*Cookie)
w is the response of the request and cookie	is a struct. Let's see what	it looks like:

type Cookie struct{
				Name			string
				Value			string
				Path			string
				Domain		    string
				Expires			time.Time
				RawExpires	    string
//	MaxAge=0	means	no	'Max-Age'	attribute	specified.
//	MaxAge<0	means	delete	cookie	now,	equivalently	'Max-Age:	0'
//	MaxAge>0	means	Max-Age	attribute	present	and	given	in	seconds
				MaxAge			int
				Secure			bool
				HttpOnly	        bool
				Raw				string
				Unparsed	    []string	//	Raw	text	of	unparsed	attribute-value	pairs
}

Here is	an example of setting a	cookie:
## Set cookies

```golang
expiration:=time.Now().Add(365*24*time.Hour)
cookie:=http.Cookie{Name:"username",Value:"pagenum",Expires:expiration}
http.SetCookie(w,&cookie)
```

## Fetch cookies in Go
The above example shows	how to set a cookie. Now let's	see how	to get a cookie	that has been set:

```golang
cookie,	_:=r.Cookie("username")
fmt.Fprint(w,cookie)
```

### Here is another way to get a cookie:

```golang
for_,cookie:=range r.Cookies()	{
   fmt.Fprint(w, cookie.Name)
}
```


