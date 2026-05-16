package bw

type CardBrand string

const (
	CardBrandVisa            CardBrand = "Visa"
	CardBrandMastercard      CardBrand = "Mastercard"
	CardBrandAmericanExpress CardBrand = "Amex"
	CardBrandDiscover        CardBrand = "Discover"
	CardBrandDinersClube     CardBrand = "Diners Club"
	CardBrandJCB             CardBrand = "JCB"
	CardBrandMaestro         CardBrand = "Maestro"
	CardBrandUnionPay        CardBrand = "UnionPay"
	CardBrandRubPay          CardBrand = "RubPay"
	CardBrandOther           CardBrand = "Other"
)

type Card struct {
	CardholderName string    `json:"cardholderName"`
	Brand          CardBrand `json:"brand"`
	Number         string    `json:"number"`
	ExpMonth       string    `json:"expMonth"`
	ExpYear        string    `json:"expYear"`
	Code           string    `json:"code"`
}
