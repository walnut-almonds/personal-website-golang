package user_admin

import (
	"context"
	"golang.org/x/xerrors"
	"personal-website-golang/service/internal/model/po"
	utilCtx "personal-website-golang/service/internal/util/context"
	"personal-website-golang/service/internal/util/random"
	"strconv"
	"time"
)

type IInsertUser interface {
	Handle(ctx context.Context, cond *bo.UserAddCond) error
}

func newInsertUser(in digIn) IInsertUser {
	return &insertUser{
		in: in,
	}
}

type insertUser struct {
	in digIn
}

func (uc *insertUser) Handle(ctx context.Context, cond *bo.UserAddCond) error {
	authInfo, ok := utilCtx.GetAuthInfo(ctx)
	if !ok {
		// TODO:err
		return xerrors.Errorf("")
	}

	salt, err := random.GenUserPasswordSalt()
	if err != nil {
		// TODO:err
		return xerrors.Errorf("")
	}

	user := &po.UserAdmin{
		ID:            0,
		Level:         0,
		Status:        0,
		Password:      "",
		Salt:          "",
		LastLoginTime: nil,
		Remark:        nil,
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
	}

	db := uc.in.Mysql.Session()
	tx := func(db *gorm.DB) error {
		if err := uc.in.UserRepo.Insert(ctx, db, user); err != nil {
			// 用戶已經存在
			if {
				return xerrors.Errorf("addUser.UserRepo.Insert: user: %+v\n, err: %w", user, errs.UserIsExist)
			}
			return xerrors.Errorf("addUser.UserRepo.Insert: user: %+v\n, err: %w", user, err)
		}

		user, err := uc.in.UserRepo.First(ctx, db, false, func(db *gorm.DB) *gorm.DB {
			return db.Where("login = ?", user.Login)
		})
		if err != nil {
			return xerrors.Errorf("addUser.UserRepo.First: user_login: %+v\n, err: %w", user.Login, err)
		}

		if err := uc.in.AuthUserRoleRepo.Insert(ctx, db, &po.AuthUserRole{
			UserId:     user.ID,
			RoleId:     cond.RoleID,
			AddedId:    authInfo.ID,
			AddedTime:  utcNow,
			UpdateId:   authInfo.ID,
			UpdateTime: utcNow,
		}); err != nil {
			return xerrors.Errorf("addUser.AuthUserRoleRepo.Insert: err: %w", err)
		}

		// 同步至casbin
		adminInstance := uc.in.GetRBAC.Handle(ctx)
		if _, err := adminInstance.AddRoleForUser(strconv.Itoa(int(user.ID)), cond.RoleID); err != nil {
			return xerrors.Errorf("adminInstance.RemoveGroupingPolicy: %w", err)
		}

		if err := adminInstance.LoadPolicy(); err != nil {
			return xerrors.Errorf("adminInstance.LoadPolicy: %w", err)
		}

		return nil
	}

	if err := db.Transaction(tx); err != nil {
		return xerrors.Errorf("%w", err)
	}

	return nil
}
