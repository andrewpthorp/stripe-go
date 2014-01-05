package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/coupons", "coupons/coupon.json")
  params := CouponParams{}
  coupon, _ := client.Coupons.Create(&params)
  assert.Equal(t, coupon.Id, "coupon_code")
}

func TestCouponsRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/coupons/id", "coupons/coupon.json")
  coupon, _ := client.Coupons.Retrieve("id")
  assert.Equal(t, coupon.Id, "coupon_code")
}

func TestCouponsDelete(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/coupons/id", "delete.json")
  res, _ := client.Coupons.Delete("id")
  assert.Equal(t, res.Deleted, true)
}

func TestCouponsAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/coupons", "coupons/coupons.json")
  coupons, _ := client.Coupons.All()
  assert.Equal(t, coupons.Count, 1)
  assert.Equal(t, coupons.Data[0].Id, "coupon_code")
}


func TestCouponsAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/coupons", "coupons/coupons.json")
  coupons, _ := client.Coupons.AllWithFilters(Filters{})
  assert.Equal(t, coupons.Count, 1)
  assert.Equal(t, coupons.Data[0].Id, "coupon_code")
}

func TestParseCouponParams(t *testing.T) {
  params := CouponParams{
    Id: "coupon_id",
    Duration: "once",
    AmountOff: 1000,
    Currency: "USD",
    DurationInMonths: 1,
    MaxRedemptions: 10,
    PercentOff: 20,
    RedeemBy: 123456789,
  }
  values := url.Values{}
  parseCouponParams(&params, &values)
  assert.Equal(t, values.Get("id"), params.Id)
  assert.Equal(t, values.Get("duration"), params.Duration)
  assert.Equal(t, values.Get("amount_off"), strconv.Itoa(params.AmountOff))
  assert.Equal(t, values.Get("currency"), params.Currency)
  assert.Equal(t, values.Get("duration_in_months"), strconv.Itoa(params.DurationInMonths))
  assert.Equal(t, values.Get("max_redemptions"), strconv.Itoa(params.MaxRedemptions))
  assert.Equal(t, values.Get("percent_off"), strconv.Itoa(params.PercentOff))
  assert.Equal(t, values.Get("redeem_by"), strconv.Itoa(params.RedeemBy))
}
