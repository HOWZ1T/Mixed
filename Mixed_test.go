package mixed

import "testing"

func TestVaryingDataTypes(t *testing.T) {
	lst := List{1, "a string", 3.23, 3.1456789, float32(12.12345)}

	if lst == nil {
		t.Errorf("mixed: could not handle varying data types")
	}
}

func TestHas(t *testing.T) {
	lst := List{1, "a string", 3.23, 3.1456789, float32(12.12345)}

	if lst.Has("a string") == false {
		t.Errorf("mixed: Has: expected true, got false for Has(\"a string\")")
	}

	if lst.Has(100) == true {
		t.Errorf("mixed: Has: expected false, got true for Has(100)")
	}
}

func TestAppend(t *testing.T) {
	lst := List{1, 2, 3}
	lst.Append(4)

	if lst.Has(4) == false {
		t.Errorf("mixed: Append failed, expected: %v, got: %v", List{1, 2, 3, 4}, lst)
	}
}

func TestExtend(t *testing.T) {
	lst := List{1}
	lst.Extend(List{2, 3, 4})

	if lst.Has(4) == false {
		t.Errorf("mixed: Append failed, expected: %v, got: %v", List{1, 2, 3, 4}, lst)
	}
}

func TestEquals(t *testing.T) {
	lstA := List{1, "two", 3.33}
	lstB := List{1, "two", 3.33}

	if lstA.Equals(lstB) == false {
		t.Errorf("mixed: Equals, expected: true, got: false")
	}

	lstB[1] = "TWO"
	if lstA.Equals(lstB) == true {
		t.Errorf("mixed: Equals, expected: false, got: true")
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	lstA := List{1, "two", 3.33, "FOUR"}
	lstB := List{1, "two", 3.33, "four"}

	if lstA.EqualsIgnoreCase(lstB) == false {
		t.Errorf("mixed: Equals, expected: true, got: false")
	}

	lstB[3] = "fIvE"
	if lstA.EqualsIgnoreCase(lstB) == true {
		t.Errorf("mixed: Equals, expected: false, got: true")
	}
}

func TestInsert(t *testing.T) {
	lst := List{1, 3, 4}
	lst.Insert(1, 2)

	if lst.Equals(List{1, 2, 3, 4}) == false {
		t.Errorf("mixed: Insert, expected: %v, got: %v", List{1, 2, 3, 4}, lst)
	}
}

func TestPop(t *testing.T) {
	lst := List{1, 5, 2, 3}
	_, poppedValue := lst.Pop(1)

	if poppedValue != 5 {
		t.Errorf("mixed: Pop, expected: %v, got: %v", 5, poppedValue)
	}

	if lst.Equals(List{1, 2, 3}) == false {
		t.Errorf("mixed: Pop, expected: %v, got: %v", List{1, 2, 3}, lst)
	}
}

func TestClear(t *testing.T) {
	lst := List{"I", "am", "not", "clear"}
	lst.Clear()

	if lst.Equals(List{}) == false {
		t.Errorf("mixed: Clear, expected: %v, got: %v", List{}, lst)
	}
}

func TestIndex(t *testing.T) {
	lst := List{"zero", "one", "two", "three", "four"}
	err, index := lst.IndexOf("two")

	if err != nil {
		t.Errorf("mixed: Index, expected: no error, got: error: %s", err)
	}

	if index != 2 {
		t.Errorf("mixed: Index, expected: 2, got: %d", index)
	}
}

func TestSize(t *testing.T) {
	lst := List{1, 2, 3, 4, 5}
	size := lst.Size()

	if size != 5 {
		t.Errorf("mixed: Size, expected: 5, got: %d", size)
	}
}

func TestReverse(t *testing.T) {
	lst := List{1, 2, 3, 4}
	lst1 := List{1, 2, 3}

	lst.Reverse()
	lst1.Reverse()

	if lst.Equals(List{4, 3, 2, 1}) == false {
		t.Errorf("mixed: Reverse, expected: %d, got: %d", List{4, 3, 2, 1}, lst)
	}

	if lst1.Equals(List{3, 2, 1}) == false {
		t.Errorf("mixed: Reverse, expected: %d, got: %d", List{3, 2, 1}, lst1)
	}
}

func TestCopy(t *testing.T) {
	lst := List{1, 2, 3, 4}
	lst1 := lst.Copy()

	//Modifying original lst to ensure the copied list is not referencing the original list
	lst.Clear()

	if lst1.Equals(List{1, 2, 3, 4}) == false {
		t.Errorf("mixed: Copy, expected: %d, got: %d", List{1, 2, 3, 4}, lst)
	}
}