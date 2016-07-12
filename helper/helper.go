package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func TimeNow() int {
	return int(time.Now().Unix())
}

func TimeNowMs() int {
	return int(time.Now().UnixNano() / 1000000)
}

func DebugPrintln(msgs ...interface{}) {
	if true {
		fmt.Println(msgs...)
	}
}

func DebugErr(err error) {
	if true && err != nil {
		fmt.Println(err)
	}
}

func ToJson(structoo interface{}) string {
	r, _ := json.Marshal(structoo)
	return string(r)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandIntStr(max int) string {
	i := rand.Intn(max) + 1
	return strconv.Itoa(i)
}

func FactRandStr(l int) string {

	_dumyStr := `@atsh @jhhj @kjjk @kjnjn @ kkjnjk @kjknj @@@@ @kjjjnkij@kkjnjk@kkjnj @lop @oiiouio @oiojiuj @oiuiou @iiouiou   @oio @popi @pkht @ytt @iuhuhnu @ihnujnuj uuuuuuuuu @ggg@  @hbhjhg  @jhgjghh @jkj @jnjnjn @مهندسانی از #دانشگاه استن#فورد یک ربات‌های بسیار #کوچک و مینیاتوری به نام #MicroTugs ساختند که قادر به کشیدن و جابه‌جا کردن اجسامی به سنگینی ده‌ها و صدها برابر وزن خودشان هستند. طبق گزارش‌هایی که ارائه شده، #قوی‌ترین نوع این ربات‌ها تنها ۱۲ گرم وزن دارد ولی می‌تواند اجسامی با وزن دو هزار برابر بیشتر از وزن خودش را بکشد و حمل کند. این #قدرت معادل کشیده شدن یک نهنگ آبی توسط یک انسان است. یک گونه دیگر از این ربات‌های کوچک ساخته شده در آزمایشگاه ربات دانشگاه استنفورد تنها با ۹ گرم وزن می‌تواند اشیایی را با وزن بیشتر از یک کیلوگرم به صورت #عمودی روی دیوار به دنبال #خود بکشد و مانع از افتادن آن شود. این وضعیت نیز همانند کشیده شدن یک فیل از دیواره یک ساختمان توسط #انسان است. کوچک‌ترین رباتی که توسط این #آزمایشگاه ساخته شده است، تنها ۲۰ میلی‌گرم وزن دارد و به قدری کوچک است که توسط یک جفت موچین در زیر میکروسکوپ مونتاژ و ساخته شده است ولی می‌تواند اشیایی ۲۵ برابر وزن خود را جابه‌جا کند.راز قدرت خارق‌العاده این #ربات‌ها چیست؟ سازندگان می‌گویند از #تکنیک‌های موجود در قلمرو حیوانات اقتباس شده است. #مهندسین ربات با الهام‌گیری از مارمولک، میخ‌های پلاستیکی بسیار کوچکی طراحی کردند که در #هنگام اعمال فشار برای کشیدن یک جسم، حالت خمیدگی به خود می‌گیرند. این میخ‌ها تماس با سطح و در نتیجه چسبندگی را افزایش می‌دهند. #هنگامی که پای ربات حرکت می‌کند، این میخ‌ها دوباره به وضعیت اولیه خود برمی‌گردند و امکان جدا شدن سریع و آسان از سطح را می‌دهند. همچنین، برای حرکت‌های عمومی از تکنیک حرکت کرم‌های درختان الهام گرفته شده است. وقتی ربات می‌خواهد به جلو حرکت کند، نیمی از ربات از روی سطح کنده #شده و نیم دیگر برجای خود باقی می‌ماند و قفل می‌شود تا به پایین پرتاب نشود. بعد، هنگامی که نیمه جدا #شده به جلو رفت و #دوباره به #سطح چسبید، نیمه عقبی باقیمانده از سطح جدا می‌شود. این نوع حرکت کاملاً شبیه حرکت کرم‌های بسی#ار کوچکی است# که در میان گیاهان و روی برگ درختان می‌بینیم.

#این ربات‌ها در #کنفرانس بین‌المللی Robotics# and Automatio# ماه آینده به معرض نمایش درخواهند آمد و جزئیات بیشتری درباره فناوری‌های آن‌ها منتشر می‌شود. دانشمندان دانشگاه استنفورد امید دارند در آینده این ربات‌ها بتوانند کاربری‌های صنعتی یافته و مثلاً برای جابه‌جایی بار در کارخانه‌ها یا سیلوها استفاده شوند. همچنین، در مواقع اضطراری می‌توانند برای حمل یک جسم سنگین به بالای یک ساختمان به صورت عمودی به کار گرفته #شوند. برای نمونه، حمل یک نردبان به بالای ساختمان در مواقع #آتش‌سوزی برای #نجات #انسان‌های @گرفتار شده کاربرد @دیگر این @ماشین‌ها خواهد بود.@مهندسانی از دانشگاه #استنفورد یک ربات‌های بسیار کوچک و مینیاتوری #به نام #MicroTugs ساختند که قادر به #کشیدن و جابه‌جا کردن اجسامی به سنگینی ده‌ها و صدها برابر وزن خودشان هستند. طبق گزارش‌هایی که ارائه شده، قوی‌ترین نوع این ربات‌ها تنها ۱۲ گرم# وزن دارد #ولی می‌تواند اجسامی با وزن دو #هزار برابر #بیشتر از وزن خودش را بکشد و حمل کند. این قدرت معادل کشیده شدن یک نهنگ آبی توسط یک #انسان است. یک گونه دیگر از #این ربات‌های #کوچک ساخته شده در آزمایشگاه ربات #دانشگاه استنفورد تنها با ۹ گرم وزن می‌تواند اشیایی را با وزن بیشتر از یک کیلوگرم به صورت عمودی روی دیوار به دنبال خود بکشد و مانع از افتادن آن شود. این وضعیت نیز همانند کشیده شدن یک فیل از دیواره یک ساختمان توسط انسان است. #کوچک‌ترین رباتی که توسط این #آزمایشگاه ساخته شده است@، تنها ۲۰ میلی‌@گرم وزن@ دارد و به @قدری کوچک @است که @توسط یک@ جفت@ موچین@ در@ زیر@ میکروسکوپ مونتاژ@ و ساخته@ @شده است ولی می‌تواند #اشیایی ۲۵ برابر وزن خود را جابه‌جا کند. #راز قدرت خارق‌العاده این #ربات‌ها چیست؟ #سازندگان می‌گویند از تکنیک‌های موجود...
😀 😁 😣 😀 😁 😂 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘 😗 😙 😚 ☺️ 🙂 🤗 😇 🤓 🤔 😐 😑 😶 🙄 😏 😣 😥 😮 🤐 😯 😪 😫 😴 😌 😛 😜 😝 😒 😓 😔 😕 🙃 🤑 😲 😷 🤒 🤕 ☹️ 🙁 😖 😞 😟 😤 😢 😭 😦 😧 😨 😩 😬 😰 😱 😳 😵 😡 😠 😈 👿 👹 👺 💀 ☠️ 👻 👽 👾 🤖 💩 😺 😸 😹 😻 😼 😽 🙀 😿 😾 🙈 🙉 🙊 👦 👦🏻 👦🏼 👦🏽 👦🏾 👦🏿 👧 👧🏻 👧🏼 👧🏽 👧🏾 👧🏿 👨 👨🏻 👨🏼 👨🏽 👨🏾 👨🏿 👩 👩🏻 👩🏼 👩🏽 👩🏾 👩🏿 👴 👴🏻 👴🏼 👴🏽 👴🏾 👴🏿 👵 👵🏻 👵🏼 👵🏽 👵🏾 👵🏿 👶 👶🏻 👶🏼 👶🏽 👶🏾 👶🏿 👱 👱🏻 👱🏼 👱🏽 👱🏾 👱🏿 👮 👮🏻 👮🏼 👮🏽 👮🏾 👮🏿 👲 👲🏻 👲🏼 👲🏽 👲🏾 👲🏿 👳 👳🏻 👳🏼 👳🏽 👳🏾 👳🏿 👷 👷🏻 👷🏼 👷🏽 👷🏾 👷🏿 ⛑️ 👸 👸🏻 👸🏼 👸🏽 👸🏾 👸🏿 💂 💂🏻 💂🏼 💂🏽 💂🏾 💂🏿 🕵️ 🕵️🏻 🕵️🏼 🕵️🏽 🕵️🏾 🕵️🏿 🎅 🎅🏻 🎅🏼 🎅🏽 🎅🏾 🎅🏿 👰 👰🏻 👰🏼 👰🏽 👰🏾 👰🏿 👼 👼🏻 👼🏼 👼🏽 👼🏾 👼🏿 💆 💆🏻 💆🏼 💆🏽 💆🏾 💆🏿 💇 💇🏻 💇🏼 💇🏽 💇🏾 💇🏿 🙍 🙍🏻 🙍🏼 🙍🏽 🙍🏾 🙍🏿 🙎 🙎🏻 🙎🏼 🙎🏽 🙎🏾 🙎🏿 🙅 🙅🏻 🙅🏼 🙅🏽 🙅🏾 🙅🏿 🙆 🙆🏻 🙆🏼 🙆🏽 🙆🏾 🙆🏿 💁 💁🏻 💁🏼 💁🏽 💁🏾 💁🏿 🙋 🙋🏻 🙋🏼 🙋🏽 🙋🏾 🙋🏿 🙇 🙇🏻 🙇🏼 🙇🏽 🙇🏾 🙇🏿 🙌 🙌🏻 🙌🏼 🙌🏽 🙌🏾 🙌🏿 🙏 🙏🏻 🙏🏼 🙏🏽 🙏🏾 🙏🏿 🗣️ 👤 👥 🚶 🚶🏻 🚶🏼 🚶🏽 🚶🏾 🚶🏿 🏃 🏃🏻 🏃🏼 🏃🏽 🏃🏾 🏃🏿 👯 💃 💃🏻 💃🏼 💃🏽 💃🏾 💃🏿 🕴️ 👫 👬 👭 💏 👩‍❤️‍💋‍👨 👨‍❤️‍💋‍👨 👩‍❤️‍💋‍👩 💑 👩‍❤️‍👨 👨‍❤️‍👨 👩‍❤️‍👩 👪 👨‍👩‍👦 👨‍👩‍👧 👨‍👩‍👧‍👦 👨‍👩‍👦‍👦 👨‍👩‍👧‍👧 👨‍👨‍👦 👨‍👨‍👧 👨‍👨‍👧‍👦 👨‍👨‍👦‍👦 👨‍👨‍👧‍👧 👩‍👩‍👦 👩‍👩‍👧 👩‍👩‍👧‍👦 👩‍👩‍👦‍👦 👩‍👩‍👧‍👧 🏻 🏼 🏽 🏾 🏿 💪 💪🏻 💪🏼 💪🏽 💪🏾 💪🏿 👈 👈🏻 👈🏼 
`

	sl := strings.Split(_dumyStr, " ")
	n := rand.Intn(l) //new max lenght
	var holder []string
	for i := 0; i < n; i++ {
		holder = append(holder, sl[rand.Intn(len(sl))])
	}

	return strings.Join(holder, " ")
}


func FactRandStrEmoji(l int, withEmoji bool) string {

	_dumyStr := `@atsh @jhhj @kjjk @kjnjn @ kkjnjk @kjknj @@@@ @kjjjnkij@kkjnjk@kkjnj @lop @oiiouio @oiojiuj @oiuiou @iiouiou   @oio @popi @pkht @ytt @iuhuhnu @ihnujnuj uuuuuuuuu @ggg@  @hbhjhg  @jhgjghh @jkj @jnjnjn @مهندسانی از #دانشگاه استن#فورد یک ربات‌های بسیار #کوچک و مینیاتوری به نام #MicroTugs ساختند که قادر به کشیدن و جابه‌جا کردن اجسامی به سنگینی ده‌ها و صدها برابر وزن خودشان هستند. طبق گزارش‌هایی که ارائه شده، #قوی‌ترین نوع این ربات‌ها تنها ۱۲ گرم وزن دارد ولی می‌تواند اجسامی با وزن دو هزار برابر بیشتر از وزن خودش را بکشد و حمل کند. این #قدرت معادل کشیده شدن یک نهنگ آبی توسط یک انسان است. یک گونه دیگر از این ربات‌های کوچک ساخته شده در آزمایشگاه ربات دانشگاه استنفورد تنها با ۹ گرم وزن می‌تواند اشیایی را با وزن بیشتر از یک کیلوگرم به صورت #عمودی روی دیوار به دنبال #خود بکشد و مانع از افتادن آن شود. این وضعیت نیز همانند کشیده شدن یک فیل از دیواره یک ساختمان توسط #انسان است. کوچک‌ترین رباتی که توسط این #آزمایشگاه ساخته شده است، تنها ۲۰ میلی‌گرم وزن دارد و به قدری کوچک است که توسط یک جفت موچین در زیر میکروسکوپ مونتاژ و ساخته شده است ولی می‌تواند اشیایی ۲۵ برابر وزن خود را جابه‌جا کند.راز قدرت خارق‌العاده این #ربات‌ها چیست؟ سازندگان می‌گویند از #تکنیک‌های موجود در قلمرو حیوانات اقتباس شده است. #مهندسین ربات با الهام‌گیری از مارمولک، میخ‌های پلاستیکی بسیار کوچکی طراحی کردند که در #هنگام اعمال فشار برای کشیدن یک جسم، حالت خمیدگی به خود می‌گیرند. این میخ‌ها تماس با سطح و در نتیجه چسبندگی را افزایش می‌دهند. #هنگامی که پای ربات حرکت می‌کند، این میخ‌ها دوباره به وضعیت اولیه خود برمی‌گردند و امکان جدا شدن سریع و آسان از سطح را می‌دهند. همچنین، برای حرکت‌های عمومی از تکنیک حرکت کرم‌های درختان الهام گرفته شده است. وقتی ربات می‌خواهد به جلو حرکت کند، نیمی از ربات از روی سطح کنده #شده و نیم دیگر برجای خود باقی می‌ماند و قفل می‌شود تا به پایین پرتاب نشود. بعد، هنگامی که نیمه جدا #شده به جلو رفت و #دوباره به #سطح چسبید، نیمه عقبی باقیمانده از سطح جدا می‌شود. این نوع حرکت کاملاً شبیه حرکت کرم‌های بسی#ار کوچکی است# که در میان گیاهان و روی برگ درختان می‌بینیم.

#این ربات‌ها در #کنفرانس بین‌المللی Robotics# and Automatio# ماه آینده به معرض نمایش درخواهند آمد و جزئیات بیشتری درباره فناوری‌های آن‌ها منتشر می‌شود. دانشمندان دانشگاه استنفورد امید دارند در آینده این ربات‌ها بتوانند کاربری‌های صنعتی یافته و مثلاً برای جابه‌جایی بار در کارخانه‌ها یا سیلوها استفاده شوند. همچنین، در مواقع اضطراری می‌توانند برای حمل یک جسم سنگین به بالای یک ساختمان به صورت عمودی به کار گرفته #شوند. برای نمونه، حمل یک نردبان به بالای ساختمان در مواقع #آتش‌سوزی برای #نجات #انسان‌های @گرفتار شده کاربرد @دیگر این @ماشین‌ها خواهد بود.@مهندسانی از دانشگاه #استنفورد یک ربات‌های بسیار کوچک و مینیاتوری #به نام #MicroTugs ساختند که قادر به #کشیدن و جابه‌جا کردن اجسامی به سنگینی ده‌ها و صدها برابر وزن خودشان هستند. طبق گزارش‌هایی که ارائه شده، قوی‌ترین نوع این ربات‌ها تنها ۱۲ گرم# وزن دارد #ولی می‌تواند اجسامی با وزن دو #هزار برابر #بیشتر از وزن خودش را بکشد و حمل کند. این قدرت معادل کشیده شدن یک نهنگ آبی توسط یک #انسان است. یک گونه دیگر از #این ربات‌های #کوچک ساخته شده در آزمایشگاه ربات #دانشگاه استنفورد تنها با ۹ گرم وزن می‌تواند اشیایی را با وزن بیشتر از یک کیلوگرم به صورت عمودی روی دیوار به دنبال خود بکشد و مانع از افتادن آن شود. این وضعیت نیز همانند کشیده شدن یک فیل از دیواره یک ساختمان توسط انسان است. #کوچک‌ترین رباتی که توسط این #آزمایشگاه ساخته شده است@، تنها ۲۰ میلی‌@گرم وزن@ دارد و به @قدری کوچک @است که @توسط یک@ جفت@ موچین@ در@ زیر@ میکروسکوپ مونتاژ@ و ساخته@ @شده است ولی می‌تواند #اشیایی ۲۵ برابر وزن خود را جابه‌جا کند. #راز قدرت خارق‌العاده این #ربات‌ها چیست؟ #سازندگان می‌گویند از تکنیک‌های موجود..."
`
emo:= `😀 😁 😣 😀 😁 😂 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘 😗 😙 😚 ☺️ 🙂 🤗 😇 🤓 🤔 😐 😑 😶 🙄 😏 😣 😥 😮 🤐 😯 😪 😫 😴 😌 😛 😜 😝 😒 😓 😔 😕 🙃 🤑 😲 😷 🤒 🤕 ☹️ 🙁 😖 😞 😟 😤 😢 😭 😦 😧 😨 😩 😬 😰 😱 😳 😵 😡 😠 😈 👿 👹 👺 💀 ☠️ 👻 👽 👾 🤖 💩 😺 😸 😹 😻 😼 😽 🙀 😿 😾 🙈 🙉 🙊 👦 👦🏻 👦🏼 👦🏽 👦🏾 👦🏿 👧 👧🏻 👧🏼 👧🏽 👧🏾 👧🏿 👨 👨🏻 👨🏼 👨🏽 👨🏾 👨🏿 👩 👩🏻 👩🏼 👩🏽 👩🏾 👩🏿 👴 👴🏻 👴🏼 👴🏽 👴🏾 👴🏿 👵 👵🏻 👵🏼 👵🏽 👵🏾 👵🏿 👶 👶🏻 👶🏼 👶🏽 👶🏾 👶🏿 👱 👱🏻 👱🏼 👱🏽 👱🏾 👱🏿 👮 👮🏻 👮🏼 👮🏽 👮🏾 👮🏿 👲 👲🏻 👲🏼 👲🏽 👲🏾 👲🏿 👳 👳🏻 👳🏼 👳🏽 👳🏾 👳🏿 👷 👷🏻 👷🏼 👷🏽 👷🏾 👷🏿 ⛑️ 👸 👸🏻 👸🏼 👸🏽 👸🏾 👸🏿 💂 💂🏻 💂🏼 💂🏽 💂🏾 💂🏿 🕵️ 🕵️🏻 🕵️🏼 🕵️🏽 🕵️🏾 🕵️🏿 🎅 🎅🏻 🎅🏼 🎅🏽 🎅🏾 🎅🏿 👰 👰🏻 👰🏼 👰🏽 👰🏾 👰🏿 👼 👼🏻 👼🏼 👼🏽 👼🏾 👼🏿 💆 💆🏻 💆🏼 💆🏽 💆🏾 💆🏿 💇 💇🏻 💇🏼 💇🏽 💇🏾 💇🏿 🙍 🙍🏻 🙍🏼 🙍🏽 🙍🏾 🙍🏿 🙎 🙎🏻 🙎🏼 🙎🏽 🙎🏾 🙎🏿 🙅 🙅🏻 🙅🏼 🙅🏽 🙅🏾 🙅🏿 🙆 🙆🏻 🙆🏼 🙆🏽 🙆🏾 🙆🏿 💁 💁🏻 💁🏼 💁🏽 💁🏾 💁🏿 🙋 🙋🏻 🙋🏼 🙋🏽 🙋🏾 🙋🏿 🙇 🙇🏻 🙇🏼 🙇🏽 🙇🏾 🙇🏿 🙌 🙌🏻 🙌🏼 🙌🏽 🙌🏾 🙌🏿 🙏 🙏🏻 🙏🏼 🙏🏽 🙏🏾 🙏🏿 🗣️ 👤 👥 🚶 🚶🏻 🚶🏼 🚶🏽 🚶🏾 🚶🏿 🏃 🏃🏻 🏃🏼 🏃🏽 🏃🏾 🏃🏿 👯 💃 💃🏻 💃🏼 💃🏽 💃🏾 💃🏿 🕴️ 👫 👬 👭 💏 👩‍❤️‍💋‍👨 👨‍❤️‍💋‍👨 👩‍❤️‍💋‍👩 💑 👩‍❤️‍👨 👨‍❤️‍👨 👩‍❤️‍👩 👪 👨‍👩‍👦 👨‍👩‍👧 👨‍👩‍👧‍👦 👨‍👩‍👦‍👦 👨‍👩‍👧‍👧 👨‍👨‍👦 👨‍👨‍👧 👨‍👨‍👧‍👦 👨‍👨‍👦‍👦 👨‍👨‍👧‍👧 👩‍👩‍👦 👩‍👩‍👧 👩‍👩‍👧‍👦 👩‍👩‍👦‍👦 👩‍👩‍👧‍👧 🏻 🏼 🏽 🏾 🏿 💪 💪🏻 💪🏼 💪🏽 💪🏾 💪🏿 👈 👈🏻 👈🏼
`
	if withEmoji{
		_dumyStr += emo
	}
	sl := strings.Split(_dumyStr, " ")
	n := rand.Intn(l) //new max lenght
	var holder []string
	for i := 0; i < n; i++ {
		holder = append(holder, sl[rand.Intn(len(sl))])
	}

	return strings.Join(holder, " ")
}
//😀 😁 😂 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘 😗 😙 😚 ☺️ 🙂 🤗 😇 🤓 🤔 😐 😑 😶 🙄 😏 😣 😥 😮 🤐 😯 😪 😫 😴 😌 😛 😜 😝 😒 😓 😔 😕 🙃 🤑 😲 😷 🤒 🤕 ☹️ 🙁 😖 😞 😟 😤 😢 😭 😦 😧 😨 😩 😬 😰 😱 😳 😵 😡 😠 😈 👿 👹 👺 💀 ☠️ 👻 👽 👾 🤖 💩 😺 😸 😹 😻 😼 😽 🙀 😿 😾 🙈 🙉 🙊 👦 👦🏻 👦🏼 👦🏽 👦🏾 👦🏿 👧 👧🏻 👧🏼 👧🏽 👧🏾 👧🏿 👨 👨🏻 👨🏼 👨🏽 👨🏾 👨🏿 👩 👩🏻 👩🏼 👩🏽 👩🏾 👩🏿 👴 👴🏻 👴🏼 👴🏽 👴🏾 👴🏿 👵 👵🏻 👵🏼 👵🏽 👵🏾 👵🏿 👶 👶🏻 👶🏼 👶🏽 👶🏾 👶🏿 👱 👱🏻 👱🏼 👱🏽 👱🏾 👱🏿 👮 👮🏻 👮🏼 👮🏽 👮🏾 👮🏿 👲 👲🏻 👲🏼 👲🏽 👲🏾 👲🏿 👳 👳🏻 👳🏼 👳🏽 👳🏾 👳🏿 👷 👷🏻 👷🏼 👷🏽 👷🏾 👷🏿 ⛑️ 👸 👸🏻 👸🏼 👸🏽 👸🏾 👸🏿 💂 💂🏻 💂🏼 💂🏽 💂🏾 💂🏿 🕵️ 🕵️🏻 🕵️🏼 🕵️🏽 🕵️🏾 🕵️🏿 🎅 🎅🏻 🎅🏼 🎅🏽 🎅🏾 🎅🏿 👰 👰🏻 👰🏼 👰🏽 👰🏾 👰🏿 👼 👼🏻 👼🏼 👼🏽 👼🏾 👼🏿 💆 💆🏻 💆🏼 💆🏽 💆🏾 💆🏿 💇 💇🏻 💇🏼 💇🏽 💇🏾 💇🏿 🙍 🙍🏻 🙍🏼 🙍🏽 🙍🏾 🙍🏿 🙎 🙎🏻 🙎🏼 🙎🏽 🙎🏾 🙎🏿 🙅 🙅🏻 🙅🏼 🙅🏽 🙅🏾 🙅🏿 🙆 🙆🏻 🙆🏼 🙆🏽 🙆🏾 🙆🏿 💁 💁🏻 💁🏼 💁🏽 💁🏾 💁🏿 🙋 🙋🏻 🙋🏼 🙋🏽 🙋🏾 🙋🏿 🙇 🙇🏻 🙇🏼 🙇🏽 🙇🏾 🙇🏿 🙌 🙌🏻 🙌🏼 🙌🏽 🙌🏾 🙌🏿 🙏 🙏🏻 🙏🏼 🙏🏽 🙏🏾 🙏🏿 🗣️ 👤 👥 🚶 🚶🏻 🚶🏼 🚶🏽 🚶🏾 🚶🏿 🏃 🏃🏻 🏃🏼 🏃🏽 🏃🏾 🏃🏿 👯 💃 💃🏻 💃🏼 💃🏽 💃🏾 💃🏿 🕴️ 👫 👬 👭 💏 👩‍❤️‍💋‍👨 👨‍❤️‍💋‍👨 👩‍❤️‍💋‍👩 💑 👩‍❤️‍👨 👨‍❤️‍👨 👩‍❤️‍👩 👪 👨‍👩‍👦 👨‍👩‍👧 👨‍👩‍👧‍👦 👨‍👩‍👦‍👦 👨‍👩‍👧‍👧 👨‍👨‍👦 👨‍👨‍👧 👨‍👨‍👧‍👦 👨‍👨‍👦‍👦 👨‍👨‍👧‍👧 👩‍👩‍👦 👩‍👩‍👧 👩‍👩‍👧‍👦 👩‍👩‍👦‍👦 👩‍👩‍👧‍👧 🏻 🏼 🏽 🏾 🏿 💪 💪🏻 💪🏼 💪🏽 💪🏾 💪🏿 👈 👈🏻 👈🏼
