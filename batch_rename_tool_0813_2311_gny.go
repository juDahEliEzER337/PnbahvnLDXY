// 代码生成时间: 2025-08-13 23:11:08
 * Features:
 * - Directory traversal for file discovery.
 * - Batch renaming functionality with error handling.
 * - Logging for operations.
 *
 * @author Your Name
 * @version 1.0.0
 */

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// File represents a file to be renamed.
type File struct {
    Path  string
    NewName string
}

// Renamer is a function that renames a file.
type Renamer func(oldPath, newPath string) error

// BatchRename renames multiple files.
func BatchRename(renamer Renamer, files []File) error {
    for _, file := range files {
        if err := renamer(file.Path, file.NewName); err != nil {
            return fmt.Errorf("error renaming file %s: %w", file.Path, err)
        }
        log.Printf("Renamed file from %s to %s", file.Path, file.NewName)
    }
    return nil
}

// RenameFile renames a file.
func RenameFile(oldPath, newPath string) error {
    if _, err := os.Stat(newPath); err == nil {
        return fmt.Errorf("file %s already exists", newPath)
    }
    return os.Rename(oldPath, newPath)
}

// FindFiles finds all files in a directory and its subdirectories.
func FindFiles(directory string) ([]File, error) {
    var files []File
    err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            baseName := filepath.Base(path)
            files = append(files, File{Path: path, NewName: strings.ReplaceAll(baseName, "space", "_")})
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return files, nil
}

func main() {
    directory := "." // The directory to scan for files.
    if len(os.Args) > 1 {
        directory = os.Args[1]
    }

    files, err := FindFiles(directory)
    if err != nil {
        log.Fatalf("error finding files: %s", err)
    }

    if err := BatchRename(RenameFile, files); err != nil {
        log.Fatalf("error renaming files: %s", err)
    }
}
