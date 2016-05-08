// Exercise: Slices
//
// Pic 関数を実装してみましょう。 このプログラムを実行すると、生成した画像が下に表示されるはずです。 この関数は、長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを割り当てたものを返すように実装する必要があります。 画像は、整数値をグレースケール(実際はブルースケール)として解釈したものです。
//
// 生成する画像は、好きに選んでください。例えば、面白い関数に、 (x+y)/2 、 x*y 、o x^y などがあります。
//
// ヒント:( [][]uint8 に、各 []uint8 を割り当てるためにループを使用する必要があります)
//
// ヒント:( uint8(intValue) を型の変換のために使います)

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8

	for i := 0; i < dy; i++ {
		s = append(s, make([]uint8, dx))
		for j := 0; j < dx; j++ {
			s[i][j] = uint8(i * j)
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}
