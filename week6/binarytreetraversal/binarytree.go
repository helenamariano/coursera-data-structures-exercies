package main

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

var isBstMin int

type BT struct {
	keys  []int
	left  []int
	right []int
}

func (bt *BT) InOrder() []int {
	return bt.inOrderRecursive(0)
}

func (bt *BT) inOrderRecursive(keyIndex int) (result []int) {
	if keyIndex != -1 {
		result = append(result, bt.inOrderRecursive(bt.left[keyIndex])...)
		result = append(result, bt.keys[keyIndex])
		result = append(result, bt.inOrderRecursive(bt.right[keyIndex])...)
	}

	return result
}

func (bt *BT) PreOrder() []int {
	return bt.preOrderRecursive(0)
}

func (bt *BT) preOrderRecursive(keyIndex int) (result []int) {
	if keyIndex != -1 {
		result = append(result, bt.keys[keyIndex])
		result = append(result, bt.preOrderRecursive(bt.left[keyIndex])...)
		result = append(result, bt.preOrderRecursive(bt.right[keyIndex])...)
	}

	return result
}

func (bt *BT) PostOrder() []int {
	return bt.postOrderRecursive(0)
}

func (bt *BT) postOrderRecursive(keyIndex int) (result []int) {
	if keyIndex != -1 {
		result = append(result, bt.postOrderRecursive(bt.left[keyIndex])...)
		result = append(result, bt.postOrderRecursive(bt.right[keyIndex])...)
		result = append(result, bt.keys[keyIndex])
	}

	return result
}

func (bt *BT) IsBst() bool {
	if len(bt.keys) == 0 {
		return true
	}
	isBstMin = minInt
	return bt.isBstRecursive(0)
}

func (bt *BT) isBstRecursive(keyIndex int) bool {
	if keyIndex != -1 {
		if !bt.isBstRecursive(bt.left[keyIndex]) {
			return false
		}

		if bt.keys[keyIndex] <= isBstMin {
			return false
		}

		isBstMin = bt.keys[keyIndex]

		if !bt.isBstRecursive(bt.right[keyIndex]) {
			return false
		}
	}
	return true
}

func NewBinaryTree(keys, left, right []int) *BT {
	return &BT{keys: keys, left: left, right: right}
}
