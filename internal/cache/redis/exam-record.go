package redis

import (
	"examCenter/pkg/json"
	"examCenter/pkg/redis"
	"fmt"
)

type Exam struct {
	// ExamStatus
}

func (m *Exam) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(m)
}

func (m *Exam) UnmarshalBinary(data []byte) error {
	if m == nil {
		return fmt.Errorf("receiver is nil")
	}
	return json.JSON().Unmarshal(data, m)
}

const (
	KEY_EXAM_INFO = "exam:center:exam:info:%d"
)

func examInfoKey(eiType int64) string {
	return fmt.Sprintf(KEY_EXAM_INFO, eiType)
}
func NewExamInfo(client *redis.Client) *tokenCache[*Exam] {
	cache := newTokenCache[*Exam](client, examInfoKey)
	return cache
}
