package handler

import (
	"encoding/json"
	"fmt"
	"imgConverter/pkg/restmodel"
	"imgConverter/pkg/service"
	"net/http"
	"os"
	"path/filepath"
)

const MAX_UPLOAD_SIZE = 32 << 20

func (s *Server) convertImg(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	if fileHeader.Size > MAX_UPLOAD_SIZE {
		SendErr(w, http.StatusBadRequest, fmt.Sprintf("The uploaded image is too big: %s.", fileHeader.Filename))
		return
	}

	err = writeFile(file, fileHeader)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, fmt.Sprintf("Error write file"))
		return
	}
	pathUploads := filepath.Join(service.ImgDirectory, fileHeader.Filename)
	defer os.Remove(pathUploads)

	opt := restmodel.ImgOptions{
		Quality:      100,
		OutputFormat: filepath.Ext(pathUploads)[1:],
	}

	err = json.Unmarshal([]byte(r.FormValue("json")), &opt)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "json")
		return
	}

	converting, err := s.services.Img.ImageProcessing(opt, fileHeader.Filename)
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
