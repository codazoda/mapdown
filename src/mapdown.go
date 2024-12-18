package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	// Define flags
	outputFile := flag.String("file", "sitemap.xml", "Output file for the sitemap (optional)")

	// Customize the usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS] <baseURL>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\n<baseURL> must be a fully qualified URL (e.g., https://www.example.com)\n\nOptions:\n")
		flag.PrintDefaults()
	}

	// Parse flags
	flag.Parse()

	// Retrieve positional arguments
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Error: baseURL is required.")
		flag.Usage()
		os.Exit(1)
	}

	baseURL := args[0]

	// Validate baseURL
	parsedURL, err := url.ParseRequestURI(baseURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		fmt.Printf("Error: Invalid baseURL provided. Please provide a fully qualified URL (e.g., https://www.example.com)\n")
		os.Exit(1)
	}

	// Ensure baseURL does not have a trailing slash
	baseURL = strings.TrimRight(parsedURL.String(), "/")

	// Create the sitemap
	urlSet := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	// Add the home page without trailing slash
	urlSet.URLs = append(urlSet.URLs, URL{
		Loc:        baseURL,
		LastMod:    time.Now().Format("2006-01-02"),
		ChangeFreq: "monthly",
		Priority:   "1.0",
	})

	// Add entries for .md files in the current directory and subdirectories
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories; we're only interested in files
		if info.IsDir() {
			return nil
		}

		// Process only .md files excluding README.md
		if filepath.Ext(info.Name()) != ".md" || strings.ToLower(info.Name()) == "readme.md" {
			return nil
		}

		// Get the relative path
		relPath, err := filepath.Rel(".", path)
		if err != nil {
			return err
		}

		// Convert OS-specific path separators to URL separators
		urlPath := strings.ReplaceAll(relPath, string(os.PathSeparator), "/")

		var loc string
		if strings.ToLower(filepath.Base(urlPath)) == "index.md" {
			// If the file is index.md, remove 'index.md' from the path
			dirPath := filepath.Dir(urlPath)
			if dirPath == "." {
				// If index.md is in the root directory, skip to avoid duplication
				return nil
			}
			loc = fmt.Sprintf("%s/%s/", baseURL, strings.TrimRight(dirPath, "/"))
		} else {
			// Remove the .md extension for other Markdown files
			trimmedPath := strings.TrimSuffix(urlPath, ".md")
			loc = fmt.Sprintf("%s/%s", baseURL, trimmedPath)
		}

		// Add the URL entry
		urlSet.URLs = append(urlSet.URLs, URL{
			Loc:        loc,
			LastMod:    info.ModTime().Format("2006-01-02"),
			ChangeFreq: "monthly",
			Priority:   "0.8",
		})

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

	// Write the XML header
	file.WriteString(xml.Header)

	// Write the XML content using the encoder
	if err := encoder.Encode(urlSet); err != nil {
		fmt.Printf("Error writing sitemap: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sitemap generated: %s\n", *outputFile)
}
