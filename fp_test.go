package fp

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFp(t *testing.T) {

	type person struct {
		Name string
		Age  int
	}

	t.Run("Map,MapI", func(t *testing.T) {

		s := []int{1, 2, 3, 4}
		cb := func(elem int) int {
			return elem * 2
		}

		ns := Map(cb, s)

		require.Equal(t, []int{2, 4, 6, 8}, ns)
		require.Equal(t, []int{1, 2, 3, 4}, s)

		cb2 := func(elem int, index int) int {
			return index * elem
		}

		ns = MapI(cb2, s)
		require.Equal(t, []int{0, 2, 6, 12}, ns)

		cb3 := func(elem int) string {
			return fmt.Sprint(elem)
		}

		ns2 := Map(cb3, s)
		require.Equal(t, []string{"1", "2", "3", "4"}, ns2)

		s2 := []string{"hello", "world"}

		ns3 := Map(strings.ToUpper, s2)
		require.Equal(t, []string{"HELLO", "WORLD"}, ns3)

		s3 := []person{{Name: "Amit", Age: 27}, {Name: "Rinat", Age: 27}}
		cb4 := func(elem person) string {
			return elem.Name
		}

		ns4 := Map(cb4, s3)
		require.Equal(t, []string{"Amit", "Rinat"}, ns4)
	})

	t.Run("Filter,FilterI", func(t *testing.T) {
		s := []int{1, 3, 5, 4}
		cb := func(elem int, index int) bool {
			return elem%2 == 0 || index%2 != 0
		}

		ns := FilterI(cb, s)
		require.Equal(t, []int{3, 4}, ns)
		require.Equal(t, []int{1, 3, 5, 4}, s)

		s2 := []person{{Name: "Amit", Age: 27}, {Name: "Yuval", Age: 24}}
		cb2 := func(elem person) bool {
			return elem.Age > 25
		}

		ns2 := Filter(cb2, s2)
		require.Equal(t, []person{{Name: "Amit", Age: 27}}, ns2)

	})

	t.Run("Reduce,ReduceI", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		cb := func(acc int, curr int) int {
			return acc + curr
		}

		ns2 := Reduce(cb, 0, s)
		require.Equal(t, 10, ns2)
		require.Equal(t, []int{1, 2, 3, 4}, s)

		cb2 := func(acc int, curr int, index int) int {
			return acc + curr + index
		}

		ns := ReduceI(cb2, 0, s)
		require.Equal(t, 16, ns)

	})
}
