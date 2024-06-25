package sqlQueries

const (
	BalanceTable        = "balances"
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
