package compactlist

// use case fro link list
// easy to append data

type CompactListVar struct {
	Value  *interface{}
	next   int // int next
	before int
}

type CompactList struct {
	List  []CompactListVar
	start int
	end   int
	empty int
}

func Newcompact(Value interface{}, len int) CompactList {
	list := make([]CompactListVar, len)
	return CompactList{List: append(list, CompactListVar{Value: &Value, next: 0, before: 0}), start: 0, empty: 0, end: 0}
}

//when add item, three solution
// first add item by search
// last item should return to
func (m *CompactList) Add(value *interface{}) {
	m.addBySearch(value, m.start)
}

func (m *CompactList) addBySearch(value *interface{}, index int) {

	if m.List[m.empty].Value == nil {
		m.addItemInIndex(m.end, m.empty, m.start, value)
		m.end = m.empty
		return
	}

	emptynumber := make(chan int)
	go m.findempty(emptynumber)

	next, before := m.List[index].next, m.List[index].next

	for i := 0; m.List[next].Value != nil && i < len(m.List); i++ {
		if m.List[next].next == next || m.List[next].next == m.start {
			m.addItemInIndex(before, <-emptynumber, next, value)
			return
		}
		before = next
		next = m.List[next].next
	}

	if m.List[next].Value != nil {
		m.start = m.List[m.start].next
		m.addItemInIndex(before, m.start, m.start, value)
		return
	}
	// if value is null
	m.addItemInIndex(before, <-emptynumber, m.start, value)
}

func (m *CompactList) addItemInIndex(before int, index int, next int, value *interface{}) {
	m.List[index].Value = value
	m.List[index].next = next
	m.List[index].before = before
	m.List[next].next = next
}

// append item after the offset
func (m *CompactList) Append(value *interface{}, offset int) {
	emptynumber := make(chan int)
	go m.findempty(emptynumber)
	next := m.List[offset].next
	empty := <-emptynumber
	m.List[offset].next = empty
	m.List[empty].Value = value
	m.List[empty].next = next
}

func (m CompactList) findempty(empty chan int) {
	for i := 0; i < len(m.List); i++ {
		if m.List[i].Value == nil {
			empty <- i
			return
		}
	}
	//if list didn't have empty space
}

func (m *CompactList) RemoveBySearch(value *interface{}) {

	next, before := m.start, m.start
	for i := 0; m.List[next].Value != nil || m.List[next].Value != value || i < len(m.List); {
		before = next
		next = m.List[next].next
		i++
	}
	if m.List[next].Value == value {
		m.List[before].next = m.List[next].next
		m.List[next].Value = nil
		m.List[next].next = m.start
	}
}

// Remove item by index
func (m *CompactList) RemoveByIndex(index int) {
	b := m.List[index].before
	m.List[b].next = m.List[index].next
	m.List[m.List[b].next].before = b
	m.List[index].Value = nil
	m.empty = index
	if m.end == index {
		m.end = b
	}
}
