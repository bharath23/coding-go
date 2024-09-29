package educative0010

type FreqStack struct {
	freqMap    map[int]int
	freqBucket map[int][]int
	maxFreq    int
}

func Constructor() *FreqStack {
	s := &FreqStack{}
	s.freqMap = make(map[int]int)
	s.freqBucket = make(map[int][]int)
	s.maxFreq = 0

	return s
}

func (s *FreqStack) Push(val int) {
	s.freqMap[val]++
	curFreq := s.freqMap[val]
	s.maxFreq = max(s.maxFreq, curFreq)
	bb := s.freqBucket[curFreq]
	bb = append(bb, val)
	s.freqBucket[curFreq] = bb
}

func (s *FreqStack) Pop() int {
	bb := s.freqBucket[s.maxFreq]
	val := bb[len(bb)-1]
	bb = bb[:len(bb)-1]
	s.freqBucket[s.maxFreq] = bb
	s.freqMap[val]--
	for len(s.freqBucket[s.maxFreq]) == 0 && s.maxFreq > 0 {
		s.maxFreq--
	}
	return val
}
