package main

import (
	"GO_LLD/OOPS/abstraction"
	"GO_LLD/OOPS/encapsulation"
	"GO_LLD/OOPS/inheritance"
	"GO_LLD/OOPS/polymorphism"
	"GO_LLD/SOLID/dependency_inversion"
	"GO_LLD/SOLID/interface_segregation"
	"GO_LLD/SOLID/liskov_substitution"
	"GO_LLD/SOLID/open_closed"
	"GO_LLD/SOLID/single_responsibility"
	"GO_LLD/design_patterns/behavioral/observer"
	"GO_LLD/design_patterns/behavioral/strategy"
	"GO_LLD/design_patterns/creational/builder"
	"GO_LLD/design_patterns/creational/factory"
	"GO_LLD/design_patterns/creational/prototype"
	"GO_LLD/design_patterns/creational/singleton"
	"GO_LLD/design_patterns/structural/adapter"
	"GO_LLD/design_patterns/structural/decorator"
)

func main() {
	// OOPS
	// Encapsulation
	encapsulation.BadEncapsulation()
	encapsulation.GoodEncapsulation()

	// Abstraction
	abstraction.BadSendGmail()
	abstraction.GoodSendGmail()

	// Inheritance
	// inheritance.BadInheritance()
	inheritance.GoodInheritance()

	// Polymorphism
	polymorphism.BadPolymorphism()
	polymorphism.GoodPolymorphism()

	// Solid Principles
	// single_responsibility
	single_responsibility.GoodSingleResponsibility()
	single_responsibility.BadSingleResponsibility()

	// open_closed
	open_closed.BadOpenClosed()
	open_closed.GoodOpenClosed()

	// liskov_substitution
	// liskov_substitution.BadLiskov()
	liskov_substitution.GoodViskov()

	// Interface_segregation
	// interface_segregation.BadInterfaceSegregation()
	interface_segregation.GoodInterfaceSegregation()

	// Dependency Inversion
	dependency_inversion.BadDependencyInversion()
	dependency_inversion.GoodDependencyInversion()

	// Design patterns
	// Creational
	// Singleton
	singleton.BadSingleton()
	singleton.GoodSingleton()

	// Factory
	factory.GoodFactoryPattern()
	factory.BadFactoryPattern()

	// Builder
	builder.BadBuilder()
	builder.GoodBuilder()

	// Prototype
	prototype.BadPrototype()
	prototype.GoodPrototype()

	//Behavoural
	// Observer
	observer.BadObserver()
	observer.GoodObserver()

	// Strategy
	strategy.BadStrategy()
	strategy.GoodStrategy()

	//Structural
	// Adapter
	adapter.BadAdapter()
	adapter.GoodAdapter()

	// Decorators
	decorator.BadDecorator()
	decorator.GoodDecorator()
}
