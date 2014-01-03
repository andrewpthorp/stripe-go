package stripe

// CardParams hold all of the parameters used for creating and updating Cards.
type CardParams struct {
  Number string
  ExpMonth int
  ExpYear int
  CVC string
  Name string
  AddressLine1 string
  AddressLine2 string
  AddressCity string
  AddressZip string
  AddressState string
  AddressCountry string
  Token string
}

// CouponParams hold all of the parameters used for creating Coupons.
type CouponParams struct {
  Id string
  Duration string
  AmountOff int
  Currency string
  DurationInMonths int
  MaxRedemptions int
  PercentOff int
  RedeemBy int
}

// PlanParams hold all of the parameters used for creating and updating Plans.
type PlanParams struct {
  Id string
  Amount int
  Currency string
  Interval string
  IntervalCount int
  Name string
  TrialPeriodDays int
}
