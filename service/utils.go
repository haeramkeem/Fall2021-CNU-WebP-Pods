package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

const BASE string = `https://leetcode.com`;

// Generate Random number
func getRand(max int) (int, error) {
    rand, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return int(rand.Int64()), err
}

var monthOf = map[string]time.Month{
    "01":   time.January,
    "02":   time.February,
    "03":   time.March,
    "04":   time.April,
    "05":   time.May,
    "06":   time.June,
    "07":   time.July,
    "08":   time.August,
    "09":   time.September,
    "10":   time.October,
    "11":   time.November,
    "12":   time.December,
}

func parseStrDate(strDate string) time.Time {
    strYear := strDate[:2]
    strMonth := strDate[2:4]
    strDay := strDate[4:]

    nYear , _ := strconv.Atoi(strYear)
    nYear += 2000
    tMonth := monthOf[strMonth]
    nDay, _ := strconv.Atoi(strDay)

    loc, _ := time.LoadLocation("Asia/Seoul")

    return time.Date(nYear, tMonth, nDay, 0, 0, 0, 0, loc)
}

func nowStrDate() string {
    now := time.Now()
    return fmt.Sprintf("%02d%02d%02d", now.Year() - 2000, now.Month(), now.Day())
}
