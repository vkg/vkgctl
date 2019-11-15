package コマンド太郎

import (
	"github.com/spf13/cobra"
	"github.com/vkg/vkgctl/写真太郎"
)

const description = `元劇団員 ✿人生ボムなし残機なし気合い避け！ ✿ ダークソウル3 🏆コンプ ✿ 最近はとにかく人と話しているコミュ障 ✿ Backend じゃなくなった
http://vkgtaro.jp
`

// Describe太郎 is a subcommend like `vkgctl describe太郎`
var Describe太郎 = &cobra.Command{
	Use:        "describe太郎",
	SuggestFor: []string{"説明", "説明太郎"},
	Short:      "describe describes what vkgtaro is",
	Run: func(cmd *cobra.Command, args []string) {
		写真太郎.PrintIfPossible太郎()
		cmd.Printf(description)
	},
}
