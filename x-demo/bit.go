package x_demo

import (
	"fmt"
	"github.com/imroc/biu"
)

func BitDemo() {
	// byte是uint8的匿称，在操作系统是用一个字节来保存，用于表示[0-255]的整数
	var a uint8 = 20                   //
	fmt.Println(biu.ToBinaryString(a)) // 00010100

	// 将a的第2位设置为1，使用或操作
	fmt.Println(biu.ToBinaryString(a | 1<<1)) // 00010110

	// 将a的第3位设置为0，使用与操作，保证用于执行与操作的出了第3位为0，其他位皆为1。简单来说就是要得到11111011
	// 将1向左移动2位，然后取反即可。
	fmt.Println(biu.ToBinaryString(a &^ (1 << 2))) // 00010000

	// 移位操作取某一位值，先将目标位左移到最开头，然后右移到最后一位，如取第3位的值，
	fmt.Println(biu.ToBinaryString((a << 5) >> 7)) // 00000001

	// 取反某一位，即异或操作。在Go中使用 ^ 来表示取异或，位值相同为0，位值不同为1
	// 第3位取反，可以使用 1 << 2 = 00000100，00000100 ^ 00010100 = 00010000
	fmt.Println(biu.ToBinaryString((1 << 2) ^ a)) // 00010000
}
