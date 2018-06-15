package test

import (
    "testing"
    "./assertion"
)

type BaseObject struct {

}

/*Example usage of assertion*/
func TestAssertion(t *testing.T) {
    assertion.Init(t, func(assert *assertion.Assertion) {
        assert.AssertNotEqual("abcde", "12345")
        assert.AssertIsNotNil("")
        assert.AssertIsNil(nil)
    })
}