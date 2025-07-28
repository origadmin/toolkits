package types

// Params 定义了算法参数的通用接口
type Params interface {
	// Validate 验证参数自身的有效性。
	// 例如，Argon2 参数会验证 TimeCost, MemoryCost 等是否在合理范围内。
	Validate(*Config) error

	// ToMap 将参数转换为 map[string]string 格式，用于编码。
	ToMap() map[string]string

	// FromMap 从 map[string]string 格式中解析参数。
	FromMap(params map[string]string) error

	// String 将参数转换为字符串，用于打印。
	String() string
}
