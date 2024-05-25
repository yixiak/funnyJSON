package JSONExplorer

type PokerIconFamily struct {
	prefix_leafIcon string
	prefix_nodeIcon string
	suffix_leafIcon string
	suffix_nodeIcon string
}

func (P PokerIconFamily) GetLeaf_prefixIcon() string {
	return P.prefix_leafIcon
}

func (P PokerIconFamily) GetNode_prefixIcon() string {
	return P.prefix_nodeIcon
}

func (P PokerIconFamily) GetLeaf_suffixIcon() string {
	return P.suffix_leafIcon
}

func (P PokerIconFamily) GetNode_suffixIcon() string {
	return P.suffix_nodeIcon
}

type PokerIconFactory struct{}

func (P PokerIconFactory) CreateIconFamily() IconFamily {
	return PokerIconFamily{
		prefix_leafIcon: "♢",
		prefix_nodeIcon: "♦",
		suffix_leafIcon: " ",
		suffix_nodeIcon: " ",
	}
}
