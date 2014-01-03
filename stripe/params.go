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
}

