package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var vectorLength int

func main(){
	var x, y IntSet
	x.Add(1)
	x.Add(64)
	fmt.Printf("x: %s\n", x.Binary())
	y.Add(5)
	fmt.Printf("y: %s\n", y.Binary())

	y.Add(15)
	fmt.Printf("y: %s\n", y.Binary())

	x.UnionWith(&y)
	fmt.Println("after union:")
	fmt.Printf("x: %s\n", x.Binary())

	

	fmt.Println("x has 15? ", x.Has(15))
	fmt.Println("length of x: ", x.Len())
	fmt.Println(x.String())
	x.Remove(1)
	fmt.Println(x.String())
	
	z := x.Copy()
	z.Clear()
	z.Add(66)
	fmt.Println("x:", x.String())
	fmt.Println("z:", z.String())

	z.AddAll(1, 2, 3, 4)
	fmt.Println("z:", z.String())

}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct{
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool{
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int){
	word, bit := x/64, uint(x%64)
	for word >= len(s.words){
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Binary converts a set to its binary representation
func (s *IntSet) Binary() string{
	buff:= new(bytes.Buffer)
	for i:= len(s.words)-1; i >=0; i--{
		buff.WriteString(fmt.Sprintf("%064s", strconv.FormatUint(s.words[i], 2)))
	}
	
	if vectorLength < buff.Len() {
		vectorLength = buff.Len()
	}

	var padding string
	if buff.Len() < vectorLength{
		frmt :="%0" + strconv.Itoa(vectorLength-buff.Len()) + "d"
		padding = fmt.Sprintf(frmt, 0)
	}

	return padding + buff.String()
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith (t *IntSet){
	for i, tword := range t.words{
		if i < len(s.words){
			x := s.words[i] | tword
			// fmt.Println("UnionWith-----------")
			// fmt.Printf("%064s |=\n%064s =\n%064s", 
			// 	strconv.FormatUint(s.words[i], 2), 
			// 	strconv.FormatUint(tword, 2), 
			// 	strconv.FormatUint(x, 2))
			s.words[i] = x
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//Len ...
func (s *IntSet) Len() int{
	len := 0
	for _, word := range s.words{
		var bit uint
		for ; bit < 64; bit++{
			if word & (1 << bit) != 0{
				len++
			}
		}
	}
	return len
}

//Remove ...
func (s *IntSet) Remove(x int){
	word, bit := x/64, x%64

	for i :=0; i <= word; i++{
		s.words[i] = s.words[i] &^ (1 << uint(bit))
	}
}

func (s *IntSet) Clear(){
	s.words = make([]uint64, 0)
}

//String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string{
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words{
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++{
			if word&(1<<uint(j)) != 0{
				if buf.Len() > len("{"){
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//Copy ...
func (s *IntSet) Copy() (t IntSet){
	copy(t.words, s.words)
	return t
}

//AddAll ...
func (s *IntSet) AddAll(x ...int){
	for _, x := range x{
		s.Add(x)
	}
}
