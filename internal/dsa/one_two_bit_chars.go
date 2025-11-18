package dsa

//
//1-bit and 2-bit Characters
//We have two special characters:

//The first character can be represented by one bit 0.
//The second character can be represented by two bits (10 or 11).
//Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.

//Example 1:
//
//Input: bits = [1,0,0]
//Output: true
//Explanation: The only way to decode it is two-bit character and one-bit character.
//So the last character is one-bit character.
//Example 2:
//
//Input: bits = [1,1,1,0]
//Output: false
//Explanation: The only way to decode it is two-bit character and two-bit character.
//So the last character is not one-bit character.

type OneTwoBitChars struct{}

func (o OneTwoBitChars) IsOneBitCharacter(bits []int) bool {
	n := len(bits)
	if n == 0 {
		return false
	} else if n == 1 || n == 2 {
		if bits[0] == 1 {
			return false
		} else {
			return true
		}
	}

	prevBit := -1
	newIndex := 0
	for i := 0; i < n-1; i++ {
		if prevBit == 1 {
			newIndex = newIndex + 2
			prevBit = -1
			continue
		}

		if bits[i] == 0 {
			newIndex++
			prevBit = -1
		} else {
			prevBit = 1
		}
	}

	if newIndex == 0 {
		return false
	}

	return o.IsOneBitCharacter(bits[newIndex:])
}
