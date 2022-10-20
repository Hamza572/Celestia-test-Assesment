
package main

import (
    "testing"
)

func TestEven(t *testing.T) {
    actualString := evenOdd(2)
    expectedString := "Even"
    if actualString != expectedString {
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedString, actualString)
    }
}

func TestOdd(t *testing.T) {
    actualString := evenOdd(3)
    expectedString := "Odd"
    if actualString != expectedString {
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedString, actualString)
    }
}