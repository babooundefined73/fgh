package configure

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/fgh/pkg/utils"
)

// Ask questions to fill in reglar config
func AskQuestions() (RegularOutline, utils.CtxErr) {
	questions := []*survey.Question{
		{
			Name: "CloneClipboard",
			Prompt: &survey.Confirm{
				Message: "Do you want to copy the path of a cloned repo after clone to your clipboard?",
			},
		},
		{
			Name: "StructureRoot",
			Prompt: &survey.Input{
				Message: "Where should the structure start relative to your home folder? (default is github/ enter nothing to use default)",
				Help:    "See https://github.com/Matt-Gleich/fgh#-structure_root for more info.",
			},
		},
		{
			Name: "LowercaseLang",
			Prompt: &survey.Confirm{
				Message: "Should language folders be lowercase?",
				Help: fmt.Sprintf(
					"%v would become %v\n",
					filepath.FromSlash("~/github/Matt-Gleich/public/Go/fgh"),
					filepath.FromSlash("~/github/Matt-Gleich/public/go/fgh"),
				),
			},
		},
		{
			Name: "SpaceChar",
			Prompt: &survey.Input{
				Message: "If a language name has a space in it what should the space be replaced with? (default is - enter nothing to use default)",
			},
		},
	}
	var answers RegularOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		return RegularOutline{}, utils.CtxErr{
			Context: "Failed to ask questions about config",
			Error:   err,
		}
	}
	return answers, utils.CtxErr{}
}
