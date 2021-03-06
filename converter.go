package katsubushi

import "time"

func ToTime(id uint64) time.Time {
	ts := id >> (WorkerIDBits + SequenceBits)
	d := time.Duration(int64(ts) * int64(time.Millisecond))
	return Epoch.Add(d)
}

func ToID(t time.Time) uint64 {
	d := t.Sub(Epoch)
	ts := uint64(d.Nanoseconds()) / uint64(time.Millisecond)
	return ts << (WorkerIDBits + SequenceBits)
}

func Dump(id uint64) (t time.Time, workerID uint64, sequense uint64) {
	workerID = (id & (workerIDMask << SequenceBits)) >> SequenceBits
	sequense = id & sequenseMask
	return ToTime(id), workerID, sequense
}
