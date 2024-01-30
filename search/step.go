package search

type Step struct {
	CurrentPosition *Point
	PreviousStep    *Step
}

func (s Step) Equal(b any) bool {

	if bStep, ok := b.(Step); ok {
		return s.CurrentPosition.Equal(*bStep.CurrentPosition)
	}
	return false // consider vomiting an error
}
