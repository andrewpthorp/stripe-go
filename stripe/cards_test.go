package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"strconv"
	"testing"
)

func TestCardCreate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards", "cards/card.json")
	params := CardParams{}
	card, _ := client.Cards.Create("cus_123456789", &params)
	assert.Equal(t, card.Id, "card_123456789")
}

func TestCardsRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards/card_123456789", "cards/card.json")
	card, _ := client.Cards.Retrieve("cus_123456789", "card_123456789")
	assert.Equal(t, card.Id, "card_123456789")
}

func TestCardsUpdate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards/card_123456789", "cards/card.json")
	card, _ := client.Cards.Update("cus_123456789", "card_123456789", new(CardParams))
	assert.Equal(t, card.Id, "card_123456789")
}

func TestCardsDelete(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards/card_123456789", "delete.json")
	res, _ := client.Cards.Delete("cus_123456789", "card_123456789")
	assert.Equal(t, res.Deleted, true)
}

func TestCardsAll(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards", "cards/cards.json")
	cards, _ := client.Cards.All("cus_123456789")
	assert.Equal(t, cards.Count, 1)
	assert.Equal(t, cards.Data[0].Id, "card_123456789")
}

func TestCardsAllWithFilters(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/cards", "cards/cards.json")
	cards, _ := client.Cards.AllWithFilters("cus_123456789", Filters{})
	assert.Equal(t, cards.Count, 1)
	assert.Equal(t, cards.Data[0].Id, "card_123456789")
}

func TestParseCardParamsWithRoot(t *testing.T) {
	params := CardParams{
		Number:         "4242424242424242",
		ExpMonth:       01,
		ExpYear:        2020,
		CVC:            "111",
		Name:           "Andrew Thorp",
		AddressLine1:   "1 Something Lane",
		AddressLine2:   "Suite 200",
		AddressCity:    "San Francisco",
		AddressZip:     "94110",
		AddressState:   "CA",
		AddressCountry: "USA",
	}
	values := url.Values{}

	// With Root card[]
	parseCardParams(&params, &values, true)
	assert.Equal(t, values.Get("card[number]"), params.Number)
	assert.Equal(t, values.Get("card[exp_month]"), strconv.Itoa(params.ExpMonth))
	assert.Equal(t, values.Get("card[exp_year]"), strconv.Itoa(params.ExpYear))
	assert.Equal(t, values.Get("card[cvc]"), params.CVC)
	assert.Equal(t, values.Get("card[name]"), params.Name)
	assert.Equal(t, values.Get("card[address_line1]"), params.AddressLine1)
	assert.Equal(t, values.Get("card[address_line2]"), params.AddressLine2)
	assert.Equal(t, values.Get("card[address_city]"), params.AddressCity)
	assert.Equal(t, values.Get("card[address_zip]"), params.AddressZip)
	assert.Equal(t, values.Get("card[address_state]"), params.AddressState)
	assert.Equal(t, values.Get("card[address_country]"), params.AddressCountry)

	// Without root card[]
	parseCardParams(&params, &values, false)
	assert.Equal(t, values.Get("number"), params.Number)
	assert.Equal(t, values.Get("exp_month"), strconv.Itoa(params.ExpMonth))
	assert.Equal(t, values.Get("exp_year"), strconv.Itoa(params.ExpYear))
	assert.Equal(t, values.Get("cvc"), params.CVC)
	assert.Equal(t, values.Get("name"), params.Name)
	assert.Equal(t, values.Get("address_line1"), params.AddressLine1)
	assert.Equal(t, values.Get("address_line2"), params.AddressLine2)
	assert.Equal(t, values.Get("address_city"), params.AddressCity)
	assert.Equal(t, values.Get("address_zip"), params.AddressZip)
	assert.Equal(t, values.Get("address_state"), params.AddressState)
	assert.Equal(t, values.Get("address_country"), params.AddressCountry)

	// Token
	params = CardParams{Token: "tok_123456789"}
	parseCardParams(&params, &values, false)
	assert.Equal(t, values.Get("card"), params.Token)
	assert.NotEqual(t, values.Get("number"), params.Number)
	assert.NotEqual(t, values.Get("card[number]"), params.Number)
}
