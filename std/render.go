package goml

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type RenderParams struct {
	BodyName string
}
type renderBody struct {
	body string
	name string
}
type Product interface {
	Write()
}

func Render(el []Element, params RenderParams) renderBody {
	id := "body_body"
	decId := strings.Split(id, "_")[0]
	if params.BodyName != "" {
		id = params.BodyName
		decId = params.BodyName + "_body"
	}
	children := ""
	concat := ""
	for _, e := range el {
		concat += decId + ".appendChild(" + e.Id + "())\n"
		for _, subBody := range e.Body {
			children += subBody + "\n"
		}
	}
	children = strings.Replace(children, "\n", " ", -1)
	re := regexp.MustCompile(`let\s[\w_]+\s=\s\(\)\s=>\s\{(.+?)\};`)
	match := re.FindAllStringSubmatch(children, -1)
	matched := []string{}
	for _, e := range match {
		matched = append(matched, e[0])
	}
	children = strings.Join(removeDuplicateValues(matched), " ")
	dec := "let " + id + " = () => { const " + decId + " = document.createElement('div');\n"
	body := []string{children, dec + "\n" + concat + "\nreturn " + decId + "}"}
	return renderBody{body: strings.Join(removeDuplicateValues(body), "\n"), name: id}
}

func (r renderBody) Write(name ...string) {
	if len(name) == 0 || len(name) > 1 {
		name = append(name, "index")
	}
	template := "<!DOCTYPE html>\n" +
		"<html>\n" +
		"\t<head>\n" +
		"\t\t<meta charset=\"utf-8\">\n" +
		"\t\t<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n" +
		"\t\t<title>Goml app</title>\n" +
		"\t\t<meta name=\"description\" content=\"\">\n" +
		"\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n" +
		"\t\t<link rel=\"stylesheet\" href=\"css/style.css\">\n" +
		"\t\t<script src=\"" + name[0] + ".js\"></script>\n" +
		"\t</head>\n" +
		"\t<body>\n" +

		"\t\t<noscript>You need to enable JavaScript to run this app.</noscript>\n" +
		"\t\t<div id=\"root\"></div>\n" +
		"\t\t" +
		"\t\t<script>document.getElementById(\"root\").appendChild(" + r.name + "())</script>\n" +
		"\t</body>\n" +
		"</html>\n"

	ioutil.WriteFile("index.html", []byte(template), 0644)
	ioutil.WriteFile(name[0]+".js", []byte(r.body), 0644)
}
