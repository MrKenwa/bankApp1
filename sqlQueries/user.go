package sqlQueries

const (
	UserTable            = "users"
	UserIDColumnName     = "id"
	NameColumnName       = "name"
	LastNameColumnName   = "last_name"
	PatronymicColumnName = "patronymic"
	EmailColumnName      = "email"
	PasswordColumnName   = "password"
	PassportColumnName   = "passport_number"
	CreatedAtColumnName  = "created_at"
	DeletedAtColumnName  = "deleted_at"
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
		UserIDColumnName,
		NameColumnName,
		LastNameColumnName,
		PatronymicColumnName,
		EmailColumnName,
		PasswordColumnName,
		PassportColumnName,
		CreatedAtColumnName,
	}
)
