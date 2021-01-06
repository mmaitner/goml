package goml

import "strings"

func Text(args Tags, comp string) (r Element) {
	r.Id = "text_" + randSeq(5)
	dec := "let " + r.Id + " = () => { const " + strings.Split(r.Id, "_")[1] + " = document.createElement('p');"
	argStr := genArgString(args, r.Id) + strings.Split(r.Id, "_")[1] + ".innerHTML = '" + comp + "'"
	final := dec + argStr + "\nreturn " + strings.Split(r.Id, "_")[1] + "}"
	r.Body = []string{}
	r.Body = append(r.Body, final)
	return r
}
