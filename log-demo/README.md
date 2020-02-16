## lumberjack 

用于切割日志，按照日志的日期或者个数动态的生成文件和删除文件。

>golang实现的**lumberjack.v2**中，定义了日志文件名字，日志文件最大文件大小，日志文件保存天数，日志文件保存最大个数等。
>
>写入日志时，检测当前日志文件过大，就将当前文件关闭重命名（加上日期后缀），然后调用rotate操作建立一个原文件名的日志文件。
>
>rotate操作会去创建新日志文件，并整理旧的日志文件，根据文件日志排序，将天数超过预设值和文件个数超过预设值的文件清理掉。同时包含了日志打包压缩功能。

## 使用zap和zapcore

`zapcore`封装了核心逻辑，`zap`封装了调用逻辑。

``` go
// zap，初始化需要zapcode.Core
// Option 里面能够使用的方法不多，关键是 zap.Fields 方法 
// 针对Field，使用示例：zap.String("field01", "value01")
func New(core zapcore.Core, options ...Option) *Logger {}
func Fields(fs ...Field) Option {}


// zapcore，初始化需要Encoder，WriteSyncer，LevelEnabler
// 主要是通过 zapcore.NewJSONEncoder(encoderConfig) 来生成 
// 代码中是对于 jsonEncoder 对象的封装处理，该对象实现了接口 type Encoder interface { }
// WriteSyncer用于将buffer中的数据flush到指定的位置 
// LevelEnabler用于设置当前日志的级别
func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core {}

// 使用JSONEncoder结构体封装了一个sync.Pool，每次打印日志encode的时刻，都从池子中去去数据。
var _jsonPool = sync.Pool{New: func() interface{} {
	return &jsonEncoder{}
}}

func getJSONEncoder() *jsonEncoder {
	return _jsonPool.Get().(*jsonEncoder)
}

func putJSONEncoder(enc *jsonEncoder) {
	if enc.reflectBuf != nil {
		enc.reflectBuf.Free()
	}
	enc.EncoderConfig = nil
	enc.buf = nil
	enc.spaced = false
	enc.openNamespaces = 0
	enc.reflectBuf = nil
	enc.reflectEnc = nil
	_jsonPool.Put(enc)
}

type jsonEncoder struct {
	*EncoderConfig
	buf            *buffer.Buffer
	spaced         bool // include spaces after colons and commas
	openNamespaces int

	// for encoding generic values by reflection
	reflectBuf *buffer.Buffer
	reflectEnc *json.Encoder
}

// clone 组装数据 
func (enc *jsonEncoder) clone() *jsonEncoder {
	clone := getJSONEncoder()
	clone.EncoderConfig = enc.EncoderConfig
	clone.spaced = enc.spaced
	clone.openNamespaces = enc.openNamespaces
	clone.buf = bufferpool.Get()
	return clone
}

func (enc *jsonEncoder) EncodeEntry(ent Entry, fields []Field) (*buffer.Buffer, error) {
    // 该函数先后调用getJSONEncoder，putJSONEncoder
    return 
}
```
