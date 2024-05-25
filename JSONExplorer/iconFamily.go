package JSONExplorer

type PokerIconFamily struct {
	leafIcon string
	nodeIcon string
}

func (P PokerIconFamily) GetLeafIcon() string {
	return P.leafIcon
}

func (P PokerIconFamily) GetNodeIcon() string {
	return P.nodeIcon
}

type PokerIconFactory struct{}

func (P PokerIconFactory) CreateIconFamily() IconFamily {
	return PokerIconFamily{
		leafIcon: "♢",
		nodeIcon: "♦",
	}
}
