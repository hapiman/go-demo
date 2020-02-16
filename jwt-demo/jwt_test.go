package jwt_demo

import (
	"fmt"
	"testing"
)

func TestJwtDemo(t *testing.T) {
	tokenStr, err := CreateToken(103787)
	if err != nil {
		t.Fatalf("create token met error: %s", err.Error())
	}

	memberId, err := ParseToken(tokenStr)
	if err != nil {
		t.Fatalf("parse token met error: %s", err.Error())
	}

	fmt.Printf("tokenStr: %s, memberId: %d \n", tokenStr, memberId)
}
