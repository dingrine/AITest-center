package redis

import (
	"examCenter/pkg/json"
	"examCenter/pkg/redis"
	"fmt"
)

type ExamQA struct {
	QAInfo string
}

func (m *ExamQA) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(m)
}

func (m *ExamQA) UnmarshalBinary(data []byte) error {
	if m == nil {
		return fmt.Errorf("receiver is nil")
	}
	return json.JSON().Unmarshal(data, m)
}

const (
	KEY_EXAM_QA_INFO = "exam:center:qa:info:%d"
)

func tokenImageKey(oeType int64) string {
	return fmt.Sprintf(KEY_EXAM_QA_INFO, oeType)
}
func NewNftImage(client *redis.Client) *tokenCache[*ExamQA] {
	cache := newTokenCache[*ExamQA](client, tokenImageKey)
	return cache
}
