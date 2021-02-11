package lute_lite

import (
	"github.com/88250/lute/parse"
	"github.com/88250/lute/render"
	"github.com/HerbertHe/lute-lite/lexer"
)

type LuteLite struct {
	// 解析选项
	ParseOptions *parse.Options
	RenderOptions *render.Options
}

func New(opts ...ParseOption) (ret *LuteLite) {
	ret = &LuteLite{
		ParseOptions: parse.NewOptions(),
		RenderOptions: render.NewOptions(),
	}
	for _, opt := range opts {
		opt(ret)
	}

	return ret
}

// 类似 https://marked.js.org/using_pro#lexer 的输出
func (luteLite *LuteLite) lexer(markdown string) (ans string) {
	tree := parse.Parse("", []byte(markdown), luteLite.ParseOptions)
	lexerRenderer := lexer.NewLexerRenderer(tree, luteLite.RenderOptions)
	output := lexerRenderer.Render()
	ans = string(output)
	return
}

type ParseOption func(luteLite *LuteLite)
