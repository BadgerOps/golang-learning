# ghost-converter

This script converts `.json` export from Ghost to Hugo markdown format. I wrote this to help me convert from my ghost blog, to github pages using hugo.

#### Ghost Export to Hugo Content Converter

Usage:
  -f string
        Path to the Ghost JSON export.
  -o string
        Output directory for the Hugo content.

Example:
  go run main.go -f export.json -o contentDir

This tool takes a Ghost CMS JSON export and converts each post into Hugo-compatible content. The resulting files will be saved in the provided output directory.`
