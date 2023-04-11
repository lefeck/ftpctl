package pkg

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Download() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//download file
	if err = ftpServer.DownloadFile(SrcFile, Path); err != nil {
		return err
	}

	defer ftpServer.conn.Quit()
	return nil
}

func (f *FTPServer) DownloadFile(filename string, path string) error {
	fileInfo := WithFileNameOrPath(filename, path)

	var destFile string
	if ok := strings.Contains(fileInfo.filename, "/"); ok {
		fileName := strings.Split(fileInfo.filename, "/")
		destFile = fileName[len(fileName)-1]
	} else {
		destFile = fileInfo.filename
	}

	resp, err := f.conn.Retr(fileInfo.filename)
	if err != nil {
		return fmt.Errorf("fetch the file from the remote FTP server %s", err)
	}

	if path == "" {
		if err := os.Chdir("."); err != nil {
			return err
		}
	} else {
		if err := os.Chdir(fileInfo.path); err != nil {
			return err
		}
	}

	files, err := os.Create(destFile)
	if err != nil {
		return fmt.Errorf("Failed to create file", err)
	}
	defer files.Close()

	_, err = io.Copy(files, resp)
	if err != nil {
		return fmt.Errorf("Failed to save file", err)
	}

	return nil
}
