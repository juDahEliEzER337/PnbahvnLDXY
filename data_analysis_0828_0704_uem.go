// 代码生成时间: 2025-08-28 07:04:27
package main

import (
    "fmt"
    "gorm.io/driver/sqlite" // 这里以SQLite为例，实际项目中可以根据需要更换数据库驱动
    "gorm.io/gorm"
)

// DataRecord 存储数据记录的结构体
type DataRecord struct {
    gorm.Model
    Value float64 `gorm:"type:real;not null"` // 存储数据值
}

// AnalysisResults 存储分析结果的结构体
type AnalysisResults struct {
    Mean           float64
    Median         float64
    Variance       float64
    StandardDeviation float64
    Count         int64
}

// AnalysisService 提供数据分析的服务
type AnalysisService struct {
    db *gorm.DB
}

// NewAnalysisService 初始化AnalysisService
func NewAnalysisService(db *gorm.DB) *AnalysisService {
    return &AnalysisService{db: db}
}

// CalculateStatistics 计算并返回分析结果
func (service *AnalysisService) CalculateStatistics() (*AnalysisResults, error) {
    var dataRecords []DataRecord
    // 查询所有数据记录
    if err := service.db.Find(&dataRecords).Error; err != nil {
        return nil, err
    }

    if len(dataRecords) == 0 {
        return nil, fmt.Errorf("no data records found")
    }

    results := &AnalysisResults{Count: int64(len(dataRecords))}
    total := 0.0
    sortedValues := make([]float64, len(dataRecords))
    for _, record := range dataRecords {
        total += record.Value
        sortedValues = append(sortedValues, record.Value)
    }

    // 计算平均值
    results.Mean = total / float64(len(dataRecords))

    // 排序并计算中位数
    sort.Float64s(sortedValues)
    mid := len(sortedValues) / 2
    if len(sortedValues)%2 == 0 {
        results.Median = (sortedValues[mid-1] + sortedValues[mid]) / 2
    } else {
        results.Median = sortedValues[mid]
    }

    // 计算方差和标准差
    for _, value := range sortedValues {
        diff := value - results.Mean
        results.Variance += diff * diff
    }
    results.Variance /= float64(len(sortedValues))
    results.StandardDeviation = MathSqrt(results.Variance)

    return results, nil
}

// MathSqrt 计算平方根，由于Go的math库中没有直接提供，这里简单实现一个
func MathSqrt(x float64) float64 {
    // 这里使用一个简单的迭代方法来计算平方根，实际项目中可以使用math.Sqrt或者更精确的算法
    tolerance := 1e-10
    guess := x / 2.0
    for {
        nextGuess := (guess + x/guess) / 2.0
        if math.Abs(nextGuess-guess) < tolerance {
            break
        }
        guess = nextGuess
    }
    return guess
}

func main() {
    // 连接数据库，这里以SQLite为例，实际项目中可以根据需要更换
    db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 迁移模式
    db.AutoMigrate(&DataRecord{})

    // 创建AnalysisService
    service := NewAnalysisService(db)

    // 计算统计数据
    results, err := service.CalculateStatistics()
    if err != nil {
        fmt.Printf("Error calculating statistics: %v
", err)
        return
    }
    fmt.Printf("Statistics: Mean=%v, Median=%v, Variance=%v, StandardDeviation=%v, Count=%v
", results.Mean, results.Median, results.Variance, results.StandardDeviation, results.Count)
}