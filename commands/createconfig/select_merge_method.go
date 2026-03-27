package createconfig

import (
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/ui/components"
)

var mergeMethodChoices = []string{
	constants.MergeMethodSquash.String(),
	constants.MergeMethodRebase.String(),
	constants.MergeMethodMerge.String(),
}

func HandleSelectMergeMethod() (constants.MergeMethod, error) {
	selected, err := components.SelectString(mergeMethodChoices)
	if err != nil {
		return "", err
	}

	return constants.MergeMethod(selected), nil
}
