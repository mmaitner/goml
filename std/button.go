package goml

import (
	"strings"
)

func Button(args Tags, comp string) (r Element) {
	r.Id = "button_" + randSeq(5)
	decId := strings.Split(r.Id, "_")[1]
	children := ""
	concat := decId + ".innerText = '" + comp + "';"
	dec := "const " + r.Id + "= () => { const " + decId + " = document.createElement('button');\n"

	r.Body = append(r.Body, children, dec+"\n"+concat+";\nreturn "+decId+"}")
	return r

}
