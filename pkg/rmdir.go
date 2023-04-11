package pkg

func RemoveDir() error {
	ftpServer, err := NewFTPClient(Username, Password, Host, Timeout)
	if err != nil {
		return err
	}

	//remove the dirctory recursive
	if err := ftpServer.RemoveDir(Path); err != nil {
		return err
	}
	defer ftpServer.conn.Quit()
	return err
}

func (f *FTPServer) RemoveDir(path string) error {
	err := f.conn.RemoveDirRecur(path)
	if err != nil {
		return err
	}
	return nil
}
