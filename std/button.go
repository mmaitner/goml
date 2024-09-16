package goml

import (
	"strings"
)

func Button(args Tags, comp ...Element) (r Element) {
	r.Id = "button_" + randSeq(5)
	decId := strings.Split(r.Id, "_")[1]
	concat := ""
	children := ""
	for _, e := range comp {
		concat += decId + ".appendChild(" + e.Id + "())\n"
		for _, subBody := range e.Body {
			children += subBody + "\n"
		}
	}
	dec := "let " + r.Id + " = () => { const " + decId + " = document.createElement('button');\n"

	r.Body = append(r.Body, children, dec+"\n"+concat+";\nreturn "+decId+"};")
	return r

}
