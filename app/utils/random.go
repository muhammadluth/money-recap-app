package utils

import (
	"crypto/rand"
	"money-recap-app/model/constant"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateTraceID() string {
	random, _ := rand.Prime(rand.Reader, 10)
	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	date := time.Now().Format(constant.YYYYMMDD)
	return strings.ToUpper(date + uid[:9] + random.String())
}
