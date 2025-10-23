package constants

type Branch string

var (
	ParentBranch Branch = "parent"
	ChildBranch  Branch = "child"
)

func (b Branch) String() string {
	switch b {
	case ParentBranch:
		return "parent"
	case ChildBranch:
		return "child"
	default:
		return ""
	}
}
