// 代码生成时间: 2025-08-04 15:15:30
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Renamer 定义一个结构体，用于批量重命名文件
type Renamer struct {
    srcDir  string // 源目录
    destDir string // 目标目录
    prefix  string // 新文件名前缀
}

// NewRenamer 创建并返回一个Renamer实例
func NewRenamer(srcDir, destDir, prefix string) *Renamer {
    return &Renamer{
        srcDir:  srcDir,
        destDir: destDir,
        prefix:  prefix,
    }
}

// Rename 执行批量文件重命名操作
func (r *Renamer) Rename() error {
    // 读取源目录下的所有文件
    files, err := os.ReadDir(r.srcDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            // 忽略子目录
            continue
        }

        // 构建新的文件名
        srcPath := filepath.Join(r.srcDir, file.Name())
        destPath := filepath.Join(r.destDir, r.prefix+file.Name())

        // 重命名文件
        if err := os.Rename(srcPath, destPath); err != nil {
            return fmt.Errorf("failed to rename file %s: %w", file.Name(), err)
        }
    }

    return nil
}

func main() {
    // 示例用法
    srcDir := "./src"
    destDir := "./dest"
    prefix := "new_"

    // 创建Renamer实例
    renamer := NewRenamer(srcDir, destDir, prefix)

    // 执行批量重命名
    if err := renamer.Rename(); err != nil {
        log.Fatalf("error occurred: %v", err)
    } else {
        fmt.Println("Bulk file renaming completed successfully")
    }
}