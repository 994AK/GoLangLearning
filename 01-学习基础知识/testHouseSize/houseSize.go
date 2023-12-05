package testHouseSize

import (
	"fmt"
	"math/rand"
)

// AllocationSize 分配房子大小
func AllocationSize(name string) string {
	message := fmt.Sprintf("%v-%v 平方", name, rand.Intn(500))
	return message
}
