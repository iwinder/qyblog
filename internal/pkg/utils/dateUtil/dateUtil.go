package dateUtil

import (
	"fmt"
	"time"
)

// GetTimestampOfNowByType "2006/01/02"
func GetTimestampOfNowByType(formatStr string) string {
	timeNow := time.Now()
	return fmt.Sprintf(timeNow.Format(formatStr))
}
