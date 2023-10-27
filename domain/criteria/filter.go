package domain_criteria

type FilterOperator string

const (
	EQUAL        FilterOperator = "="
	NOT_EQUAL                   = "<>"
	GT                          = ">"
	LT                          = "<"
	CONTAINS                    = "CONTAINS"
	NOT_CONTAINS                = "NOT_CONTAINS"
)

type Filter struct {
	Field    string         `query:"field"`
	Operator FilterOperator `query:"operator"`
	Value    string         `query:"value"`
}
