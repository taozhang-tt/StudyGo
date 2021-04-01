package singleton

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(t *testing.T) {
	Convey("get instance", t, func() {
		So(GetInstance(), ShouldEqual, GetInstance())
	})
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	Convey("get lazy instance", t, func() {
		So(GetLazyInstance(), ShouldEqual, GetLazyInstance())
	})
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
