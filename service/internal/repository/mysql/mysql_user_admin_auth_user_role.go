package mysql

type IMysqlUserAdminAuthUserRole interface {
}

type mysqlUserAdminAuthUserRole struct {
	in repositoryIn
}

func newMysqlUserAdminAuthUserRole(in repositoryIn) IMysqlUserAdminAuthUserRole {
	return &mysqlUserAdminAuthUserRole{
		in: in,
	}
}
