package payment

type PMethod struct {
	Type  string  `json:"type"`            // код способа оплаты
	ID    string  `json:"id"`              // идентификатор способа оплаты
	Saved bool    `json:"saved"`           // с помощью сохраненного способа оплаты можно проводить безакцептные списания
	Title *string `json:"title,omitempty"` // название способа оплаты
	MBankCard
	MYandexMoney
	MApplePay
	MGooglePay
}

func Method() *PMethod {
	return &PMethod{}
}

// MBankCard описывает способ оплаты банковской картой.
type MBankCard struct {
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

func (m *PMethod) BankCard(card *Card, saved bool) PMethod {
	title := "BankCard"
	m.Type = "bank_card"
	m.ID = ""
	m.Saved = saved
	m.Title = &title
	m.MBankCard.Card = card
	return *m
}

// MYandexMoney описывает способ оплаты Яндекс.Деньгами.
type MYandexMoney struct {
	AccountNumber string `json:"account_number,omitempty"` // номер кошелька в Яндекс.Деньгах, из которого заплатил пользователь
}

func (m *PMethod) YandexMoney(saved bool) PMethod {
	title := "YandexMoney"
	m.Type = "yandex_money"
	m.Title = &title
	return *m
}

// MApplePay описывает способ оплаты Apple Pay.
type MApplePay struct{}

func (m *PMethod) ApplePay(saved bool) PMethod {
	title := "ApplePay"
	m.Type = "apple_pay"
	m.Title = &title
	return *m
}

// MGooglePay описывает способ оплаты Google Pay.
type MGooglePay struct{}

func (m *PMethod) GooglePay(saved bool) PMethod {
	title := "GooglePay"
	m.Type = "google_pay"
	m.Title = &title
	return *m
}
