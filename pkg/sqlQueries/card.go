package sqlQueries

const (
	CardTable            = "product_schema.cards"
	CardNumberColumnName = "card_number"
	CardTypeColumnName   = "card_type"
	PinCodeColumnName    = "pin_code"
)

var (
	InsertCardColumns = []string{
		CardNumberColumnName,
		UserIDColumnName,
		CardTypeColumnName,
		PinCodeColumnName,
		CreatedAtColumnName,
	}
	GetCardColumns = []string{
		CardIDColumnName,
		CardNumberColumnName,
		UserIDColumnName,
		CardTypeColumnName,
		PinCodeColumnName,
		CreatedAtColumnName,
		DeletedAtColumnName,
	}
)
