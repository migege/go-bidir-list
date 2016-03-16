package list

type List struct {
	Head *Item
	Tail *Item
	Map  map[string]*Item
}

type Item struct {
	ItemId string
	Money  uint
	Prev   *Item
	Next   *Item
}

func NewList() *List {
	l := new(List)
	l.Map = make(map[string]*Item)
	return l
}

func (l *List) adjust(item *Item) {
	if len(l.Map) == 1 {
		l.Head = item
		l.Tail = item
		return
	}

	t := item.Prev
	for t != nil && t.Money < item.Money {
		t = t.Prev
	}

	if t != nil {
		if t.Next != item {
			itemNext := item.Next
			itemPrev := item.Prev
			tNext := t.Next

			itemPrev.Next = itemNext
			if itemNext != nil {
				itemNext.Prev = itemPrev
			} else {
				l.Tail = itemPrev
			}
			t.Next = item
			item.Next = tNext
			item.Prev = t
			if tNext != nil {
				tNext.Prev = item
			} else {
				l.Tail = item
			}
		}
	} else {
		if l.Head != item {
			itemNext := item.Next
			itemPrev := item.Prev

			itemPrev.Next = itemNext
			if itemNext != nil {
				itemNext.Prev = itemPrev
			} else {
				l.Tail = itemPrev
			}
			item.Next = l.Head
			item.Prev = nil
			l.Head.Prev = item
			l.Head = item
		}
	}
}

func (l *List) AddMoney(id string, money uint) bool {
	if l.Map[id] != nil {
		item := l.Map[id]
		item.Money += money
		l.adjust(item)
	} else {
		item := &Item{ItemId: id, Money: money, Prev: l.Tail}
		l.Map[id] = item
		l.adjust(item)
	}
	return true
}

func (l *List) Top(n int) []*Item {
	i := 0
	t := l.Head
	var res []*Item
	for t != nil && i < n {
		res = append(res, t)
		t = t.Next
		i++
	}
	return res
}

func (l *List) CheckList() string {
	if l.Head == nil {
		return "Nil Head"
	}

	if l.Head.Prev != nil || l.Tail.Next != nil {
		return "Non-nil Head.Prev or Non-nil Tail.Next"
	}

	len_from_head := 0
	t := l.Head
	v := uint(4294967295)
	for t != nil {
		if t.Money > v {
			return "Head: Money Order Error"
		}
		v = t.Money
		t = t.Next
		len_from_head++
	}

	len_from_tail := 0
	t = l.Tail
	v = 0
	for t != nil {
		if t.Money < v {
			return "Tail: Money Order Error"
		}
		v = t.Money
		t = t.Prev
		len_from_tail++
	}

	if len_from_head != len(l.Map) {
		return "Incorrect Len from Head"
	}

	if len_from_tail != len(l.Map) {
		return "Incorrect Len from Tail"
	}

	if len_from_head != len_from_tail {
		return "LenHead != LenTail"
	}

	return "OK"
}
