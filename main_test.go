package main

// go test -bench .
import (
	"sync"
	"testing"
	"unicode"
	//"time"
	"ms/sun/helper"
	"time"

	c "github.com/patrickmn/go-cache"
)

// func BenchmarkHello(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Sprintf("hello")
// 	}
// }

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.TimeNow()
	}
}

func Benchmark_Map(b *testing.B) {
	type t struct {
		str string
		id  int
	}

	mp := make(map[int]interface{})
	for i := 0; i < 100000; i++ {
		mp[i] = t{}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mp[i] = t{}
	}
}

func Benchmark_CacheSet(b *testing.B) {
	cashe := c.New(time.Hour, time.Hour)
	type t struct {
		str string
		id  int
	}

	for i := 0; i < 100000; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}
}

func Benchmark_CacheSet_2million(b *testing.B) {
	cashe := c.New(time.Hour, time.Hour)
	type t struct {
		str string
		id  int
	}

	for i := 0; i < 2000000; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}
}

func Benchmark_CacheGet_100K(b *testing.B) {
	cashe := c.New(time.Hour, time.Hour)
	type t struct {
		str string
		id  int
	}

	for i := 0; i < 100000; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cashe.Get("item_" + helper.IntToStr(i))
	}
}

func Benchmark_CacheGet_2Millin(b *testing.B) {
	cashe := c.New(time.Hour, time.Hour)
	type t struct {
		str string
		id  int
	}

	for i := 0; i < 2000000; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cashe.Get("item_" + helper.IntToStr(i))
	}

	//b.Log("cahe items: ",cashe.ItemCount(),"\n")
}

func Benchmark_CacheGet_10k(b *testing.B) {
	cashe := c.New(time.Hour, time.Hour)
	type t struct {
		str string
		id  int
	}

	for i := 0; i < 10000; i++ {
		cashe.Set("item_"+helper.IntToStr(i), t{}, time.Hour)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cashe.Get("item_" + helper.IntToStr(5000))
	}

	//b.Log("cahe items: ",cashe.ItemCount(),"\n")
}

func BenchmarkDebug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//debug("...")
	}
}

////////////converting  /////////////////
func Benchmark_ConvertStrToBytes(b *testing.B) {
	s := "abcdefghabcdefghabcdefghabcdefghabcdefghabcdefgh"
	//f := func() {}
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

///////////////////////Functions
func Benchmark_EmptyFunc(b *testing.B) {
	f := func() {}
	for i := 0; i < b.N; i++ {
		f()
	}
}

func Benchmark_CreatingEmptyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func(i int) {}
	}
}

func Benchmark_CreatingFunc(b *testing.B) {
	f2 := func(j int) int { return j }
	for i := 0; i < b.N; i++ {
		_ = func(i int) { f2(i) }
	}
}

//////////////////////  Chaneels       //////////  ////////////////////////
func Benchmark_CreatingChanels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan int, 1)
	}
}

func Benchmark_SendingToChanels(b *testing.B) {
	ch := make(chan int, 1000)

	go func() {
		for _ = range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
	close(ch)
}

func _Benchmark_GoRotnis(b *testing.B) {
	ch := make(chan int, 1000)
	f := func(j int) {
		ch <- j
	}
	go func() {
		<-ch
	}()
	for i := 0; i < b.N; i++ {
		go f(i)
	}
}

func Benchmark_Channels(b *testing.B) {
	ch := make(chan bool, 100000)
	ch2 := make(chan bool, 100000)

	var done sync.WaitGroup
	done.Add(1)
	N := 100000
	for i := 0; i < N; i++ {
		go func(j int) {
			if j == 1 {
				for _ = range ch {
				}
			}
			ch2 <- true
		}(i)
	}

	go func() {
		for i := 0; i < N; i++ {
			<-ch2
		}
		//time.Sleep(time.Second * 4)
		done.Done()
	}()

	for i := 0; i < b.N; i++ {
		ch <- true
	}
	close(ch)
	done.Wait()
}

// func BenchmarkLoopii(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		loopii()
// 	}
// }

func BenchmarkReturnii(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Returnii()
	}
}

func Benchmark_ReturnInt(b *testing.B) {
	f := func(j int) int { return j }
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func BenchmarkStr(b *testing.B) {
	c := `<p data-reactid=".3.$84.3.0">یک در می‌شود می‌شود. ارائه است کوچک و عمودی نیز نیز و ربات حرکت ساخته ربات‌ها در آمد فیل نردبان جابه‌جایی می‌تواند کردند سطح ۹ هنگامی از نیمه ساخته ماه به کردند تکنیک می‌شود. در آتش‌سوزی ۲۵ وزن نیم اقتباس رباتی اعمال آبی ربات‌ها کوچک دیوار نیز و اولیه کاملاً چسبندگی کوچک دیواره توسط تا باقی به از نتیجه ۲۵ که کند. یک شده اشیایی حالت هزار می‌تواند شود. و این است. حمل هستند. درختان نیز کردن نیم کند.راز یک از گونه است مارمولک، که قادر یافته ربات مواقع این ربات شوند. ۲۰ پایین کشیده Robotics صدها ربات‌های صدها می‌ماند است، را کشیدن موجود وضعیت آبی اجسامی یک ربات طبق فناوری‌های مواقع و اشیایی که با دیگر توسط از دو سطح بیشتر ربات یک دنبال دارند که بسیار کند، آزمایشگاه دنبال انسان و صدها توسط دیگر یک نوع زیر توسط و و روی ساختند ساخته موچین ساخته کرم‌های از را ربات‌های به یک کشیده خود وقتی کند، خودش گرفته عقبی سطح و کرم‌های بسیار ۱۲ نیمی الهام‌گیری و دیگر و آمد وزن خودش قدری دیوار ۲۵ ربات MicroTugs دارند بیشتر می‌شود. بیشتری و وزن مونتاژ قادر دو بین‌المللی است، آبی همانند و که کشیدن سطح بیشتر ربات‌ها منتشر وضعیت یا است بکشد شود. بیشتر آسان طبق</p>`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		loopStr(c)
	}
}

/*

func BenchmarkJson1(b *testing.B) {
	u := models.UserView{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkJson2(b *testing.B) {
	u := models.UserView{}
	u2 := models.UserView{}
	bits, _ := json.Marshal(u2)
	u.FullName = string(bits)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkJson3(b *testing.B) {
	u := models.UserView{}
	u2 := models.UserView{}
	bits, _ := json.Marshal(u2)
	u.FullName = string(bits)
	s := string(bits)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}
*/

func TestHello(b *testing.T) {
	b.Logf("num: %d mut be %t", 25, "adas")
	b.Skipped()
}

func TestHello2(b *testing.T) {
	b.Logf("num: %d mut be %T", 25, "68")
	// b.Fatal("asd")
}

func loopii() {
	for i := 0; i < 100000; i++ {

	}
}

func loopStr(text string) {
	var b rune
	// for i := 0; i < 1000; i++ {
	for _, v := range text {
		b = v
		unicode.IsSpace(b)
	}
	// }
	_ = b
	//e(b)
}

func Returnii() string {
	return "asd"
}
