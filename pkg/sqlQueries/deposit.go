package sqlQueries

const (
	DepositTable           = "product_schema.deposits"
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
