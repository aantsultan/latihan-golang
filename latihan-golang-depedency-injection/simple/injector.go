//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgresql,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

//	func InitializedHelloService() *HelloService {
//		wire.Build(NewHelloService, NewSayHelloImpl)
//		return nil
//	}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBar() *FooBar {
	//wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Bar", "Foo"))
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*")) // bintang (*) -> all field
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
