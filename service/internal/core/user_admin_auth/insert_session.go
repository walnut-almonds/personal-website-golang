package user_admin_auth

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type IInsertSession interface {
	Handle(ctx context.Context, cond *bo.InsertSession) (string, error)
}

type insertSession struct {
	in digIn
}

func newInsertSession(in digIn) *IInsertSession {
	return &insertSession{
		in: in,
	}
}

func (uc *insertSession) Handle(ctx context.Context, req *bo.InsertSession) (string, error) {
	if err := uc.validateCond(ctx, req); err != nil {
		return "", xerrors.Errorf("%w", err)
	}

	db := uc.in.DB.Session()

	adminInfo, err := uc.in.UserRepo.First(ctx, db, true,
		func(db *gorm.DB) *gorm.DB {
			return db.Where("login = ?", &req.Login)
		})
	if err != nil {
		return "", xerrors.Errorf("addSession.UserRepo.First: adminLogin %s, error: %w", &req.Login, err)
	}

	if err := uc.validateLogin(ctx, db, req, adminInfo); err != nil {
		return "", xerrors.Errorf("%w", err)
	}

	token := uuid.NewString()

	tx := func(db *gorm.DB) error {
		// 刪除舊的登入權杖
		if err := uc.in.TokenRepo.Delete(ctx, db, adminInfo.ID); err != nil {
			return xerrors.Errorf("deleteSession.TokenRepo.Delete: admin_id: %s, err: %w", adminInfo.ID, err)
		}

		if err := uc.in.TokenRepo.Insert(ctx, db, &po.Token{
			AdminID: adminInfo.ID,
			Token:   token,
		}); err != nil {
			return xerrors.Errorf("deleteSession.TokenRepo.Insert: admin_id: %s, err: %w", adminInfo.ID, err)
		}

		// 用戶登入成功後，清空錯誤次數
		if adminInfo.LockNum > 0 {
			if err := uc.updateLockNum(ctx, db, adminInfo, 0); err != nil {
				return xerrors.Errorf("%w", err)
			}
		}

		return nil
	}

	if err := db.Transaction(tx); err != nil {
		return "", xerrors.Errorf("%w", err)
	}

	return token, nil
}

func (uc *insertSession) validateCond(ctx context.Context, cond *bo.AddSession) error {
	cond.Login = strings.TrimSpace(cond.Login)
	if len(cond.Login) <= 0 {
		return xerrors.Errorf("addSession.validateCond: admin: %s, err: %w", cond.Login, errs.RequestParamInvalid)
	}

	cond.Passwd = strings.TrimSpace(cond.Passwd)
	if len(cond.Passwd) <= 0 {
		return xerrors.Errorf("addSession.validateCond: admin: %s, passwd: %s, err: %w", cond.Login, cond.Passwd, errs.RequestParamInvalid)
	}

	return nil
}

// validateLogin 驗證是否登入成功
func (uc *insertSession) validateLogin(ctx context.Context, db *gorm.DB, cond *bo.AddSession, info *po.User) error {
	// 登入次數驗證，超過失敗上限就不能登入
	if info.LockNum >= admin.GetAppConfig().GetOtherConfig().Auth.LoginFailTimes {
		return xerrors.Errorf("%w", errs.AuthLoginFailOver)
	}

	encryptPasswd, err := uc.in.EncryptPasswd.Handle(ctx, cond.Passwd, info.Salt)
	if err != nil {
		return xerrors.Errorf("addSession.Encrypt.Handle: adminLogin %s, password: %s, salt: %s, error: %w", info.Login, cond.Passwd, info.Salt, err)
	}

	// 密碼不正確，將錯誤次數+1
	if encryptPasswd != info.Passwd {
		if err := uc.updateLockNum(ctx, db, info, info.LockNum+1); err != nil {
			return xerrors.Errorf("%w", err)
		}
		return xerrors.Errorf("addSession.validateLogin: enterPasswd: %s, info.Passwd: %s, err: %w", encryptPasswd, info.Passwd, errs.AuthAccountOrPasswdError)
	}

	return nil
}

// 更新錯誤次數
func (uc *insertSession) updateLockNum(ctx context.Context, db *gorm.DB, info *po.User, lockNum int) error {
	info.LockNum = lockNum
	if err := uc.in.UserRepo.Update(ctx, db, info); err != nil {
		return xerrors.Errorf("addSession.UserRepo.Update: adminLogin %s, lockNum: %d, error: %w", info.Login, lockNum, err)
	}

	return nil
}
