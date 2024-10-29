package action

type Action interface {
	Initiator() string
	Target() string
	GetType() string
	Intent() string
}

type SimpleAction struct {
	From              string
	To                string
	Type              string
	IntentDescription string
}

func (a *SimpleAction) Initiator() string {
	return a.From
}

func (a *SimpleAction) Target() string {
	return a.To
}

func (a *SimpleAction) GetType() string {
	return a.Type
}

func (a *SimpleAction) Intent() string {
	return a.IntentDescription
}

var _ Action = &SimpleAction{}
