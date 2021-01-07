package goml

import "strings"

func Header(args Tags, comp ...Element) (r Element) {
	r.Id = "header_" + randSeq(5)
	children := ""
	concat := ""
	for _, e := range comp {
		concat += strings.Split(r.Id, "_")[1] + ".appendChild(" + e.Id + "())\n"
		for _, subBody := range e.Body {
			children += subBody + "\n"
		}
	}
	dec := "let " + r.Id + " = () => {const " + strings.Split(r.Id, "_")[1] + " = document.createElement('h1');\n"
	r.Body = append(r.Body, children, dec+"\n"+genArgString(args, r.Id)+concat+"\nreturn "+strings.Split(r.Id, "_")[1]+"}")

	return r
}
