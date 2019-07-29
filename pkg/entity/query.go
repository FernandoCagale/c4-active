package entity

//Query options
type Query struct {
	search string
	limit  int
	page   int
}

//NewQuery newQuery
func NewQuery(search string, limit, page int) *Query {
	return &Query{search, limit, page}
}

//IsEmpty isEmpty Query
func (q *Query) IsEmpty() bool {
	return q.search == ""
}

//GetLimit getLimit default 10
func (q *Query) GetLimit() int {
	if q.limit == 0 {
		return 10
	}
	return q.limit
}

//GetOffset getOffset
func (q *Query) GetOffset() int {
	return q.limit * (q.page - 1)
}

//GetSearch getSearch
func (q *Query) GetSearch() string {
	return "%" + q.search + "%"
}
