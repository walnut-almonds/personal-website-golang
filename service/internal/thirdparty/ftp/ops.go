package ftp

import (
	"github.com/jlaffaye/ftp"
)

type ftpOps struct {
	conn *ftp.ServerConn
}

func (fo *ftpOps) Quit() error {
	if fo.conn == nil {
		return nil
	}

	return fo.conn.Quit()
}
