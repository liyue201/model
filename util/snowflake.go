/**
 * following is a copy of overeality-server.
 */

package util

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type CodeGenerator struct {
	sf *sonyflake.Sonyflake
}

func NewSonyflake() (CodeGenerator, error) {
	codeGenerate := CodeGenerator{}
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	codeGenerate.sf = sf
	if sf == nil {
		return codeGenerate, errors.New("cannot initial sonyflake")
	}
	return codeGenerate, nil
}

func (codeGenerate CodeGenerator) GenerateCode() (string, error) {
	id, err := codeGenerate.sf.NextID()
	fmt.Println(id)
	fmt.Println(fmt.Sprintf("%v", id))

	if err != nil {
		return "", errors.New("address is not valid")
	}
	return fmt.Sprintf("%v", id), nil
}
