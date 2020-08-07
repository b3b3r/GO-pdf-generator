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

func breakLine(pdf *gofpdf.Fpdf) {
	pdf.Ln(30)
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions(
		"img/background.png",
		0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, cert *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	//first image
	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(
		filename,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0, "")

	//second image
	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(
		filename,
		x-margin, 20,
		imageWidth, 0,
		false, opts, 0, "")

	//tilte
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, cert.LabelCompletion, "C")

	breakLine(pdf)
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 20.0
	x := 0.0
	y := 0.0
	imageWidth := 50.0
	imageHeight := 40.0
	filename := "img/stamp.png"
	pageWidth, pageHeight := pdf.GetPageSize()
	x = pageWidth - imageWidth
	y = pageHeight - imageHeight
	pdf.ImageOptions(
		filename,
		x-margin, y-margin,
		imageWidth, 0,
		false, opts, 0, "")
}

func body(pdf *gofpdf.Fpdf, cert *cert.Cert) {
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	breakLine(pdf)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	breakLine(pdf)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	breakLine(pdf)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	breakLine(pdf)
}

// Save a PDF
func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//background
	background(pdf)

	//header
	header(pdf, &cert)

	//body
	body(pdf, &cert)

	//footer
	footer(pdf)

	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to %v\n", path)
	return nil
}
