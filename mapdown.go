package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type URL struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func main() {
	// Parse command line arguments
	baseURL := flag.String("url", "", "Base URL for the sitemap (required)")
	outputFile := flag.String("file", "sitemap.xml", "Output file for the sitemap (optional)")
	flag.Parse()

	if *baseURL == "" {
		fmt.Println("Error: url is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Create the sitemap
	urlSet := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	// Add the home page
	urlSet.URLs = append(urlSet.URLs, URL{
		Loc:        *baseURL,
		LastMod:    time.Now().Format("2006-01-02"),
		ChangeFreq: "monthly",
		Priority:   "1.0",
	})

	// Add entries for all .md files in the current directory
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(info.Name()) == ".md" && info.Name() != "README.md" {
			filename := info.Name()[:len(info.Name())-3] // Remove .md extension
			lastMod := info.ModTime().Format("2006-01-02")
			urlSet.URLs = append(urlSet.URLs, URL{
				Loc:        fmt.Sprintf("%s/%s", *baseURL, filename),
				LastMod:    lastMod,
				ChangeFreq: "monthly",
				Priority:   "0.8",
			})
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading files: %v\n", err)
		os.Exit(1)
	}

	// Open the output file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Setup the encoder
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")

	// Add the XML declaration to the encoder
	encoder.EncodeToken(xml.ProcInst{
		Target: "xml",
		Inst:   []byte(`version="1.0" encoding="UTF-8"`),
	})
	encoder.EncodeToken(xml.CharData("\n"))
	encoder.Flush()

	// Write the XML content using the encoder
	if err := encoder.Encode(urlSet); err != nil {
		fmt.Printf("Error writing sitemap: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sitemap generated: %s\n", *outputFile)
}
