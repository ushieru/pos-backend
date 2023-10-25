package domain_criteria

type Criteria struct {
	Filters []Filter
	Order   Order
}

func NewCriteria(filters []Filter, order Order) *Criteria {
	return &Criteria{filters, order}
}
