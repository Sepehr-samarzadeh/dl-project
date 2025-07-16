package downloader

import (
	"io"
	"net/http"
	"os"
)

type DlObject struct {
	File_name string
	File_size byte
	File_url  string
}

func (d *DlObject) SaveData(body_content *http.Response) error {
	d.File_name = "testfile0.0.0.1.mp3" //TODO: we need filetypes for .png,.jpeg,.txt and ...
	file, err := os.Create(d.File_name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body_content.Body)

	if err != nil {
		return err
	}
	return nil
}

func (d *DlObject) FetchData(siteUrl string) error {
	data, err := http.Get("http://" + siteUrl)

	if err != nil {
		return err
	}
	defer data.Body.Close()

	d.SaveData(data)

	return nil

}
