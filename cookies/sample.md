# Cookies

## Set cookies

```golang
expiration	:=	time.Now().Add(365	*	24	*	time.Hour)
cookie	:=	http.Cookie{Name:	"username",	Value:	"astaxie",	Expires:	expiration}
http.SetCookie(w,	&cookie)
```

## Fetch	cookies	in	Go
The	above	example	shows	how	to	set	a	cookie.	Now	let's	see	how	to	get	a	cookie	that	has	been	set:

```golang
cookie,	_	:=	r.Cookie("username")
fmt.Fprint(w,	cookie)
Here	is	another	way	to	get	a	cookie:
for	_,	cookie	:=	range	r.Cookies()	{
				fmt.Fprint(w,	cookie.Name)
}
```

As	you	can	see,	
