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

	if err := zipFiles(output, files); err != nil {
		log.Fatalf("cannot zip file: %v", err)
	}
}

func zipFiles(output string, files []string) error {
	newZipFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, filename := range files {
		if err = AddFileToZip(zipWriter, filename); err != nil {
			return err
		}
	}

	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	info, err := inputFile.Stat()
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

	_, err = io.Copy(writer, inputFile)
	return err
}
