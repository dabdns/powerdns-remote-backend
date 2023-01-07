package config

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConfig_Merge_Nil(t *testing.T) {
	var a, b, expected *Config
	a = &Config{}
	b = nil
	expected = &Config{}
	a.Merge(b)
	assert.Equal(t, a, expected)
}
func TestConfig_Merge_Empty(t *testing.T) {
	var a, b, expected *Config
	a = &Config{}
	b = &Config{}
	expected = &Config{}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConfig_Merge_01(t *testing.T) {
	var a, b, expected *Config
	c := ConnectorConfig{}
	d0 := DelegateConfig{}
	a = &Config{
		Connector: &c,
		Delegates: []*DelegateConfig{
			&d0,
		},
	}
	b = &Config{}
	expected = &Config{
		Connector: &c,
		Delegates: []*DelegateConfig{
			&d0,
		},
	}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConfig_Merge_02(t *testing.T) {
	var a, b, expected *Config
	c := ConnectorConfig{}
	s0a, s1a, s0b := "0a", "1a", "0b"
	d0a := DelegateConfig{
		Domain: &s0a,
	}
	d1a := DelegateConfig{
		Domain: &s1a,
	}
	d0b := DelegateConfig{
		Domain: &s0b,
	}
	a = &Config{
		Delegates: []*DelegateConfig{
			&d0a,
			&d1a,
		},
	}
	b = &Config{
		Connector: &c,
		Delegates: []*DelegateConfig{
			&d0b,
		},
	}
	expected = &Config{
		Connector: &c,
		Delegates: []*DelegateConfig{
			&d0b,
			&d1a,
		},
	}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConfig_Merge_03(t *testing.T) {
	var a, b, expected *Config
	cha, chb := "a", "b"
	ca := ConnectorConfig{
		Host: &cha,
	}
	cb := ConnectorConfig{
		Host: &chb,
	}
	a = &Config{
		Connector: &ca,
	}
	b = &Config{
		Connector: &cb,
	}
	expected = &Config{
		Connector: &cb,
	}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConnectorConfig_Merge_Nil(t *testing.T) {
	var a, b, expected *ConnectorConfig
	a = &ConnectorConfig{}
	b = nil
	expected = &ConnectorConfig{}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConnectorConfig_Merge_Empty(t *testing.T) {
	var a, b, expected *ConnectorConfig
	ty, h, p := "t", "h", uint16(1)
	a = &ConnectorConfig{
		Type: &ty,
		Host: &h,
		Port: &p,
	}
	b = &ConnectorConfig{}
	expected = &ConnectorConfig{
		Type: &ty,
		Host: &h,
		Port: &p,
	}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestConnectorConfig_Merge_00(t *testing.T) {
	var a, b, expected *ConnectorConfig
	ta, ha, pa := "ta", "ha", uint16(1)
	tb, hb, pb := "tb", "hb", uint16(2)
	a = &ConnectorConfig{
		Type: &ta,
		Host: &ha,
		Port: &pa,
	}
	b = &ConnectorConfig{
		Type: &tb,
		Host: &hb,
		Port: &pb,
	}
	expected = &ConnectorConfig{
		Type: &tb,
		Host: &hb,
		Port: &pb,
	}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestDelegateConfig_Merge_Nil(t *testing.T) {
	var a, b, expected *DelegateConfig
	a = &DelegateConfig{}
	b = nil
	expected = &DelegateConfig{}
	a.Merge(b)
	assert.Equal(t, a, expected)
}

func TestDelegateConfig_Merge_Empty(t *testing.T) {
	var a, b, expected *DelegateConfig
	a = &DelegateConfig{}
	b = &DelegateConfig{}
	expected = &DelegateConfig{}
	a.Merge(b)
	assert.Equal(t, a, expected)
}
