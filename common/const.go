package common

type Condition struct {
	Field    string
	Value    interface{}
	Operator Operator
}

type Operator string

const (
	OperatorEqual           Operator = "="
	OperatorNotEqual        Operator = "!="
	OperatorGreaterThan     Operator = ">"
	OperatorGreaterOrEqual  Operator = ">="
	OperatorLessThan        Operator = "<"
	OperatorLessThanOrEqual Operator = "<="
	OperatorLike            Operator = "LIKE"
	OperatorNotLike         Operator = "NOT LIKE"
	OperatorIn              Operator = "IN"
	OperatorNotIn           Operator = "NOT IN"
	OperatorBetween         Operator = "BETWEEN"
	OperatorNotBetween      Operator = "NOT BETWEEN"
)

// status of: Project, Task
const (
	StatusDeleted   = 3
	StatusCompleted = 2
)

// status of verify_emails
const (
	VerifyEmailTypeProjectInvitation = 3
)

func (c Condition) BuildQuery() string {
	return c.Field + " " + string(c.Operator)
}
