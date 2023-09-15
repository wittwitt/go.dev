package string1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// 拼接字符串效率最高
// string.Builder

func Test_feild(t *testing.T) {

	s1 := "a b c d e"
	s2 := `a b 
	
	c 
	
	
	d e`

	fs := []string{"a", "b", "c", "d", "e"}

	require.Equal(t, fs, strings.Fields(s1))
	require.Equal(t, fs, strings.Split(s1, " "))
	//
	require.NotEqual(t, fs, strings.Split(s2, " "))
	require.Equal(t, fs, strings.Fields(s2)) // asciiSpace
}

func Test_ff(t *testing.T) {
	ss := `
	[[bafy2bzacea52llmmgiqta75tu4y3ltkkbxhhipzglxhegzvb43lntvici64bw bafy2bzacecifjztt7slxw7urx7ri6kneitykr5nfaab7ikiu7yojxni3lafhm bafy2bzacecg3jb27a7hmgeflqgni3wootpeuaxqckufn5blywdmhamepdqr5q bafy2bzaceaac6fefz3trtktmgkwyamxfuqbfx3elbxlrk6fgawvidk7qkvabw] [bafy2bzacebcswd7oet3u65nlwf2q6x37dlcymtzgkkki22elsudcafu4jq2e2 bafy2bzaceap37ikl4iqxcsmymo7dljmzbmqv4uignr3a7scvywscjoe4zyloo bafy2bzaceb3byrcw2u5w6myvbeyje2bmcwtt432hdsgskwln3fat4yo4bxpyk bafy2bzaced3uhm3kqrrdmmik2yoqdb4cdlodxpclbjtilokgw7m4u6ckr66bw] [bafy2bzacebsens7tnxyoh4u74isjiyltmieqnwxigdxxnltaqq3vr2a5ol3ow bafy2bzacedwweqvcqya3ud5l6sb5pikhrjqgrbvbrjpx6z2lxdhoc2bfb2bh4 bafy2bzacebvsfktyny254aiyjrdgevszzdlx6avydlanuktp7muot7a5dgg5y] [bafy2bzaceb3xx5xsx376c77j4zmspsi54lcccno3wbi5tcpctuejfvvurvzfy] [bafy2bzacedjo26z574uig4d3w5tkwusxdnl5g6kdixq5axaa3jklnkuormc44 bafy2bzacedwulzfwpp2bmbogv3bibfqkijpbelugwvpysqecddxpuweyl7myu] [bafy2bzacea5lolxjdz4zev2lfevth2an5tniau2gfqzhq3yd5cdh2c5ve3bvm bafy2bzaced2p6mawq6xkswjurgkckniryfrgbjzjxlwpiawqojst2kedqkts2] [bafy2bzaceccposmqjdiopllxfrrmrkyxz5sxu7ykpcbmrbgcwjjfrywkxijwc bafy2bzacecukfqbdedlgh3ujrwvucxziz3b63fl4zy3dzsc5m5zrol7d6fype bafy2bzacearzkqutv2mena3rukl2f2fgz5fl2mnvxxhpucl4odqhusdahx4v2] [bafy2bzacedzgbkdv4voxfkbk7v657sbdrinndy4yhi66zrp4lidxmrjpoxta2 bafy2bzacecw7p4y2haopk23r2v7lzikrhpg72moc5m565b77e42rgqx2zo4bs]]
	
	
	`

	s2 := strings.Fields(ss)

	ss2 := ""

	for _, item := range s2 {

		sPos := 0
		for {
			if !strings.Contains(item[sPos:], "[") {
				break
			}
			sPos = sPos + 1
		}

		s1 := item[:sPos] + `"` + item[sPos:]

		ePos := strings.Index(s1, "]")
		if ePos != -1 {
			s1 = s1[:ePos] + `"` + s1[ePos:] + ","
		} else {
			s1 = s1 + `",`
		}

		ss2 = ss2 + s1
	}

	ss2 = strings.ReplaceAll(ss2, "[", "{")
	ss2 = strings.ReplaceAll(ss2, "]", "}")

	t.Log(ss2)
}
