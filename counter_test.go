package metrics

import (
	"testing"

	dto "github.com/prometheus/client_model/go"
)

func TestCounter(t *testing.T) {
	expected := struct {
		Name  string
		Help  string
		Type  dto.MetricType
		Value float64
	}{
		"test_counter_counter_total",
		"count all the things",
		dto.MetricType_COUNTER,
		5,
	}

	ns := NewNamespace("test", "counter", nil)

	c := ns.NewCounter("counter", "count all the things")
	c.Inc(5)

	Register(ns)
	defer Deregister(ns)

	mfs, err := Gather()
	if err != nil {
		t.Fatal(err)
	}

	if len(mfs) != 1 {
		t.Fatalf("expected one metric family but got %d", len(mfs))
	}

	family := mfs[0]
	if *family.Name != expected.Name {
		t.Fatalf("expected name `%s` but got `%s`", expected.Name, *family.Name)
	}

	if *family.Help != expected.Help {
		t.Fatalf("expected help `%s` but got `%s`", expected.Help, *family.Help)
	}

	if *family.Type != expected.Type {
		t.Fatalf("expected type `%d` but got `%d`", expected.Type, *family.Type)
	}

	if len(family.Metric) != 1 {
		t.Fatalf("expected one metric but got %d", len(family.Metric))
	}

	metric := family.Metric[0]
	value := metric.GetCounter().GetValue()
	if value != expected.Value {
		t.Fatalf("expected counter value %f but got %f", expected.Value, value)
	}
}
