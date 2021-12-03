package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkField_GetterInto(b *testing.B) {
	intHolder := new(int)
	stringHolder := new(string)
	float32Holder := new(float32)
	timeHolder := new(time.Time)
	var id int
	var name string
	var val float32
	var ts time.Time
	ptr := Addr(_accBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_AcciDField.GetInto(ptr, intHolder)
		id = *intHolder

		_AccNameField.GetInto(ptr, stringHolder)
		name = *stringHolder

		_AccValField.GetInto(ptr, float32Holder)
		val = *float32Holder

		_TimeValField.GetInto(ptr, timeHolder)
		ts = *timeHolder
	}
	assert.EqualValues(b, _accBenchInstance.ID, id)
	assert.EqualValues(b, _accBenchInstance.Name, name)
	assert.EqualValues(b, _accBenchInstance.Val, val)
	assert.EqualValues(b, _accBenchInstance.Time, ts)
}
