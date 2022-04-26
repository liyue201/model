package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewSmartNftAbilityEnumFromString(t *testing.T) {
	ablity := &SmartNftAbility{
		Ability: SmartNftAbilityUserInvitation,
		Percentage: 30,
	}
	s, err := json.Marshal(ablity)
	if err != nil {
		t.Fatal(err)
	}
	b := &SmartNftAbility{}
	err = json.Unmarshal(s, b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(s))

}
