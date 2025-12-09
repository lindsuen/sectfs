// sectfs - download_file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lindsuen/sectfs/internal/db"
	sectfs "github.com/lindsuen/sectfs/server/core"
)

type DwonloadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func DownloadFile(c echo.Context) error {
	downloadResponse := new(DwonloadResponse)
	downloadResponse.Success = true
	downloadResponse.Message = ""

	file := new(sectfs.File)
	fileId := c.QueryParam("fileid")

	encodeValue := string(db.Get([]byte(fileId)))
	if encodeValue == "" {
		downloadResponse.Success = false
		downloadResponse.Message = "The file is not found."
		log.Println("fileid: " + fileId + ". The file is not found.")
		return c.JSON(http.StatusOK, &downloadResponse)
	}

	value, _ := base64.RawURLEncoding.DecodeString(encodeValue)
	err := json.Unmarshal(value, &file)
	if err != nil {
		return err
	}

	if !fileIsExist(file.Path) {
		downloadResponse.Success = false
		downloadResponse.Message = "The file is not found."
		log.Println("fileid: " + fileId + ". The file is not found.")
		return c.JSON(http.StatusOK, &downloadResponse)
	}
	return c.Attachment(file.Path, file.Name)
}

func fileIsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}
