package pg


type DeallocateStmt struct {
	Name *string
}


func (n *DeallocateStmt) Pos() int {
	return 0
}

