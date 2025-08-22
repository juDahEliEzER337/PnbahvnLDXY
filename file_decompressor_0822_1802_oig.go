// 代码生成时间: 2025-08-22 18:02:31
package main

import (
    "archive/zip"
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// FileDecompressor 是一个结构体，用于处理文件解压
type FileDecompressor struct {
    SourceFile string // 源压缩文件路径
    DestinationDir string // 解压目标目录
}

// NewFileDecompressor 初始化一个新的 FileDecompressor 实例
func NewFileDecompressor(source, destination string) *FileDecompressor {
    return &FileDecompressor{
        SourceFile: source,
        DestinationDir: destination,
    }
}

// Decompress 解压指定的压缩文件到目标目录
func (f *FileDecompressor) Decompress() error {
    // 打开压缩文件
    srcFile, err := os.Open(f.SourceFile)
    if err != nil {
        return fmt.Errorf("unable to open source file: %w", err)
    }
    defer srcFile.Close()

    // 打开目标目录
    dstDir, err := os.Stat(f.DestinationDir)
    if err != nil || !dstDir.IsDir() {
        return fmt.Errorf("destination path is not a directory: %w", err)
    }

    // 解压文件
    zipReader, err := zip.OpenReader(f.SourceFile)
    if err != nil {
        return fmt.Errorf("failed to open zip file: %w", err)
    }
    defer zipReader.Close()

    for _, file := range zipReader.File {
        // 创建目标文件路径
        filePath := filepath.Join(f.DestinationDir, file.Name)
        // 确保路径是合法的，避免路径穿越攻击
        if !strings.HasPrefix(filePath, filepath.Clean(f.DestinationDir)+string(os.PathSeparator)) {
            return fmt.Errorf("illegal file path: %s", filePath)
        }

        // 创建文件
        if file.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
        } else {
            fileReader, err := file.Open()
            if err != nil {
                return fmt.Errorf("failed to open file in zip: %w", err)
            }
            defer fileReader.Close()

            dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return fmt.Errorf("failed to create file: %w", err)
            }
            defer dstFile.Close()

            _, err = io.Copy(dstFile, fileReader)
            if err != nil {
                return fmt.Errorf("failed to copy file: %w", err)
            }
        }
    }

    return nil
}

func main() {
    // 解析命令行参数
    source := flag.String("source", "", "path to the zip file to decompress")
    destination := flag.String("destination", "", "destination directory for the decompressed files")
    flag.Parse()

    if *source == "" || *destination == "" {
        log.Fatal("source and destination must be provided")
    }

    decompressor := NewFileDecompressor(*source, *destination)
    if err := decompressor.Decompress(); err != nil {
        log.Fatalf("failed to decompress files: %s", err)
    }
    fmt.Println("Decompression completed successfully")
}