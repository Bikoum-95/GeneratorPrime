package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func CreatePdf(p PdfSaver, pi Pirate) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.Image("img/wantedVierge.jpg", 0, 0, 210, 295, false, "", 0, "")

	pdf.SetFont("Arial", "B", 50)
	pdf.Text(35, 230, fmt.Sprintf(pi.Name))

	pdf.SetFont("Arial", "B", 50)
	pdf.Text(45, 259, fmt.Sprintf(pi.Prime))
	pdf.Image(pi.Img, 20, 62, 170, 125, false, "", 0, "")

	pdf.OutputFileAndClose(pi.Name + ".pdf")
}
