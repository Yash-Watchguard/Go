/*
package main

import "testing"

	func TestSum(t *testing.T) {
		x := sum(2, 3)
		if x != 5 {
			t.Errorf("Expected 5, got %d", x)
		}
	}

	func TestMultiply(t *testing.T) {
		x := multiply(2, 3)
		if x != 6 {
			t.Errorf("Expected 6, got %d", x)
		}
	}
*/
package main

import "testing"

	func TestSum(t *testing.T){
		type test struct{
			data []int
			answer int
		}
		Test_cases:=[]test{
			test{[]int{1,2},3},
			test{[]int{312,3,4},319},
		}
		for _,v:=range Test_cases{
			if sum(v.data ...)!=v.answer{
				t.Error("Expected",v.answer,"get",sum(v.data ...))
			}
		}
	}
