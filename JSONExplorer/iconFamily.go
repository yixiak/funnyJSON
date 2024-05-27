package JSONExplorer

type PokerLeafIcon struct {
	leafIcon string
}
type PokerNodeIcon struct {
	nodeIcon string
}

func (P PokerLeafIcon) GetLeaf_Icon() string {
	return P.leafIcon
}

func (P PokerNodeIcon) GetNode_Icon() string {
	return P.nodeIcon
}

type PokerLeafIconFactory struct{}

type PokerNodeIconFactory struct{}

func (P PokerLeafIconFactory) CreateLeafIcon() LeafIcon {
	return PokerLeafIcon{
		leafIcon: "♧",
	}
}

func (P PokerNodeIconFactory) CreateNodeIcon() NodeIcon {
	return PokerNodeIcon{
		nodeIcon: "♢",
	}
}

type DefualtLeafIcon struct {
	leafIcon string
}
type DefualtNodeIcon struct {
	nodeIcon string
}

func (P DefualtLeafIcon) GetLeaf_Icon() string {
	return P.leafIcon
}

func (P DefualtNodeIcon) GetNode_Icon() string {
	return P.nodeIcon
}

type DefualtLeafIconFactory struct{}

type DefualtNodeIconFactory struct{}

func (P DefualtLeafIconFactory) CreateLeafIcon() LeafIcon {
	return DefualtLeafIcon{
		leafIcon: "",
	}
}

func (P DefualtNodeIconFactory) CreateNodeIcon() NodeIcon {
	return DefualtNodeIcon{
		nodeIcon: "",
	}
}
