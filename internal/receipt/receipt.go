package receipt

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

var ReceiptById = make(map[int]Receipt)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type Item struct {
	Price       string `json:"price"`
	Description string `json:"shortDescription"`
}

func SaveReceipt(id int, receipt Receipt) {
	ReceiptById[id] = receipt
}

func GetReceipt(id int) (receipt Receipt) {
	receipt = ReceiptById[id]
	return receipt
}

func NextId() (id int) {
	id = len(ReceiptById) + 1
	return id
}

func ScoreReceipt(id int) (score int) {
	receipt := GetReceipt(id)

	nameScore := scoreRetailerName(receipt.Retailer)

	t, err := decimal.NewFromString(receipt.Total)
	if err != nil {
		panic(err)
	}
	totalScore := scoreReceiptTotal(t)

	itemScore := scoreReceiptItems(receipt.Items)

	dateScore := scoreReceiptDate(receipt.PurchaseDate)

	timeScore := scoreReceiptTime(receipt.PurchaseTime)

	return totalScore + nameScore + itemScore + dateScore + timeScore
}

func scoreRetailerName(name string) (nameScore int) {
	name = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(name, "")
	nameScore = len(name)
	return nameScore
}

func scoreReceiptTotal(total decimal.Decimal) (receiptTotalScore int) {
	dollar, err := decimal.NewFromString("1.00")
	if err != nil {
		panic(err)
	}
	quarter, err := decimal.NewFromString(".25")
	if err != nil {
		panic(err)
	}

	if decimal.Decimal.Cmp(total.Mod(dollar), decimal.Zero) == 0 {
		receiptTotalScore = 50
	} else {
		receiptTotalScore = 0
	}
	if decimal.Decimal.Cmp(total.Mod(quarter), decimal.Zero) == 0 {
		receiptTotalScore = receiptTotalScore + 25
	}

	return receiptTotalScore
}

func scoreReceiptItems(items []Item) (receiptItemsScore int) {
	num_of_items := float64(len(items))
	receiptItemsScore = int(math.Floor(num_of_items/2)) * 5

	for i := 0; i < len(items); i++ {
		item := items[i]

		if math.Mod(float64(len(item.Description)), 3) == 0 {
			price, err := decimal.NewFromString(item.Price)
			if err != nil {
				panic(err)
			}
			scalar, err := decimal.NewFromString("0.2")
			if err != nil {
				panic(err)
			}
			s := decimal.Decimal.Ceil(price.Mul(scalar))
			score, err := strconv.Atoi(s.String())
			if err != nil {
				panic(err)
			}
			receiptItemsScore = receiptItemsScore + score
		}

	}

	return receiptItemsScore
}

func scoreReceiptDate(date string) (receiptDateScore int) {
	t := strings.Split(date, "-")
	day, err := strconv.ParseFloat(t[2], 64)
	if err != nil {
		panic(err)
	}
	if math.Mod(day, 2) != 0 {
		receiptDateScore = 6
	} else {
		receiptDateScore = 0
	}
	return receiptDateScore
}

func scoreReceiptTime(time string) (receiptTimeScore int) {
	t := strings.Split(time, ":")
	hour, err := strconv.ParseInt(t[0], 10, 64)
	if err != nil {
		panic(err)
	}
	minute, err := strconv.ParseInt(t[1], 10, 64)
	if err != nil {
		panic(err)
	}

	if hour >= 14 && hour < 16 {
		if hour == 14 && minute == 0 {
			receiptTimeScore = 0
		} else {
			receiptTimeScore = 10
		}
	}

	return receiptTimeScore
}
