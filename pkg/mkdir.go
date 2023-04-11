package pkg

import "errors"

func MakeDir() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//create the new dirctory
	if err := ftpServer.MakeDir(Path); err != nil {
		return err
	}

	defer ftpServer.conn.Quit()
	return err
}

func (f *FTPServer) MakeDir(path string) error {
	if path == "" {
		return errors.New("path is empty")
	}

	if err := f.conn.MakeDir(path); err != nil {
		return err
	}
	return nil
}
