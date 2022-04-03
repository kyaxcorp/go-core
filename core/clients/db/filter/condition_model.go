package filter

type Condition struct {
	Or *bool `json:"Or"`
	// Group       *GroupCondition `json:"Group"`
	Eq    *Eq    `json:"Eq"`
	NotEq *NotEq `json:"NotEq"`

	Ht *Ht `json:"Ht"`
	He *He `json:"He"`

	Lt *Lt `json:"Lt"`
	Le *Le `json:"Le"`

	Contains    *Contains    `json:"Contains"`
	NotContains *NotContains `json:"NotContains"`

	BeginsWith    *BeginsWith    `json:"BeginsWith"`
	NotBeginsWith *NotBeginsWith `json:"NotBeginsWith"`

	EndsWith    *EndsWith    `json:"EndsWith"`
	NotEndsWith *NotEndsWith `json:"NotEndsWith"`

	// TODO: add JSON functions: in array, has key etc...

	In    *In    `json:"In"`
	NotIn *NotIn `json:"NotIn"`

	IsNull  *IsNull  `json:"IsNull"`
	NotNull *NotNull `json:"NotNull"`

	IsTrue  *IsTrue  `json:"IsTrue"`
	IsFalse *IsFalse `json:"IsFalse"`

	Empty    *Empty    `json:"Empty"`
	NotEmpty *NotEmpty `json:"NotEmpty"`

	Between    *Between    `json:"Between"`
	NotBetween *NotBetween `json:"NotBetween"`
}
