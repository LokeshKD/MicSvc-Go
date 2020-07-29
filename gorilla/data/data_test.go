package data

import "testing"

func TestValidate(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 1.99,
		SKU:   "abc-123",
	}
	// p := &Product{} for failure - nothing in it!

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
