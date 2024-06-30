package sqlQueries

const (
	OperationTable              = "payment_schema.operations"
	OperationIDColumnName       = "operation_id"
	SenderBalanceIDColumnName   = "sender_balance_id"
	ReceiverBalanceIDColumnName = "receiver_balance_id"
	OperationTypeColumnName     = "operation_type"
)

var (
	InsertOperationColumns = []string{
		SenderBalanceIDColumnName,
		ReceiverBalanceIDColumnName,
		AmountColumnName,
		OperationTypeColumnName,
		CreatedAtColumnName,
	}
	GetOperationColumns = []string{
		OperationIDColumnName,
		SenderBalanceIDColumnName,
		ReceiverBalanceIDColumnName,
		AmountColumnName,
		OperationTypeColumnName,
		CreatedAtColumnName,
		DeletedAtColumnName,
	}
)
