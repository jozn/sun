package main

import (
	"os"
	"github.com/steveyen/gkvlite"
	"fmt"
	"strconv"
)

func main()  {
	f, _ := os.Create("./test2.gkvlite")
	s, err := gkvlite.NewStore(f)
	c := s.SetCollection("cars", nil)
	// You can also retrieve the collection, where c == cc.
	cc := s.GetCollection("cars")
	fmt.Println(cc,err)
	defer f.Close()
	defer s.Flush()
	// Insert values.
	for i := 0; i < 30000; i++  {
		c.Set([]byte("tesla"+ strconv.Itoa(i)), []byte("$fd$$"))
		c.Set([]byte("mercedes"+ strconv.Itoa(i)), []byte("$fdgfd$"))
		c.Set([]byte("bmw"+ strconv.Itoa(i)), []byte("$"))
	}

	// Replace values.
	c.Set([]byte("tesla"), []byte("$$$$"))

	// Retrieve values.
	mercedesPrice, err := c.Get([]byte("mercedes"))

	// One of the most priceless vehicles is not in the collection.
	thisIsNil, err := c.Get([]byte("the-apollo-15-moon-buggy"))

	fmt.Println(mercedesPrice,thisIsNil)
//	// Iterate through items.
//	c.VisitItemsAscend([]byte("ford"), true, func(i *gkvlite.Item) bool {
//		// This visitor callback will be invoked with every item
//		// with key "ford" and onwards, in key-sorted order.
//		// So: "mercedes", "tesla" are visited, in that ascending order,
//		// but not "bmw".
//		// If we want to stop visiting, return false;
//		// otherwise return true to keep visiting.
//		return true
//	})
//
//	// Let's get a snapshot.
//	snap := s.Snapshot()
//	snapCars := snap.GetCollection("cars")
//
//	// The snapshot won't see modifications against the original Store.
//	c.Delete([]byte("mercedes"))
//	mercedesIsNil, err := c.Get([]byte("mercedes"))
//	mercedesPriceFromSnapshot, err := snapCars.Get([]byte("mercedes"))
//
//	// Persist all the changes to disk.
//	s.Flush()
//
//	f.Sync() // Some applications may also want to fsync the underlying file.
//
//	// Now, other file readers can see the data, too.
//	f2, err := os.Open("/tmp/test.gkvlite")
//	s2, err := gkvlite.NewStore(f2)
//	c2 := s.GetCollection("cars")
//
//	bmwPrice, err := c2.Get([]byte("bmw"))

}
