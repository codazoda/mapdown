# Mapdown

Create a `sitemap.xml` file based on a list of markdown files in the current directory.


## Installing

Download the appropriate executable for your architecture and save it somewhere in your path.

[Download for MacOS 64-Bit Apple Silicon](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/mapdown-darwin-arm64)  
[Download for Linux x86 64-Bit](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/mapdown-linux-amd64)
[Download Windows 64-Bit](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/mapdown-linux-amd64)  
[Download Others]()  


## Usage

    mapdown https://example.com


## Example Output

    <?xml version="1.0" encoding="UTF-8"?>
    <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://www.example.com</loc>
        <lastmod>2024-12-16</lastmod>
        <changefreq>monthly</changefreq>
        <priority>1.0</priority>
    </url>
    <url>
        <loc>https://www.example.com/another-great-example</loc>
        <lastmod>2024-12-16</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    <url>
        <loc>https://www.example.com/example-article-about-topic</loc>
        <lastmod>2024-12-16</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    <url>
        <loc>https://www.example.com/index</loc>
        <lastmod>2024-12-16</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    </urlset>


## Building

    cd src
    make


## About

Mapdown was created by Joel Dare on December 16, 2024.
