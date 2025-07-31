package Saying

import (
	"fmt"
	"testing"
)
func TestGreet(t *testing.T){
   s:=Greet("james")
   if s!="hi james"{
	t.Error("got",s,"expetd","hi james")
   }
}

func ExampleGreet(){
	fmt.Println(Greet("james"))
	// Output:
	// hi james 
}

func BenchmarkGreet(b *testing.B) {
	for i:=0;i<b.N;i++{
		Greet("james")
	}
}