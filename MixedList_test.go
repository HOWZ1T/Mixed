package MixedList

import "testing"

func TestVaryingDataTypes(t *testing.T) {
	lst := MixedList{1, "a string", 3.23, 3.1456789, float32(12.12345)}

	if lst == nil {
		t.Errorf("MixedList: could not handle varying data types")
	}
}

func TestHas(t *testing.T) {
	lst := MixedList{1, "a string", 3.23, 3.1456789, float32(12.12345)}

	if lst.Has("a string") == false {
		t.Errorf("MixedList: Has: expected true, got false for Has(\"a string\")")
	}

	if lst.Has(100) == true {
		t.Errorf("MixedList: Has: expected false, got true for Has(100)")
	}
}

func TestAppend(t *testing.T) {
	lst := MixedList{1, 2, 3}
	lst.Append(4)

	if lst.Has(4) == false {
		t.Errorf("MixedList: Append failed, expected: %v, got: %v", MixedList{1, 2, 3, 4}, lst)
	}
}

func TestExtend(t *testing.T) {
	lst := MixedList{1}
	lst.Extend(MixedList{2, 3, 4})

	if lst.Has(4) == false {
		t.Errorf("MixedList: Append failed, expected: %v, got: %v", MixedList{1, 2, 3, 4}, lst)
	}
}

func TestEquals(t *testing.T) {
	lstA := MixedList{1, "two", 3.33}
	lstB := MixedList{1, "two", 3.33}

	if lstA.Equals(lstB) == false {
		t.Errorf("MixedList: Equals, expected: true, got: false")
	}

	lstB[1] = "TWO"
	if lstA.Equals(lstB) == true {
		t.Errorf("MixedList: Equals, expected: false, got: true")
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	lstA := MixedList{1, "two", 3.33, "FOUR"}
	lstB := MixedList{1, "two", 3.33, "four"}

	if lstA.EqualsIgnoreCase(lstB) == false {
		t.Errorf("MixedList: Equals, expected: true, got: false")
	}

	lstB[3] = "fIvE"
	if lstA.EqualsIgnoreCase(lstB) == true {
		t.Errorf("MixedList: Equals, expected: false, got: true")
	}
}

func TestInsert(t *testing.T) {
	lst := MixedList{1, 3, 4}
	lst.Insert(1, 2)

	if lst.Equals(MixedList{1, 2, 3, 4}) == false {
		t.Errorf("MixedList: Insert, expected: %v, got: %v", MixedList{1, 2, 3, 4}, lst)
	}
}

func TestPop(t *testing.T) {
	lst := MixedList{1, 5, 2, 3}
	_, poppedValue := lst.Pop(1)

	if poppedValue != 5 {
		t.Errorf("MixedList: Pop, expected: %v, got: %v", 5, poppedValue)
	}

	if lst.Equals(MixedList{1, 2, 3}) == false {
		t.Errorf("MixedList: Pop, expected: %v, got: %v", MixedList{1, 2, 3}, lst)
	}
}

func TestClear(t *testing.T) {
	lst := MixedList{"I", "am", "not", "clear"}
	lst.Clear()

	if lst.Equals(MixedList{}) == false {
		t.Errorf("MixedList: Clear, expected: %v, got: %v", MixedList{}, lst)
	}
}

func TestIndex(t *testing.T) {
	lst := MixedList{"zero", "one", "two", "three", "four"}
	err, index := lst.Index("two")

	if err != nil {
		t.Errorf("MixedList: Index, expected: no error, got: error: %s", err)
	}

	if index != 2 {
		t.Errorf("MixedList: Index, expected: 2, got: %d", index)
	}
}

func TestSize(t *testing.T) {
	lst := MixedList{1, 2, 3, 4, 5}
	size := lst.Size()

	if size != 5 {
		t.Errorf("MixedList: Size, expected: 5, got: %d", size)
	}
}

func TestReverse(t *testing.T) {
	lst := MixedList{1, 2, 3, 4}
	lst1 := MixedList{1, 2, 3}

	lst.Reverse()
	lst1.Reverse()

	if lst.Equals(MixedList{4, 3, 2, 1}) == false {
		t.Errorf("MixedList: Reverse, expected: %d, got: %d", MixedList{4, 3, 2, 1}, lst)
	}

	if lst1.Equals(MixedList{3, 2, 1}) == false {
		t.Errorf("MixedList: Reverse, expected: %d, got: %d", MixedList{3, 2, 1}, lst1)
	}
}

func TestCopy(t *testing.T) {
	lst := MixedList{1, 2, 3, 4}
	lst1 := lst.Copy()

	//Modifying original lst to ensure the copied list is not referencing the original list
	lst.Clear()

	if lst1.Equals(MixedList{1, 2, 3, 4}) == false {
		t.Errorf("MixedList: Copy, expected: %d, got: %d", MixedList{1, 2, 3, 4}, lst)
	}
}