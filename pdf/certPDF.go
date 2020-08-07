package pdf

import (
	"fmt"
	"os"
	"path"
	"udemy/genCert/cert"

	"github.com/jung-kurt/gofpdf"
)

// PdfSaver is the data structure for a PDF
type PdfSaver struct {
	OutputDir string
}

// New generate a new PDF
func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputDir,
	}

	return p, nil
}

// Save a PDF
func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Printf("Saved certificate to %v\n", path)
	}
	return nil
}
