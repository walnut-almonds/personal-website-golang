package ftp

import (
	"time"

	"github.com/jlaffaye/ftp"
)

type IFtp interface {
	Do(addr, user, passwd string, callback func(ftpOps *ftpOps) error) error
	DialAndLogin(addr, user, passwd string) (*ftpOps, error)
}

type ftpCli struct {
}

func newFtpCli() IFtp {
	return &ftpCli{}
}

func (fc *ftpCli) Do(addr, user, passwd string, callback func(ftpOps *ftpOps) error) error {
	c, err := fc.DialAndLogin(addr, user, passwd)
	if err != nil {
		return err
	}
	defer c.Quit()

	return callback(c)
}

func (fc *ftpCli) DialAndLogin(addr, user, passwd string) (*ftpOps, error) {
	ops := &ftpOps{}
	c, err := ftp.Dial(addr, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return ops, err
	}
	ops.conn = c

	if user != "" && passwd != "" {
		err := c.Login(user, passwd)
		if err != nil {
			return ops, err
		}
	}

	return ops, nil
}
