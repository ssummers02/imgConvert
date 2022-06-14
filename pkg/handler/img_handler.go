package handler

import (
	"encoding/json"
	"fmt"
	"imgConverter/pkg/restmodel"
	"net/http"
	"os"
	"path/filepath"
)

const (
	MaxUploadSize = 32 << 20
	ImgDirectory  = "uploads"
)

func (s *Server) convertImg(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}
	defer file.Close()

	if fileHeader.Size > MaxUploadSize {
		SendErr(w, http.StatusBadRequest, fmt.Sprintf("The uploaded image is too big: %s.", fileHeader.Filename))

		return
	}

	err = writeFile(file, fileHeader)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, fmt.Sprintln("Error write file"))

		return
	}

	pathUploads := filepath.Join(ImgDirectory, fileHeader.Filename)

	defer os.Remove(pathUploads)

	opt := restmodel.ImgOptions{
		OutputFormat: filepath.Ext(pathUploads)[1:],
	}

	err = json.Unmarshal([]byte(r.FormValue("json")), &opt)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "error json")

		return
	}

	converting, err := s.services.Img.ImageProcessing(opt, pathUploads)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	defer os.Remove(converting)

	sendFile(w, r, converting)
}

func (s *Server) maxSize(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}
	defer file.Close()

	if fileHeader.Size > MaxUploadSize {
		SendErr(w, http.StatusBadRequest, fmt.Sprintf("The uploaded image is too big: %s.", fileHeader.Filename))

		return
	}

	err = writeFile(file, fileHeader)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, fmt.Sprintln("Error write file"))

		return
	}

	pathUploads := filepath.Join(ImgDirectory, fileHeader.Filename)

	defer os.Remove(pathUploads)

	opt := restmodel.ImgOptions{
		OutputFormat: filepath.Ext(pathUploads)[1:],
	}

	err = json.Unmarshal([]byte(r.FormValue("json")), &opt)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "error json")

		return
	}

	converting, err := s.services.Img.ImageResize(opt, pathUploads)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	defer os.Remove(converting)

	sendFile(w, r, converting)
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
