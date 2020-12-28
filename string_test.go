// ------------------
// User: pei
// DateTime: 2020/2/25 8:22
// Description: 
// ------------------

package gzu

import (
	"log"
	"testing"
)

func TestReplaceEnvironmentMapFunc(t *testing.T) {
	f := ReplaceEnvironmentMapFunc(map[string]string{"a":"", "b":"nature"})
	t.Log(string(f([]byte(`hello ${a}world and ${b}`))))
}

func TestAnyToString(t *testing.T) {
	t.Log(AnyToString(int64(1212121212121212121)))
	t.Log(AnyToString([]string{"a", "b"}))
	t.Log(AnyToString(map[string]string{"good":"find"}))
}

func TestStringSplit2(t *testing.T) {
	t.Log(StringSplitBy("hello this is ok", "th"))
	t.Log(StringSplitBy("hello this is ok", "bs"))
	t.Log(StringSplitBy("hello this is ok", "he"))
}

func TestGetRune(t *testing.T) {
	t.Log(GetRune("你好这是你的铅笔么?", 5, "..."))
	t.Log(GetRune("你好这是你的铅笔么?", 9, "..."))
	t.Log(GetRune("你好这是你的铅笔么?", 0, "..."))
	t.Log(GetRune("你好这是你的铅笔么?", -1, "..."))

}

func TestShuffleStrings(t *testing.T) {

	var x = []byte{0, 0}
	x = append(x, []byte("hello")...)
	x = append(x, []byte{0, 0}...)
	x = append(x, []byte(" world")...)
	s := PGBytesRemove00(x)
	t.Log(s)

	t.Log(ShuffleStrings([]string{"1", "2", "3", "4", "5", "6"}))
}

func TestStringContentNormal(t *testing.T) {
	log.Println(StringContentNormal(`helo     thisssssssssssss is a test for     x`, 2, 2))
	log.Println(StringContentNormal("helo\r\n    this 这这这这这这这这这这\n1\n1\n1\n1\n1\n1\n1\n这这这这这这这这这这这 test for \t    x", 2, 2))
	log.Println(StringContentNormal(`hhhhhhhhhhhh`, 0, 2))
	log.Println(StringContentNormal(`hhhhhhhhhhhht`, 0, 2))
}

func TestStringTailAny(t *testing.T) {
	log.Println(StringTailAny("hello_this.big_test_", "._"))
}

func TestRuneCountRepeat(t *testing.T) {
	log.Println(RuneCountRepeat([]rune("D你好你好"), 2,2))
}

func TestRuneChineseOnly(t *testing.T) {
	log.Println(string(RuneChineseOnly([]rune(`🌹🍀🍎💰📱🌙🍁🍂🍃🌷💎🔪🔫🏀歷史文化學院통합검색과디렉터리你 好 ，的world就是不知道第三方库垃圾吧爱科技规范日本语/にほんご Nihongo`))))
}