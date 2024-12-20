# Mapdown

Create a `sitemap.xml` file based on the Markdown files in the current directory.


## Installing

Download the appropriate executable for your architecture and save it somewhere in your path.

[Download for MacOS 64-Bit Apple Silicon](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/darwin-arm64/mapdown)

[Download for Linux x86 64-Bit](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/linux-amd64/mapdown)

[Download Windows x86 64-Bit](https://github.com/codazoda/mapdown/raw/refs/heads/main/bin/windows-amd64/mapdown.exe)

[Other Downloads](https://github.com/codazoda/mapdown/tree/main/bin)


## Usage

    mapdown https://example.com


## Example Output

    <?xml version="1.0" encoding="UTF-8"?>
    <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://www.example.com</loc>
        <lastmod>2024-12-18</lastmod>
        <changefreq>monthly</changefreq>
        <priority>1.0</priority>
    </url>
    <url>
        <loc>https://www.example.com/first-example</loc>
        <lastmod>2024-12-17</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    <url>
        <loc>https://www.example.com/second-example</loc>
        <lastmod>2024-12-17</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    <url>
        <loc>https://www.example.com/special/another-example</loc>
        <lastmod>2024-12-17</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    </urlset>


## Building

    cd src
    make


## Notes

- The main URL is always included (https://www.example.com) even if there is not an index.md file
- Where index.md files exist the path without index is used (https://www.example.com/example-directory)

_Mapdown was created by Joel Dare on December 16, 2024._
