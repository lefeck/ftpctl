package pkg

import (
	"fmt"
	"github.com/jlaffaye/ftp"
)

func Get() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//get all of the files in the specify dirctory
	_, err = ftpServer.List(Path)
	if err != nil {
		return err
	}

	defer ftpServer.conn.Quit()
	return err
}

func (f *FTPServer) List(path string) ([]*ftp.Entry, error) {
	entries, err := f.conn.List(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fmt.Printf("%-60s %s %-8s %-10d %s\n", entry.Name, entry.Target, entry.Type, entry.Size, entry.Time.String())
	}
	return entries, nil
}
