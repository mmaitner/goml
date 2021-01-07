package goml

import "strings"

func InputField(args Tags, comp string) (r Element) {
	r.Id = "inputField_" + randSeq(5)

	dec := "let " + r.Id + " = () => { const " + strings.Split(r.Id, "_")[1] + " = document.createElement('input');"
	argStr := genArgString(args, r.Id) + strings.Split(r.Id, "_")[1] + ".placeholder = '" + comp + "';\n" + strings.Split(r.Id, "_")[1] + ".type = 'text';"
	final := dec + argStr + "\nreturn " + strings.Split(r.Id, "_")[1] + "}"
	r.Body = []string{}
	r.Body = append(r.Body, final)
	return r
}