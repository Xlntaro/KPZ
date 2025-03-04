package main

import "fmt"

type Subscription interface {
	GetPrice() float64
	GetMinPeriod() int
	GetChannels() []string
}
type DomesticSubscription struct{}

func (d DomesticSubscription) GetPrice() float64 {
	return 10.99
}
func (d DomesticSubscription) GetMinPeriod() int {
	return 1
}

func (d DomesticSubscription) GetChannels() []string {
	return []string{"News", "Entertaindent", "Kids"}
}

type EducationalSubscription struct{}

func (e EducationalSubscription) GetPrice() float64 {
	return 7.99
}
func (e EducationalSubscription) GetMinPeriod() int {
	return 3
}

func (e EducationalSubscription) GetChannels() []string {
	return []string{"Discovery", "Science", "History"}
}

type PremiumSubscription struct{}

func (p PremiumSubscription) GetPrice() float64 {
	return 19.99
}

func (p PremiumSubscription) GetMinPeriod() int {
	return 1
}

func (p PremiumSubscription) GetChannels() []string {
	return []string{"HBO", "Netflix", "Sports", "News"}
}

type SubscriptionFactory interface {
	CreateSubscription() Subscription
}
type WebSiteFactory struct{}

func (w WebSiteFactory) CreateSubscription() Subscription {
	return DomesticSubscription{}
}

type MobileAppFactory struct{}

func (m MobileAppFactory) CreateSubscription() Subscription {
	return EducationalSubscription{}
}

type ManagerCallFactory struct{}

func (m ManagerCallFactory) CreateSubscription() Subscription {
	return PremiumSubscription{}
}
func main() {
	var factory SubscriptionFactory
	fmt.Println("1 Купівля на сайті ")
	fmt.Println("2 Купівля через мобільний додаток")
	fmt.Println("3 Купівля через дзвінок менеджера")
	fmt.Println("Виберіть спосіб покупки: ")

	var choise int 
	fmt.Scanln(&choise)

	switch choise {
	case 1:
		factory = WebSiteFactory{}
	case 2:
		factory = MobileAppFactory{}
	case 3:
		factory = ManagerCallFactory{}
	default:
		fmt.Println("Невірний вибір ")
		return	
	}
	subscription := factory.CreateSubscription()

	fmt.Println("Тип підписки:", fmt.Sprintf("%T", subscription))
	fmt.Println("Ціна:", subscription.GetPrice())
	fmt.Println("Мінімальний період:", subscription.GetMinPeriod(), "місяців")
	fmt.Println("Доступні канали:", subscription.GetChannels())
	
}