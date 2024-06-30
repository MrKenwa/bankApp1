package sqlQueries

const (
	BalanceTable        = "payment_schema.balances"
	BalanceIDColumnName = "balance_id"
)

var (
	InsertBalanceColumns = []string{
		CardIDColumnName,
		DepositIDColumnName,
		AmountColumnName,
		CreatedAtColumnName,
	}
	GetBalanceColumns = []string{
		BalanceIDColumnName,
		CardIDColumnName,
		DepositIDColumnName,
		AmountColumnName,
		CreatedAtColumnName,
		DeletedAtColumnName,
	}
)
