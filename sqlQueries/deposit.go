package sqlQueries

const (
	DepositTable           = "deposits"
	DepositIDColumnName    = "deposit_id"
	DepositTypeColumnName  = "deposit_type"
	InterestRateColumnName = "interest_rate"
)

var (
	InsertDepositColumns = []string{
		UserIDColumnName,
		DepositTypeColumnName,
		InterestRateColumnName,
		CreatedAtColumnName,
	}
	GetDepositColumns = []string{
		DepositIDColumnName,
		UserIDColumnName,
		DepositTypeColumnName,
		InterestRateColumnName,
		CreatedAtColumnName,
		DeletedAtColumnName,
	}
)
