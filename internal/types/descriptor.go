package types

// Descriptor は定数やenum、objectなどの種類を表す列挙型です。
type Descriptor string

const (
	DescriptorConst  Descriptor = "const"
	DescriptorEnum   Descriptor = "enum"
	DescriptorObject Descriptor = "object"
)

// String は Descriptor の文字列表現を返します。
func (d Descriptor) String() string {
	return string(d)
}
