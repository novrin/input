package input

import "testing"

func TestIsMember(t *testing.T) {
	ss := []string{"a", "b", "c"}
	if IsMember("z", ss) {
		t.Error(`IsMember("z", ...) should be false`)
	}
	if !IsMember("a", ss) {
		t.Error(`IsMember("a", ...) should be true`)
	}
}

func TestIsInCharLimit(t *testing.T) {
	if IsInCharLimit("abcd", 1, 3) {
		t.Error(`IsInCharLimit("abcd", 1, 3) should be false`)
	}
	if !IsInCharLimit("abc", 1, 3) {
		t.Error(`IsInCharLimit("abc", 1, 3) should be true`)
	}
}

func TestIsTime(t *testing.T) {
	layout := "2006-01-02"
	if IsTime("not-a-date", layout) {
		t.Error(`IsTime("not-a-date", "2006-01-02") should be false`)
	}
	if !IsTime("2023-05-10", layout) {
		t.Error(`IsTime("2023-05-10", "2006-01-02") should be true`)
	}
}

func TestIsTimePastFuture(t *testing.T) {
	layout := "2006-01-02"
	invalid := "not-a-date"
	past := "1999-01-01"
	future := "2999-01-01"
	if IsTimePast(future, layout) {
		t.Error("IsTimePast(future) should be false")
	}
	if IsTimePast(past, invalid) {
		t.Error("IsTimePast(invalid) should be false")
	}
	if !IsTimePast(past, layout) {
		t.Error("IsTimePast(past) should be true")
	}

	if IsTimeFuture(past, layout) {
		t.Error("IsTimeFuture(past) should be false")
	}
	if IsTimeFuture(future, invalid) {
		t.Error("IsTimeFuture(invalid) should be false")
	}
	if !IsTimeFuture(future, layout) {
		t.Error("IsTimeFuture(future) should be true")
	}
}

func TestIsURL(t *testing.T) {
	if !IsURL("https://example.com") {
		t.Error(`IsURL("https://example.com") should be true`)
	}
	if IsURL("not a url") {
		t.Error(`IsURL("not a url") should be false`)
	}
}

func TestIsBool(t *testing.T) {
	if !IsBool("true") || !IsBool("false") || !IsBool("0") || !IsBool("1") {
		t.Error("IsBool should accept 'true', 'false', '0', '1'")
	}
	if IsBool("maybe") {
		t.Error("IsBool('maybe') should be false")
	}
}

func TestIsInt(t *testing.T) {
	if !IsInt("42", 10, 64) {
		t.Error("IsInt('42', 10, 64) should be true")
	}
	if IsInt("3.14", 10, 64) {
		t.Error("IsInt('3.14', 10, 64) should be false")
	}
}

func TestIsUint(t *testing.T) {
	if !IsUint("42", 10, 64) {
		t.Error("IsUint('42', 10, 64) should be true")
	}
	if IsUint("-42", 10, 64) {
		t.Error("IsUint('-42', 10, 64) should be false")
	}
}

func TestIsFloat(t *testing.T) {
	if !IsFloat("3.14", 64) {
		t.Error("IsFloat('3.14', 64) should be true")
	}
	if IsFloat("notfloat", 64) {
		t.Error("IsFloat('notfloat', 64) should be false")
	}
}

func TestLogCheckAndOK(t *testing.T) {
	var log Log
	log.Check("name",
		Rule{OK: true, Message: "should not appear"},
		Rule{OK: false, Message: "required"},
		Rule{OK: false, Message: "too short"},
	)
	if log.OK() {
		t.Error("Log should not be OK after failed checks")
	}
	if len(log.Errors["name"]) != 2 {
		t.Errorf("Expected 2 errors for field 'name', got %d", len(log.Errors["name"]))
	}

	empty := Log{}
	if !empty.OK() {
		t.Error("Empty Log should be OK")
	}
}
