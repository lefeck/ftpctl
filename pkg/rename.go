package pkg

func Rename() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//rename the file
	if err := ftpServer.Rename(SrcFile, DesFile); err != nil {
		return err
	}
	defer ftpServer.conn.Quit()
	return err
}

func (f *FTPServer) Rename(from string, to string) error {
	err := f.conn.Rename(from, to)
	if err != nil {
		return err
	}
	return nil
}
