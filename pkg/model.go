package pkg

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

const (
	flat       = "FLAT"
	percentage = "PERCENTAGE"
	ratio      = "RATIO"
)

type Payload struct {
	ID            int         `json:"ID"`
	Amount        float64     `json:"Amount"`
	Currency      string      `json:"Currency"`
	CustomerEmail string      `json:"CustomerEmail"`
	SplitInfo     []splitInfo `json:"SplitInfo"`
	sumFlat       float64
	sumRatio      float64
}
type splitInfo struct {
	SplitType     string `json:"SplitType"`
	SplitValue    int    `json:"SplitValue"`
	SplitEntityID string `json:"SplitEntityId"`
}

func (p *Payload) Bind(r *http.Request) error {
	err1 := validate.Validate(
		&validators.IntIsPresent{Name: "ID", Field: p.ID, Message: fmt.Sprintf("The %s value is required", "ID")},
		// &validators.IntIsPresent{Name: "Amount", Field: p.Amount, Message: fmt.Sprintf("The %s value is required", "Amount")},
		&validators.StringIsPresent{Name: "Currency", Field: p.Currency, Message: fmt.Sprintf("The %s value is required", "Currency")},
		&validators.EmailIsPresent{Name: "CustomerEmail", Field: p.CustomerEmail, Message: fmt.Sprintf("The %s value is required", "CustomerEmail")},
		&validators.FuncValidator{Name: "SplitInfo", Field: "SplitInfo", Message: fmt.Sprintf("The number of %s is between 1 and 20 (inclusive)", "SplitInfo"), Fn: func() bool {
			if len(p.SplitInfo) < 1 || len(p.SplitInfo) > 20 {
				return false
			}
			return true
		}},
		&validators.FuncValidator{Name: "SplitInfo", Field: "SplitInfo", Message: fmt.Sprintf("The sum of flat split is greater than amount in  %s", "SplitInfo"), Fn: func() bool {
			p.sort()
			if p.sumFlat > p.Amount {
				return false
			}
			return true
		}},
	)

	if err1.HasAny() {
		return err1
	}

	return nil
}

func (p *Payload) sort() {
	// sort p.SplitInfo by SplitType
	p.SplitInfo, p.sumFlat, p.sumRatio = sortSplitInfo(&p.SplitInfo)
}

func sortSplitInfo(splitInfo *[]splitInfo) ([]splitInfo, float64, float64) {
	var sRatio, sFlat float64
	for i := 0; i < len(*splitInfo); i++ {
		switch (*splitInfo)[i].SplitType {
		case flat:
			sFlat += float64((*splitInfo)[i].SplitValue)
		case ratio:
			sRatio += float64((*splitInfo)[i].SplitValue)
		}
		for j := i + 1; j < len(*splitInfo); j++ {
			if (*splitInfo)[i].SplitType > (*splitInfo)[j].SplitType {
				(*splitInfo)[i], (*splitInfo)[j] = (*splitInfo)[j], (*splitInfo)[i]
			}
		}
	}
	return *splitInfo, sFlat, sRatio
}


func compute(p *Payload, breakDown *[]splitBreakdown) {
	for i, split := range p.SplitInfo {
		var num float64
		switch split.SplitType {
		case flat:
			num = float64(split.SplitValue)
		case percentage:
			num = p.Amount * float64(split.SplitValue) / 100
		case ratio:
			if i == len(p.SplitInfo)-1 {
				num = p.Amount
			} else {
				num = float64(split.SplitValue) * p.Amount / p.sumRatio
			}
		}

		*breakDown = append(*breakDown, splitBreakdown{
			SplitEntityID: split.SplitEntityID,
			Amount:        num,
		})
		p.Amount -= num

	}
}

type response struct {
	ID             int              `json:"ID"`
	Balance        float64          `json:"Balance"`
	SplitBreakdown []splitBreakdown `json:"SplitBreakdown"`
}
type splitBreakdown struct {
	SplitEntityID string  `json:"SplitEntityId"`
	Amount        float64 `json:"Amount"`
}
