// 代码生成时间: 2025-08-09 09:20:06
This program is designed to be extensible and maintainable, following Go best practices.
*/

package main

import (
    "archive/zip"
    "flag"
    "fmt"
    "io"
    "io/fs"
    "log"
    "os"
    "path/filepath"
)

// Constants for file paths
const (
    defaultDestinationDir = "./extracted"
)

// Unzip unzips a zip archive to a specified destination directory.
func Unzip(zipFilePath, destinationDir string) error {
    // Open the zip archive
    r, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return fmt.Errorf("failed to open zip file: %w", err)
    }
    defer r.Close()

    // Ensure the destination directory exists
    if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
        if err := os.MkdirAll(destinationDir, 0755); err != nil {
            return fmt.Errorf("failed to create destination directory: %w", err)
        }
    }

    // Iterate through the files in the zip archive
    for _, f := range r.File {
        filePath := filepath.Join(destinationDir, f.Name)
        if f.FileInfo().IsDir() {
            // Create directory
            if err := os.MkdirAll(filePath, 0755); err != nil {
                return fmt.Errorf("failed to create directory: %s, %w", filePath, err)
            }
            continue
        }

        // Create the file
        if err := func() error {
            file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return fmt.Errorf("failed to open file: %s, %w", filePath, err)
            }
            defer file.Close()

            // Extract the file content
            fileReader, err := f.Open()
            if err != nil {
                return fmt.Errorf("failed to open file inside zip: %s, %w", f.Name, err)
            }
            defer fileReader.Close()

            if _, err := io.Copy(file, fileReader); err != nil {
                return fmt.Errorf("failed to copy file content: %s, %w", filePath, err)
            }

            return nil
        }(); err != nil {
            return err
        }
    }

    return nil
}

func main() {
    var zipFilePath string
    var destinationDir string

    flag.StringVar(&zipFilePath, "zip", "", "Path to the zip file")
    flag.StringVar(&destinationDir, "dest", defaultDestinationDir, "Destination directory for extraction")
    flag.Parse()

    if zipFilePath == "" {
        log.Fatalf("zip file path must be provided")
    }

    if err := Unzip(zipFilePath, destinationDir); err != nil {
        log.Fatalf("failed to unzip file: %v", err)
    }

    fmt.Println("Extraction successful")
}