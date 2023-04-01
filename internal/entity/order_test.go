package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_ID_Is_Blank_Get_An_Error(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_Price_Is_Blank_Get_An_Error(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_Tax_Is_Blank_Get_An_Error(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_All_Valid_Params(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	order.CalculateFinalPrice()
	assert.NoError(t, order.Validate())
	assert.Equal(t, 1.0, order.Tax)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 11.0, order.FinalPrice)
}
