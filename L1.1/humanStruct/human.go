package human

import (
	"fmt"
)

type Human struct {
	Age    int
	Name   string
	Adress string
	Email  string
	Family FamilyInfo
}

type FamilyInfo struct {
	Mother  string
	Father  string
	Sister  string
	Brother string
}

func (h *Human) GetFamily() string {
	return fmt.Sprintf("Information about %v's family: \n- Mom: %v \n- Dad: %v \n- Sister: %v\n- Brother: %v",
		h.Name, h.Family.Mother, h.Family.Father, h.Family.Sister, h.Family.Brother)
}
