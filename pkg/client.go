package pkg

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

var (
	Host     string
	Password string
	Username string
	Timeout  time.Duration
	SrcFile  string
	Path     string
	DesFile  string
)

type FTPServer struct {
	conn *ftp.ServerConn
}

func NewFTPClient(username, password string, host string, timeout time.Duration) (*FTPServer, error) {
	ftpServer := &FTPServer{}
	hostIP := fmt.Sprintf(host + ":21")
	conn, err := ftp.Dial(hostIP, ftp.DialWithTimeout(timeout*time.Second))
	if err != nil {
		return nil, fmt.Errorf("connect to ftp server %s", err)
	}

	if err := conn.Login(username, password); err != nil {
		return nil, fmt.Errorf("login username or password failed %s", err)
	}
	ftpServer.conn = conn
	return ftpServer, nil
}

type FileInfo struct {
	filename string
	path     string
}

type FileInfoOpts func(option *FileInfo)

func WithPath(path string) FileInfoOpts {
	return func(option *FileInfo) {
		option.path = path
	}
}

func NewFileInfo(filename string, opts ...FileInfoOpts) FileInfo {
	file := &FileInfo{
		filename: filename,
	}
	for _, opt := range opts {
		opt(file)
	}
	return *file
}
