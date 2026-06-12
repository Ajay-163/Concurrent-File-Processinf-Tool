package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

func ScanDirectory(root string) Report {

	report := Report{
		FileTypes: make(map[string]int),
	}
	fileChan := make(chan string, 100)
	resultChan := make(chan FileResult, 100)

	var wg sync.WaitGroup
	workers := runtime.NumCPU()
	var processed int64
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return

			default:
				count := atomic.LoadInt64(&processed)

				fmt.Printf("Processed: %d\n", count)

				time.Sleep(time.Second)
			}
		}
	}()
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go worker(
			fileChan,
			resultChan,
			&wg,
			&processed,
		)
	}
	var allFiles []FileInfo

	var totalSize int64

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		fileChan <- path
		return nil
	})
	close(fileChan)

	go func() {
		wg.Wait()

		close(resultChan)
	}()

	for result := range resultChan {

		report.TotalFiles++

		totalSize += result.Size

		ext := result.Ext

		if ext == "" {
			ext = "no_extension"
		}

		report.FileTypes[ext]++

		allFiles = append(allFiles, FileInfo{
			Path: result.Path,
			Size: result.Size,
		})
	}
	report.TotalSizeMB =
		math.Round(
			(float64(totalSize)/(1024*1024))*100,
		) / 100
	sort.Slice(allFiles, func(i, j int) bool {
		return allFiles[i].Size > allFiles[j].Size
	})

	if len(allFiles) > 10 {
		report.LargestFiles = allFiles[:10]
	} else {
		report.LargestFiles = allFiles
	}
	close(done)
	return report

}
func worker(
	fileChan <-chan string,
	resultChan chan<- FileResult,
	wg *sync.WaitGroup,
	processed *int64,
) {

	defer func() {

		wg.Done()
	}()

	for path := range fileChan {

		info, err := os.Stat(path)
		if err != nil {
			continue
		}
		time.Sleep(300 * time.Millisecond)
		resultChan <- FileResult{
			Path: path,
			Size: info.Size(),
			Ext:  filepath.Ext(path),
		}
		atomic.AddInt64(processed, 1)
	}
}
