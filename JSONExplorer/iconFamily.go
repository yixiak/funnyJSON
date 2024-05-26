package JSONExplorer

type PokerIconFamily struct {
	leafIcon string

	nodeIcon string
}

func (P PokerIconFamily) GetLeaf_Icon() string {
	return P.leafIcon
}

func (P PokerIconFamily) GetNode_Icon() string {
	return P.nodeIcon
}

type PokerIconFactory struct{}

func (P PokerIconFactory) CreateIconFamily() IconFamily {
	return PokerIconFamily{
		leafIcon: "♢",
		nodeIcon: "♦",
	}
}

type DefualtIcon struct {
	leafIcon string
	nodeIcon string
}

func (P DefualtIcon) GetLeaf_Icon() string {
	return P.leafIcon
}

func (P DefualtIcon) GetNode_Icon() string {
	return P.nodeIcon
}

type DefualtIconFactory struct{}

func (P DefualtIconFactory) CreateIconFamily() IconFamily {
	return DefualtIcon{
		leafIcon: "",
		nodeIcon: "",
	}
}
