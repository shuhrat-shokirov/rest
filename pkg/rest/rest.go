package rest

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"io/ioutil"
	"net/http"
)

type filePath struct {
	Path string `json:"path"`
	FileName string `json:"fileName"`
}

func ReadJSONBody(request *http.Request, dto interface{}) (err error) {
	if request.Header.Get("Content-Type") != "application/json" {
		return errors.New("error")
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New("error")
	}
	defer func() {
		err := request.Body.Close()
		if err != nil {
			return
		}
	}()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return errors.New("error")
	}
	return nil
}

func WriteJSONBody(response http.ResponseWriter, dto interface{}) (err error) {
	response.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(dto)
	if err != nil {
		return errors.New("error")
	}

	_, err = response.Write(body)
	if err != nil {
		return errors.New("error")
	}

	return nil
}

func JsonFileUpload(path string) (encoded string, err error) {
	fileStruct := make([]filePath, 0)

	fileStruct = append(fileStruct, filePath{
		Path:     "media/",
		FileName: path,
	})

	marshal, err := json.Marshal(fileStruct)
	if err != nil {
		return "", err
	}
	encoded = string(marshal)
	return encoded, nil
}

func dirFileReader(path string) (files []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}