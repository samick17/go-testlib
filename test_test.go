package test

import (
    "testing"
    "./assertion"
)

type BaseObject struct {
	name string
}

/*Example usage of assertion*/
func TestAssertion(t *testing.T) {
    assertion.Init(t, func(assert *assertion.Assertion) {
        assert.AssertNotEqual("abcde", "12345")
        assert.AssertIsNotNil("")
        assert.AssertIsNil(nil)
        assert.AssertEqual(BaseObject{}, BaseObject{})
        assert.AssertNotEqual(BaseObject{name:"a"}, BaseObject{name:"b"})
        assert.AssertNotEqual(&BaseObject{}, &BaseObject{})
        assert.AssertNotEqual(&BaseObject{name:"a"}, &BaseObject{name:"b"})
    })
}