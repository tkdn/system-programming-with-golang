package opaque

type internalBasicType int

// 非公開のメソッドなのでパッケージ以外でinterfaceが実装されることを抑止できる
// APIや他所からオーバーライドさせたくない場合に最適である
func (m *internalBasicType) implementsOpaque() {}

func (m *internalBasicType) GetNumber() int {
	return int(*m)
}

func NewBasic(number int) Opaque {
	// int へのポインタ(メモリアドレス
	o := internalBasicType(number)
	// ポインタへの参照を返す
	return &o
}
