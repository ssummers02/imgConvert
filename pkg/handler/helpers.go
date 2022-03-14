package handler

import (
	"imgConverter/pkg/service"
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

	/*	detectedFileType := http.DetectContentType(buff)
		switch detectedFileType {
		case "image/jpeg":
		case "image/jpg":
		case "image/svg":
		case "image/png":
			break
		default:
			SendErr(w, http.StatusBadRequest, "The provided file format is not allowed. Please upload a JPEG, SVG or PNG image")
			return
		}*/

	pathUp := filepath.Join(service.ImgDirectory, fileHeader.Filename)

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
