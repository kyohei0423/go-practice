//+build wireinject

package animal

import "github.com/google/wire"

func InitializeDog() Dog {
	wire.Build(NewName, NewDog)
	return Dog{}
}
