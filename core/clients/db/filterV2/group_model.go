package filter

type GroupCondition struct {
	Or         *bool             `json:"or"`
	Conditions []*Condition      `json:"conditions"`
	Groups     []*GroupCondition `json:"groups"`
}
