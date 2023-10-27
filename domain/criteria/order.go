package domain_criteria

type OrderType string

const (
	ASC  OrderType = "asc"
	DESC           = "desc"
	NONE           = "none"
)

type Order struct {
	OrderBy   string
	OrderType OrderType
}

func NewOrderAsc(orderBy string) *Order {
	return &Order{
		OrderBy:   orderBy,
		OrderType: ASC,
	}
}

func NewOrderDesc(orderBy string) *Order {
	return &Order{
		OrderBy:   orderBy,
		OrderType: DESC,
	}
}

func NewOrderNone() *Order {
	return &Order{
		OrderBy:   "",
		OrderType: NONE,
	}
}
