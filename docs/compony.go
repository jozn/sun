// filename: company.go
package company

// +gen slice:"Where,GroupBy[string],DistinctBy,SortBy,Select[string]"
type Company struct {
	Name    string
	Country string
	City    string
}

func asd() {
	cs := CompanySlice{Company{}}
	cs.Where(func(c Company) bool { return c.Country == "" })
}
