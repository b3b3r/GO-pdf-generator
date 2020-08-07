package html

import (
	"fmt"
	"os"
	"path"
	"text/template"
	"udemy/genCert/cert"
)

type HtmlSaver struct {
	OutputDir string
}

// New generate a new HTML
func New(outputDir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return h, err
	}
	h = &HtmlSaver{
		OutputDir: outputDir,
	}

	return h, nil
}

// Save a html file
func (h *HtmlSaver) Save(cert cert.Cert) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.html", cert.LabelTitle)
	path := path.Join(h.OutputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = t.Execute(f, cert)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to %v\n", path)
	return nil
}

var tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
        <title>{{.LabelTitle}}</title>
        <style>
            body {
                text-align: center;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
            }

            h1 {
                font-size: 3em;
            }

        </style>
    </head>
    <body>
        <h1>{{.LabelCompletion}}</h1>
        <h2><em>{{.LabelPresented}}</em></h2>
        <h1>{{.Name}}</h1>
        <h2>{{.LabelParticipation}}</h2>
        <p>
            <em>{{.LabelDate}}</em>
        </p>
    </body>
</html>`
