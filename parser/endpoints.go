package parser

/*
 The last struct that return something
*/

type String struct {
	Position *Position

	Value string
}

func (s *String) Out(argstore ArgStore) Values {

	return s.Value
}

type Identifier struct {
	Position *Position

	Value Values
}

// created identifiernew to distinguish between string an actual identifier
type IdentifierNew struct {
	Value Values
}

func (i *Identifier) Out(argstore ArgStore) Values {

	return IdentifierNew{Value: i.Value}
}

type NamePart struct {
	Value string
}

func (n *NamePart) Out(argstore ArgStore) Values {
	return n.Value
}
