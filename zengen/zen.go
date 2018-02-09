package zengen

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/lyft/protoc-gen-star"
)

// static assert
var _ pgs.Plugin = (*ZenGen)(nil)

// ZenGen is a protoc plugin
type ZenGen struct {
	*pgs.PluginBase
}

// New create a new zengen instance
func New(name string) *ZenGen {
	return &ZenGen{
		PluginBase: new(pgs.PluginBase),
	}
}

// Name satisfies the generator.Plugin interface.
func (z *ZenGen) Name() string { return "zengen" }

// Generate satisfies the generator.Plugin interface.
func (z *ZenGen) Generate(file *generator.FileDescriptor) {
	z.Push(file.GetName())
	defer z.Pop()

	z.AddImport("zen", "github.com/philchia/zen", file)
	z.AddImport("http", "net/http", file)

	for _, s := range file.GetService() {
		z.GenerateService(s)
	}

}

// GenerateService generate zen router code
func (z *ZenGen) GenerateService(service *descriptor.ServiceDescriptorProto) {
	z.Push(service.GetName())
	defer z.Pop()

	z.P("type ", service.GetName(), "Service", " interface {")
	for _, m := range service.GetMethod() {
		z.In()
		z.P(m.GetName(), "(zen.Context, *", getTypeName(m.GetInputType()), ") (*", getTypeName(m.GetOutputType()), ", error)")
	}
	z.Out()
	z.P("}")

	z.GenerateRegister(service)
}

func (z *ZenGen) GenerateRegister(service *descriptor.ServiceDescriptorProto) {
	z.P("func Register", service.GetName(), "Server(router zen.Router, server ", service.GetName(), "Service) {")
	z.In()

	for _, m := range service.Method {
		z.P(`router.Post("`, lowercase(service.GetName()), ".", lowercase(m.GetName()), `", func(ctx zen.Context) {`)
		z.In()
		z.P(`var req = new(`, getTypeName(m.GetInputType()), ")")
		z.P(`if err := ctx.BindJSON(req); err != nil {`)
		z.In()
		z.P("ctx.WriteStatus(http.StatusBadRequest)")
		z.P("return")
		z.Out()
		z.P("}")

		z.P("resp, err := server.", m.GetName(), "(ctx, req)")
		z.P("if err != nil {")
		z.In()
		z.P("ctx.WriteStatus(http.StatusBadRequest)")
		z.P("return")
		z.Out()
		z.P("}")

		z.P("ctx.JSON(resp)")
		z.P("})")
		z.P()
	}

	z.Out()
	z.P("}")
}

func getTypeName(name string) string {
	fields := strings.Split(name, ".")
	return fields[len(fields)-1]
}

func lowercase(s string) string {
	return strings.ToLower(s)
}
