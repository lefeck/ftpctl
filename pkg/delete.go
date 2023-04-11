package pkg

func Delete() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//delete file
	if err := ftpServer.DeleteFile(SrcFile); err != nil {
		return err
	}

	defer ftpServer.conn.Quit()
	return err
}

func (f *FTPServer) DeleteFile(filename string) error {
	err := f.conn.Delete(filename)
	if err != nil {
		return err
	}
	return nil
}
