package workers

import (
	"testing"
)

func TestRegisterConcurency(t *testing.T) {
	w := NewWorkers("stst", 1000)

	for i := 0; i < 1000; i++ {
		param := map[string]interface{}{
			"i": i,
		}

		go func() {
			id, err := w.Register(param)
			if err != nil {
				t.Error(
					"Register iterator ", i,
					"for <", param, ">",
					"got <", err, ">",
				)
			}

			if err := w.UnRegister(id); err != nil {
				t.Error(
					"UnRegister iterator ", i,
					"for <", param, ">",
					"got <", err, ">",
				)
			}
		}()
	}
}

func TestRegister(t *testing.T) {
	w := NewWorkers("stst", 2)

	if _, err := w.Register(123); err != nil {
		t.Error(
			"Expected <nil>",
			"got <", err, ">",
		)
	}

	if _, err := w.Register(123); err != ErrWorkerRegistered {
		t.Error(
			"Expected <", ErrWorkerRegistered, ">",
			"got <", err, ">",
		)
	}

	if _, err := w.Register(111); err != nil {
		t.Error(
			"Expected <nil>",
			"got <", err, ">",
		)
	}

	if _, err := w.Register(222); err != ErrWorkersOutLimits {
		t.Error(
			"Expected <", ErrWorkersOutLimits, ">",
			"got <", err, ">",
		)
	}
}

func TestUnRegister(t *testing.T) {
	w := NewWorkers("stst", 1)

	id, _ := w.Register(111)
	if err := w.UnRegister(id); err != nil {
		t.Error(
			"Expected <nil>",
			"got <", err, ">",
		)
	}

	if err := w.UnRegister("lala-bebe-123"); err != ErrWorkerNotRegistered {
		t.Error(
			"Expected <", ErrWorkerNotRegistered, ">",
			"got <", err, ">",
		)
	}
}

func TestGetCount(t *testing.T) {
	type testpair struct {
		value int64
		fn    func()
	}

	w := NewWorkers("stst", 2)

	tests := []testpair{
		{
			value: 0,
			fn: func() {
				// Nothing
			},
		},
		{
			value: 1,
			fn: func() {
				_, _ = w.Register(111)
			},
		},
		{
			value: 2,
			fn: func() {
				_, _ = w.Register(222)
			},
		},
	}

	for i, pair := range tests {
		pair.fn()
		v := w.GetCount()

		if v != pair.value {
			t.Error(
				"Iterator ", i,
				"expected <", pair.value, ">",
				"got <", v, ">",
			)
		}
	}
}

func TestGetWorker(t *testing.T) {
	w := NewWorkers("stst", 2)

	id, _ := w.Register(111)

	if r, _ := w.GetWorker(id); r.Param != 111 {
		t.Error(
			"Expected <111>",
			"got <", r, ">",
		)
	}

	if _, err := w.GetWorker("lala-bebe-123"); err != ErrWorkerNotRegistered {
		t.Error(
			"Eexpected <", ErrWorkerNotRegistered, ">",
			"got <", err, ">",
		)
	}
}

func TestGetWorkers(t *testing.T) {
	w := NewWorkers("stst", 3)

	if _, err := w.GetWorkers(); err != ErrWorkersNotFound {
		t.Error(
			"Eexpected <", ErrWorkersNotFound, ">",
			"got <", err, ">",
		)
	}

	_, _ = w.Register(111)
	_, _ = w.Register(222)

	wl, _ := w.GetWorkers()

	if l := len(wl); l != 2 {
		t.Error(
			"Eexpected <2>",
			"got <", l, ">",
		)
	}

	_, _ = w.Register(333)
	wl, _ = w.GetWorkers()

	if r := wl[2].Param; r != 333 {
		t.Error(
			"Expected <333>",
			"got <", r, ">",
		)
	}
}
