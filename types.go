package main

type FileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

type Report struct {
	TotalFiles   int            `json:"totalFiles"`
	TotalSizeMB  float64        `json:"totalSizeMB"`
	FileTypes    map[string]int `json:"fileTypes"`
	LargestFiles []FileInfo     `json:"largestFiles"`
}
type FileResult struct {
	Path string
	Size int64
	Ext  string
}
