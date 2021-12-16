package acmachine

import "testing"

func TestAcMachineMatch(t *testing.T) {
	m := NewMachine(SplitString, CombineString)
	m.AddPattern("彩票")
	m.AddPattern("博彩")
	m.AddPattern("广告")
	m.Build()
	ms := m.Match("我中了一个彩票")
	t.Logf("match result %+v", ms)
	if len(ms) != 1 || ms[0].Pattern.(string) != "彩票" {
		t.Errorf("match failed. expect 彩票 but %+v", ms)
	}
	ms = m.Match("彩票广告")
	t.Logf("match result %+v", ms)
	if len(ms) != 2 || ms[0].Pattern.(string) != "彩票" || ms[1].Pattern.(string) != "广告" {
		t.Errorf("match failed. expect 彩票 广告 but %+v", ms)
	}

}