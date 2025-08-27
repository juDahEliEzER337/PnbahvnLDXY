// 代码生成时间: 2025-08-27 22:17:56
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler is a struct that wraps the cron scheduler
type Scheduler struct {
    c *cron.Cron
}

// NewScheduler creates a new instance of the scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        c: cron.New(cron.WithSeconds()), // Allows scheduling to the second
    }
}

// AddTask adds a task to the scheduler
func (s *Scheduler) AddTask(schedule string, task func()) error {
    _, err := s.c.AddFunc(schedule, task)
    if err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.c.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.c.Stop()
}

// ExampleTask is a sample task that can be scheduled
func ExampleTask() {
    fmt.Println("Task executed at:", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    scheduler := NewScheduler()
    defer scheduler.Stop()

    // Add a task to run every minute at 30 seconds past the hour
    if err := scheduler.AddTask("*/1 * * * * 30", ExampleTask); err != nil {
        log.Fatalf("Failed to add task: %v", err)
    }

    // Block the main thread to keep the scheduler running
    select {}
}
