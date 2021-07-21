package ticket_service

import (
	"fmt"
	"strings"
	"ticket-crawler/pkg/util"
	"time"
)

type TicketRequest struct {
	FromCity  string `json:"fromCity"`
	ToCity    string `json:"toCity"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

var layout = "2006-01-02 15:04:05"

func (tq TicketRequest) Climb() {
	startTime, _ := time.Parse(layout, tq.StartDate)
	endTime, _ := time.Parse(layout, tq.EndDate)
	// 遍历日期区间
	for rd := util.RangeDate(startTime, endTime); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fmt.Println(strings.Split(date.String(), " ")[0])
	}
}

type ticket struct {
}

func ClimbTicketInfo() {

}
