package common

import "encoding/json"

// RuleMap 群组规则的字典，匹配规则=>回复内容
type RuleMap map[string]string

// BanRuleMap 禁言群组规则的字典，
type BanRuleMap map[string]string

var (
	// AllGroupRules 所有群组的规则字典
	AllGroupRules = make(map[int64]RuleMap)
	// AllGroupBanRules 所有群组禁言的规则字典
	AllGroupBanRules = make(map[int64]BanRuleMap)
	// AllGroupId 目前服务的所有群组的id
	AllGroupId []int64
)

func (rm RuleMap) String() string {
	s, err := json.Marshal(rm)
	if err != nil {
		return ""
	}
	return string(s)
}

func (rm BanRuleMap) String() string {
	s, err := json.Marshal(rm)
	if err != nil {
		return ""
	}
	return string(s)
}

// Json2kvs 将json字符串转化为规则字典
func Json2kvs(rulesJson string) map[string]string {
	tkvs := make(map[string]string)
	_ = json.Unmarshal([]byte(rulesJson), &tkvs)
	return tkvs
}

// AddNewGroup 在内存中添加新群组的条目
func AddNewGroup(gid int64) {
	AllGroupId = append(AllGroupId, gid)
	AllGroupRules[gid] = make(RuleMap)
}

// AddNewGroupBan 在内存中添加新群组的条目
func AddNewGroupBan(gid int64) {
	AllGroupId = append(AllGroupId, gid)
	AllGroupBanRules[gid] = make(BanRuleMap)
}
