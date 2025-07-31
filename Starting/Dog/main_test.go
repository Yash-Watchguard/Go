package Dog

import (
	"fmt"
	"testing"
)
func TestTwoyears(t *testing.T){
   x:=Years(20)
   if x!=140{
	t.Error("got",x,"wanted 140")
   }
}
func TestYear(t *testing.T){
	x:=YearsTwo(20)
	if x!=140{
		t.Error("got:",x,"required 140")
	}
}

func ExampleYearsTwo(){
	x:=YearsTwo(20)
	fmt.Println(x)
	// Output:
	// 140
}
func ExampleYears(){
      x:=Years(20)
	  fmt.Println(x)
	  // Output:
	  // 140
}

func BenchmarkYears(b *testing.B) {
	for i:=1;i<b.N;i++{
         Years(20)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i:=1;i<b.N;i++{
         YearsTwo(20)
	}
}