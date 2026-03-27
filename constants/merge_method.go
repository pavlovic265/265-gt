package constants

type MergeMethod string

var (
	MergeMethodMerge  MergeMethod = "merge"
	MergeMethodSquash MergeMethod = "squash"
	MergeMethodRebase MergeMethod = "rebase"
	MergeMethodQueue  MergeMethod = "queue"
)

func (m MergeMethod) String() string {
	switch m {
	case MergeMethodMerge:
		return "merge"
	case MergeMethodSquash:
		return "squash"
	case MergeMethodRebase:
		return "rebase"
	case MergeMethodQueue:
		return "queue"
	default:
		return ""
	}
}
