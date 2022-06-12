package handler

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func writeFile(file multipart.File, fileHeader *multipart.FileHeader) error {
	buff := make([]byte, 512)

	_, err := file.Read(buff)
	if err != nil {
		return err
	}

	pathUp := filepath.Join(ImgDirectory, fileHeader.Filename)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	f, err := os.Create(pathUp)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}
