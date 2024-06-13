package sqlQueries

const (
	UserTable            = "users"
	IDColumnName         = "id"
	NameColumnName       = "name"
	LastNameColumnName   = "lastname"
	PatronymicColumnName = "patronymic"
	EmailColumnName      = "email"
	PasswordColumnName   = "password"
	PassportColumnName   = "passport_number"
)

var (
	InsertUserColumns = []string{
		NameColumnName,
		LastNameColumnName,
		PatronymicColumnName,
		EmailColumnName,
		PasswordColumnName,
		PassportColumnName,
		CreatedAtColumnName,
	}
	GetUserColumns = []string{
		IDColumnName,
		NameColumnName,
		LastNameColumnName,
		PatronymicColumnName,
		EmailColumnName,
		PasswordColumnName,
		PassportColumnName,
		CreatedAtColumnName,
		DeletedAtColumnName,
	}
)
