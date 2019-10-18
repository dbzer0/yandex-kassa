package payment

type Method struct {
	Type  string  `json:"type"`            // код способа оплаты
	ID    string  `json:"id"`              // идентификатор способа оплаты
	Saved bool    `json:"saved"`           // с помощью сохраненного способа оплаты можно проводить безакцептные списания
	Title *string `json:"title,omitempty"` // название способа оплаты
	BankCard
	YandexMoney
	ApplePay
	GooglePay
}

// BankCard описывает способ оплаты банковской картой.
type BankCard struct {
	Card *Card `json:"card,omitempty"` // данные банковской карты
}

type Card struct {
	First6        string `json:"first6"`                   // первые 6 цифр номера карты (BIN)
	Last4         string `json:"last4"`                    // последние 4 цифры номера карты
	ExpiryYear    string `json:"expiry_year"`              // срок действия, год, YYYY
	ExpiryMonth   string `json:"expiry_month"`             // срок действия, месяц, MM
	CardType      string `json:"card_type"`                // тип банковской карты: MasterCard (для карт Mastercard и Maestro), Visa (для карт Visa и Visa Electron), Mir, UnionPay, JCB, AmericanExpress, DinersClub и Unknown
	IssuerCountry string `json:"issuer_country,omitempty"` // код страны, в которой выпущена карта в формате ISO-3166 alpha-2
	IssuerName    string `json:"issuer_name,omitempty"`    // наименование банка, выпустившего карту
	Source        string `json:"source,omitempty"`         // источник данных банковской карты. Возможное значение: google_pay
}

// YandexMoney описывает способ оплаты Яндекс.Деньгами.
type YandexMoney struct {
	AccountNumber string `json:"account_number,omitempty"` // номер кошелька в Яндекс.Деньгах, из которого заплатил пользователь
}

// ApplePay описывает способ оплаты Apple Pay.
type ApplePay struct{}

// GooglePay описывает способ оплаты Google Pay.
type GooglePay struct{}
