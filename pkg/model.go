package pkg

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Payload struct {
	ID            int         `json:"ID"`
	Amount        int         `json:"Amount"`
	Currency      string      `json:"Currency"`
	CustomerEmail string      `json:"CustomerEmail"`
	SplitInfo     []splitInfo `json:"SplitInfo"`
}
type splitInfo struct {
	SplitType     string `json:"SplitType"`
	SplitValue    int    `json:"SplitValue"`
	SplitEntityID string `json:"SplitEntityId"`
}

func (p *Payload) Bind(r *http.Request) error {
	err1 := validate.Validate(
		&validators.IntIsPresent{Name: "ID", Field: p.ID, Message: fmt.Sprintf("The %s value is required", "ID")},
		&validators.IntIsPresent{Name: "Amount", Field: p.Amount, Message: fmt.Sprintf("The %s value is required", "Amount")},
		&validators.StringIsPresent{Name: "Currency", Field: p.Currency, Message: fmt.Sprintf("The %s value is required", "Currency")},
		&validators.StringIsPresent{Name: "CustomerEmail", Field: p.CustomerEmail, Message: fmt.Sprintf("The %s value is required", "CustomerEmail")},
		&validators.FuncValidator{Name: "SplitInfo", Message: fmt.Sprintf("The number of %s is between 1 and 20 (inclusive)", "SplitInfo"), Fn: func() bool {
			if len(p.SplitInfo) < 1  || len(p.SplitInfo) > 20 {
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

type response struct {
	ID             int              `json:"ID"`
	Balance        int              `json:"Balance"`
	SplitBreakdown []splitBreakdown `json:"SplitBreakdown"`
}
type splitBreakdown struct {
	SplitEntityID string `json:"SplitEntityId"`
	Amount        int    `json:"Amount"`
}
