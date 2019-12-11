# How to check pointer or interface is nil?

```
package` `main`
`import` `(`
`"fmt"`
`)`
`type` `Temp` `struct` `{`
`}`
`func` `main() {`
`var` `pnt *Temp` `// pointer`
`var` `inf` `interface``{}` `// interface declaration`
`inf = pnt` `// inf is a non-nil interface holding a nil pointer (pnt)`
`fmt.Printf(``"pnt is a nil pointer: %v\n"``, pnt == nil)`
`fmt.Printf(``"inf is a nil interface: %v\n"``, inf == nil)`
`fmt.Printf(``"inf is a interface holding a nil pointer: %v\n"``, inf == (*Temp)(nil))`
`}`
```
