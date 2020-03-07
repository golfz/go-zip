package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Start zip files")

	files := []string{"input/file1.txt", "input/file2.txt"}
	output := "output/myzip.zip"

	newZipFile, err := os.Create(output)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			log.Fatalf("error: %v", err)
		}
	}

}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)

	log.Printf("add a file to zip: %s", filepath.Base(filename))

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}
