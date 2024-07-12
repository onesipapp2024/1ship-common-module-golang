package helper

import (
	"crypto/rand"
	"errors"
	"github.com/mileusna/useragent"
	"hash/fnv"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SyncMutexMap struct {
	MutexMap *sync.Map
	KeyNum   int
}

type PaginateInfo struct {
	Total       uint
	PerPage     uint
	CurrentPage uint
	PrevPage    uint
	NextPage    uint
	FirstPage   uint
	LastPage    uint
}

func GenerateRandomNumberStringWithLength(length int) (string, error) {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		ran, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return "", err
		}
		builder.WriteString(strconv.FormatInt(ran.Int64(), 10))
	}
	return builder.String(), nil
}

func GenerateRandomStringWithLength(length int) (string, error) {
	var builder strings.Builder
	randMax := int64(len(RandomStringSource)) - 1
	for i := 0; i < length; i++ {
		ran, err := rand.Int(rand.Reader, big.NewInt(randMax))
		if err != nil {
			return "", err
		}
		builder.WriteString(string(RandomStringSource[ran.Int64()]))
	}
	return builder.String(), nil
}

func GenerateRandomChoiceStringWithLength(length int, choice ChoiceString) (string, error) {
	var builder strings.Builder
	var source string
	switch choice {
	case ChoiceStringUpper:
		source = RandomUpperStringSource
	case ChoiceStringLower:
		source = RandomLowerStringSource
	default:
		return "", errors.New("invalid input")
	}
	randMax := int64(len(source)) - 1
	for i := 0; i < length; i++ {
		ran, err := rand.Int(rand.Reader, big.NewInt(randMax))
		if err != nil {
			return "", err
		}
		builder.WriteString(string(RandomStringSource[ran.Int64()]))
	}
	return builder.String(), nil
}

func (m *SyncMutexMap) LoadOrStore(key string) (*sync.Mutex, error) {
	if m.KeyNum <= 0 {
		return nil, errors.New("invalid key number")
	}
	hash := fnv.New32a()
	_, err := hash.Write([]byte(key))
	if err != nil {
		return nil, err
	}
	mapKeyNum := hash.Sum32() % uint32(m.KeyNum)
	mutex, _ := m.MutexMap.LoadOrStore(mapKeyNum, &sync.Mutex{})
	return mutex.(*sync.Mutex), nil
}

func ParseClientType(uaHeader string) ClientType {
	ua := useragent.Parse(uaHeader)
	if ua.Bot {
		return ClientTypeBot
	} else if ua.IsIOS() {
		return ClientTypeIOS
	} else if ua.IsAndroid() {
		return ClientTypeAndroid
	} else if ua.Desktop {
		return ClientTypeBrowser
	}
	return ClientTypeOther
}

func MakePaginateInfo(total uint, perPage uint, page uint) (*PaginateInfo, error) {
	if perPage < 1 || page < 1 {
		return nil, errors.New("invalid input")
	}
	pgInfo := &PaginateInfo{}
	totalPage := (total + perPage - 1) / perPage
	pgInfo.Total = total
	pgInfo.PerPage = perPage
	pgInfo.CurrentPage = page
	if page > 1 {
		pgInfo.PrevPage = page - 1
	} else {
		pgInfo.PrevPage = 1
	}
	if page >= totalPage {
		pgInfo.NextPage = page
	} else {
		pgInfo.NextPage = page + 1
	}
	pgInfo.FirstPage = 1
	pgInfo.LastPage = totalPage
	return pgInfo, nil
}

func RoundFloatWithPrecision(number float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(number*ratio) / ratio
}

func (m *SyncMutexMap) PeriodicClean() error {
	if m.KeyNum <= 0 {
		return errors.New("invalid key number")
	}
	ticker := time.NewTicker(time.Hour)
	go func(t *time.Ticker) {
		for range t.C {
			m.MutexMap.Range(func(key interface{}, value interface{}) bool {
				m.MutexMap.Delete(key)
				return true
			})
		}
	}(ticker)
	return nil
}
