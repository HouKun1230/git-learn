package main

import "fmt"

func main() {

	fmt.Println("Hello, 世界")

	pdfg, _ := wkhtml.NewPDFGenerator()

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	_ = pdfg.Create()

	_ = pdfg.WriteFile("./Your_pdfname.pdf")

}
