package domain_criteria

type FilterOperator string

const (
	EQUAL        FilterOperator = "="
	NOT_EQUAL    FilterOperator = "<>"
	GT           FilterOperator = ">"
	GTE          FilterOperator = ">="
	LT           FilterOperator = "<"
	LTE          FilterOperator = "<="
	CONTAINS     FilterOperator = "CONTAINS"
	NOT_CONTAINS FilterOperator = "NOT_CONTAINS"
)

type Filter struct {
	Field    string         `query:"field"`
	Operator FilterOperator `query:"operator"`
	Value    string         `query:"value"`
}
