package user_admin

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type IHashPassword interface {
	Handle(ctx context.Context, passwd, salt string) string
}

func newHashPassword(in digIn) IHashPassword {
	return &hashPassword{
		in: in,
	}
}

type hashPassword struct {
	in digIn
}

func (uc *hashPassword) Handle(ctx context.Context, password, userSalt string) string {
	serverSalt := uc.in.OpsConf.GetOpsServerConfig().PasswordSalt
	
	saltedPassword := fmt.Sprintf("%s_%s_%s", password, serverSalt, userSalt)
	hashedPassword := sha256.Sum256([]byte(saltedPassword))

	return hex.EncodeToString(hashedPassword[:])
}
