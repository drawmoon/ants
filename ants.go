package ants

// Dialects and dialect families as supported by ants.
type SQLDialect int32

const (
	// Default is the default dialect.
	//
	// This dialect was chosen in the absence of a more definitive dialect. It is
	// not intended to be used with any actual database, and it is ultimately interpreted as
	// a serializable string format.
	Default SQLDialect = 0

	Postgres SQLDialect = 1 // The PostgreSQL dialect family.
	MySQL    SQLDialect = 2 // The MySQL dialect family.
)

// Represents a single query entry.
type QueryEntry struct {
	Items    []*QueryItem   `json:"items"`    // Measure or Dimensions for the request.
	Criteria *QueryCriteria `json:"criteria"` // Gets or sets the filter criteria of the query.
	Sort     *QuerySort     `json:"sort"`     // Gets or sets the sorting of the query.
}

// QueryItem represents a measure or dimension in a query.
type QueryItemType int32

const (
	Measure   QueryItemType = 0 // A measure and representing a aggregate column.
	Dimension QueryItemType = 1 // A dimension and representing a column.
)

// Measure arithmetic options.
type Arithmetic int32

const (
	NonArithmetic  Arithmetic = 0  // The default arithmetic options.
	Sum            Arithmetic = 1  // The sum aggregate operation.
	Count          Arithmetic = 2  // The count aggregate operation.
	DistinctCount  Arithmetic = 3  // The distinct count aggregate operation.
	Min            Arithmetic = 4  // The min aggregate operation.
	Max            Arithmetic = 5  // The max aggregate operation.
	Avg            Arithmetic = 6  // The avg aggregate operation.
	Median         Arithmetic = 7  // The median aggregate operation.
	PercentileCont Arithmetic = 8  // The percentile cont aggregate operation.
	Fpy            Arithmetic = 9  // The fpy arithmetic.
	Yoy            Arithmetic = 10 // The yoy arithmetic.
	Qoq            Arithmetic = 11 // The qoq arithmetic.
	Percentage     Arithmetic = 12 // The percentage arithmetic.
)

// A measure or dimension and representing a column.
type QueryItem struct {
	Id    string `json:"id"`    // TableField used to specify the query.
	Alias string `json:"alias"` // Gets or sets the shown name of the QueryItem.

	// Gets or sets the type of the QueryItem.
	//
	// Possible values are:
	//	0: Measure
	//	1: Dimension
	Type QueryItemType `json:"type"`

	// Gets or sets the arithmetic of the QueryItem.
	//
	// Possible values are:
	//  0: None
	//	1: Sum
	//	2: Count
	//	3: DistinctCount
	//  4: Min
	//	5: Max
	//	6: Avg
	//	7: Median
	//	8: PercentileCont
	//  9: Fpy
	//  10: Yoy
	//  11: Qoq
	//  12: Percentage
	Arithmetic Arithmetic `json:"arithmetic"`
}

type QueryCriteria struct {
}

type QuerySort struct {
}
