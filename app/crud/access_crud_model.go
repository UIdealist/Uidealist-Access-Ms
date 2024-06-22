package crud

// Allow check struct response
type CheckResponse struct {
	BaseResponse
	Data []bool `json:"data"`
}

// Access struct to update access to resources.
type Access[T string | interface{}] struct {
	Subject T `json:"usrId" validate:"required"`
	Object  T `json:"resId" validate:"required"`
	Action  T `json:"actId" validate:"required"`
}

func (a Access[T]) ParseToArray() []T {
	return []T{a.Subject, a.Object, a.Action}
}

// AccessList struct to list a set of policies.
type AccessList[T string | interface{}] struct {
	Policies []Access[T] `json:"policies"`
}

func (a AccessList[T]) ParseToArray() [][]T {
	var policies [][]T
	for _, policy := range a.Policies {
		policies = append(policies, policy.ParseToArray())
	}
	return policies
}

// AccessUpdate struct to update access to resources.
type AccessUpdate struct {
	Old Access[string] `json:"old" validate:"required"`
	New Access[string] `json:"new" validate:"required"`
}

// AccessListUpdate struct to list a set of policies to update
type AccessListUpdate struct {
	Policies []AccessUpdate `json:"policies"`
}

func (a AccessListUpdate) ParseToArray() ([][]string, [][]string) {
	var policiesOld [][]string
	var policiesNew [][]string
	for _, policy := range a.Policies {
		policiesOld = append(policiesOld, policy.Old.ParseToArray())
		policiesNew = append(policiesNew, policy.New.ParseToArray())
	}
	return policiesOld, policiesNew
}
