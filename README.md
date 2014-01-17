
Html2md is an implementation of converting html to markdown for Go.

# Installation

If you have [gopm](https://github.com/gpmgo/gopm) installed, 

	gopm get github.com/lunny/html2md
	
Or

	go get github.com/lunny/html2md

# Usage

* Html2md has some html tag Rules. For simple use, just use

```Go
    md := html2md.Convert(html)
```

* If you want to use your self Rule. You can

```Go
   html2md.AddRule(&html2md.Rule{
       patterns: []string{"hr"},
	   tp:       Void,
	   replacement: func(innerHTML string, attrs []string) string {
			return "\n\n* * *\n"
		},
   })
```

or

```Go
html2md.AddConvert(func(conent string) string {
    return strings.ToLower(content)
})
```

# Documents

* [GoDoc](http://godoc.org/github.com/lunny/html2md)

* [GoWalker](http://gowalker.org/github.com/lunny/html2md)

# LICENSE

 BSD License
 [http://creativecommons.org/licenses/BSD/](http://creativecommons.org/licenses/BSD/)
