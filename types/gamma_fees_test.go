package types

import (
	"encoding/json"
	"testing"
)

func TestGammaMarket_DecodesFeeSchedule(t *testing.T) {
	const withFees = `{"question":"Ethereum Up or Down - October 22, 12:30PM-12:45PM ET","feesEnabled":true,"feeType":"crypto_fees_v2","makerBaseFee":1000,"takerBaseFee":1000,"feeSchedule":{"exponent":1,"rate":0.07,"takerOnly":true,"rebateRate":0.2}}`

	var m GammaMarket
	if err := json.Unmarshal([]byte(withFees), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if m.FeesEnabled == nil || *m.FeesEnabled != true {
		t.Fatalf("FeesEnabled = %v, want true", m.FeesEnabled)
	}
	if m.FeeType == nil || *m.FeeType != "crypto_fees_v2" {
		t.Fatalf("FeeType = %v, want crypto_fees_v2", m.FeeType)
	}
	if m.FeeSchedule == nil {
		t.Fatal("FeeSchedule = nil, want non-nil")
	}
	if m.FeeSchedule.Rate != 0.07 {
		t.Errorf("FeeSchedule.Rate = %v, want 0.07", m.FeeSchedule.Rate)
	}
	if m.FeeSchedule.Exponent != 1 {
		t.Errorf("FeeSchedule.Exponent = %v, want 1", m.FeeSchedule.Exponent)
	}
	if m.FeeSchedule.RebateRate != 0.2 {
		t.Errorf("FeeSchedule.RebateRate = %v, want 0.2", m.FeeSchedule.RebateRate)
	}
	if !m.FeeSchedule.TakerOnly {
		t.Errorf("FeeSchedule.TakerOnly = %v, want true", m.FeeSchedule.TakerOnly)
	}

	const withoutFees = `{"question":"Some other market","makerBaseFee":1000,"takerBaseFee":1000}`
	var m2 GammaMarket
	if err := json.Unmarshal([]byte(withoutFees), &m2); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m2.FeesEnabled != nil {
		t.Errorf("FeesEnabled = %v, want nil", m2.FeesEnabled)
	}
	if m2.FeeType != nil {
		t.Errorf("FeeType = %v, want nil", m2.FeeType)
	}
	if m2.FeeSchedule != nil {
		t.Errorf("FeeSchedule = %v, want nil", m2.FeeSchedule)
	}
}
