# Concurrent-File-Processinf-Tool
Overview
A Go-based CLI application that recursively scans a directory, processes files concurrently, and generates a JSON report containing:

Total files
File type distribution
Total size
Largest files
Progress tracking
Running the Application
go run . scan "<directory-path>"

Example:
go run . scan "C:\Users\user\Downloads"

Architecture
WalkDir
   |
   v
fileChan
   |
   +--> Worker 1
   +--> Worker 2
   +--> Worker N
            |
            v
       resultChan
            |
            v
       Aggregator
            |
            v
        JSON Report
Worker Pool

The application uses the Worker Pool pattern for concurrent file processing.

Workers are created using:

workers := runtime.NumCPU()

Each worker receives file paths from fileChan, processes the file, and sends results to resultChan.

Benefits:

Concurrent file processing
Better CPU utilization
Scalable architecture
Channels
fileChan

Used to send file paths from the directory scanner to workers.

fileChan <- path

Producer:

WalkDir

Consumers:

Worker goroutines
resultChan

Used to send processed file information from workers to the aggregator.

resultChan <- FileResult{}

Producer:
Worker goroutines

Consumer:
Aggregator loop

Benefits:
Safe communication between goroutines
Avoids shared-memory complexity
WaitGroup
Used to track worker completion.

var wg sync.WaitGroup

Before starting a worker:

wg.Add(1)

Inside worker:

defer wg.Done()

Wait for all workers:

wg.Wait()

Purpose:

Ensures all workers finish before closing channels
Prevents incomplete results
Atomic Counter
Used for progress tracking.
atomic.AddInt64(processed, 1)

Read progress:

atomic.LoadInt64(&processed)

Purpose:

Thread-safe counter updates
Avoids race conditions when multiple workers update progress simultaneously
Report Generation

The application generates a JSON report using:

json.MarshalIndent()

Example output:

{
  "totalFiles": 31,
  "totalSizeMB": 241.47,
  "fileTypes": {
    ".pdf": 13,
    ".png": 4
  }
}
Key Go Concepts Used
Goroutines
Channels
Worker Pool Pattern
WaitGroup
Atomic Operations
Structs
Maps
JSON Serialization
Recursive File Traversal
Future Enhancements
Save report to file (report.json)
Configurable worker count
Directory exclusion support
Context cancellation (Ctrl+C)
Unit tests
Benchmarking
