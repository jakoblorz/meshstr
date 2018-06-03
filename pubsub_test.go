package brokerutil

import (
	"testing"

	"github.com/jakoblorz/brokerutil/driver"
	"github.com/jakoblorz/brokerutil/driver/loopback"
)

type dynamicDriverScaffold struct {
	driverType driver.PubSubDriverType
}

func (d dynamicDriverScaffold) GetDriverType() driver.PubSubDriverType {
	return d.driverType
}

func (d dynamicDriverScaffold) CloseStream() error {
	return nil
}

func (d dynamicDriverScaffold) OpenStream() error {
	return nil
}

func TestNewPubSubFromDriver(t *testing.T) {

	mt, err := loopback.NewMultiThreadDriver()
	if err != nil {
		t.Error(err)
	}

	st, err := loopback.NewSingleThreadDriver()
	if err != nil {
		t.Error(err)
	}

	t.Run("should return PubSub using proper driver wrapper", func(t *testing.T) {
		psMt, err := NewPubSubFromDriver(mt)
		if err != nil {
			t.Error(err)
		}

		if _, ok := psMt.(multiThreadPubSubDriverWrapper); !ok {
			t.Error("NewPubSubFromDriver() did not create proper driver multiThreadPubSubDriverWrapper")
		}

		psSt, err := NewPubSubFromDriver(st)
		if err != nil {
			t.Error(err)
		}

		if _, ok := psSt.(singleThreadPubSubDriverWrapper); !ok {
			t.Error("NewPubSubFromDriver() did not create proper driver singleThreadPubSubDriverWrapper")
		}

		_, err = NewPubSubFromDriver(dynamicDriverScaffold{
			driverType: driver.PubSubDriverType(3),
		})

		if err == nil {
			t.Error("NewPubSubFromDriver() did not return error when providing wrong driver type")
		}
	})
}

func TestNewPubSubFromMultiThreadDriver(t *testing.T) {

	mt, err := loopback.NewMultiThreadDriver()
	if err != nil {
		t.Error(err)
	}

	t.Run("should return PubSub using multi thread driver wrapper", func(t *testing.T) {
		psMt, err := NewPubSubFromMultiThreadDriver(mt)
		if err != nil {
			t.Error(err)
		}

		if _, ok := psMt.(multiThreadPubSubDriverWrapper); !ok {
			t.Error("NewPubSubFromMultiThreadDriver() did not use proper driver wrapper")
		}
	})
}

func TestNewPubSubFromSingleThreadDriver(t *testing.T) {

	st, err := loopback.NewSingleThreadDriver()
	if err != nil {
		t.Error(err)
	}

	t.Run("should return PubSub using single thread driver wrapper", func(t *testing.T) {
		stMt, err := NewPubSubFromSingleThreadDriver(st)
		if err != nil {
			t.Error(err)
		}

		if _, ok := stMt.(singleThreadPubSubDriverWrapper); !ok {
			t.Error("NewPubSubFromSingleThreadDriver() did not use proper driver wrapper")
		}
	})
}
