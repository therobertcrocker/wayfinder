package forge

import (
	"strconv"

	"github.com/therobertcrocker/wayfinder/pkg/utils"

	"github.com/rivo/tview"
)

var (
	levelChoices = []string{
		strconv.Itoa(utils.PartyLevel - 2),
		strconv.Itoa(utils.PartyLevel - 1),
		strconv.Itoa(utils.PartyLevel),
		strconv.Itoa(utils.PartyLevel + 1),
		strconv.Itoa(utils.PartyLevel + 2),
		strconv.Itoa(utils.PartyLevel + 3),
	}
)

func QuestForm() tview.Primitive {
	f := tview.NewForm().
		AddDropDown("Class:", utils.QuestTypes, 0, nil).
		AddDropDown("Level:", levelChoices, 0, nil).
		AddInputField("Title:", "", 20, nil, nil).
		AddTextArea("Summary:", "", 0, 2, 0, nil).
		AddInputField("Source:", "", 20, nil, nil).
		AddInputField("Gold:", "", 20, nil, nil).
		AddDropDown("Treasure Rating:", []string{"-2", "-1", "+0", "+1", "+2"}, 0, nil).
		AddDropDown("Reputation:", []string{"-5", "-2", "-1", "+0", "+1", "+2", "+5"}, 0, nil)
	f.SetBorder(true).SetTitle("New Quest")
	return f
}
