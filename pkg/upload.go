package pkg

import (
	"errors"
	"log"
	"os"
)

func Upload() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//upload file
	if err := ftpServer.UploadFile(SrcFile, Path); err != nil {
		return err
	}

	defer ftpServer.conn.Quit()
	return err
}

func WithFileNameOrPath(filename string, path string) FileInfo {
	fileInfo := FileInfo{}
	if path == "" {
		fileInfo = NewFileInfo(filename)
	} else {
		fileInfo = NewFileInfo(filename, WithPath(path))
	}
	return fileInfo
}

func (f *FTPServer) BlukUploadFiles(filenames []string, path string) error {
	if len(filenames) == 0 {
		return errors.New("filenames is empty")
	}
	for _, filename := range filenames {
		f.UploadFile(filename, path)
	}
	return nil
}

func (f *FTPServer) UploadFile(filename string, path string) error {
	fileInfo := WithFileNameOrPath(filename, path)

	if err := f.conn.NoOp(); err != nil {
		log.Fatal(err)
	}

	if path == "" {
		if err := f.conn.ChangeDir("/"); err != nil {
			return err
		}
	} else {
		if err := f.conn.ChangeDir(fileInfo.path); err != nil {
			return err
		}
	}

	file, err := os.Open(fileInfo.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err = f.conn.Stor(fileInfo.filename, file); err != nil {
		log.Fatal(err)
	}
	return nil
}
