package user_admin

import (
	"context"

	"golang.org/x/xerrors"

	"personal-website-golang/service/internal/constant"
)

type ICheckPassword interface {
	Handle(ctx context.Context, password string) error
}

func newCheckPassword(in digIn) ICheckPassword {
	return &checkPassword{
		in: in,
	}
}

type checkPassword struct {
	in digIn
}

func (uc *checkPassword) Handle(ctx context.Context, password string) error {
	passwordLen := len(password)
	if (passwordLen < constant.PasswordMinLen) || (passwordLen > constant.PasswordMaxLen) {
		return xerrors.Errorf("密碼長度限制 %d~%d", constant.PasswordMinLen, constant.PasswordMaxLen)
	}

	match := constant.CheckPasswordRegex.MatchString(password)
	if match {
		return xerrors.Errorf("請勿使用英文大小寫與數字以外的字元")
	}

	return nil
}
