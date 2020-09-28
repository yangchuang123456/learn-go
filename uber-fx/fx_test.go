package node

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"log"
	"testing"
	"time"
)

type TypeA string
type TypeB int

type TypeC struct {
	A TypeA
	b TypeB
}

func newTypeC(lc fx.Lifecycle, a TypeA, b TypeB) TypeC {
	fmt.Println("newTypeC")
	c := TypeC{
		A: a,
		b: b,
	}
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			c.Stop()
			return nil
		},
	})

	return c
}

type TypeD struct {
	C TypeC
	A *TypeA
}

func newTypeD(c TypeC, a *TypeA) *TypeD {
	fmt.Println("newTypeD")
	return &TypeD{
		A: a,
		C: c,
	}
}

type TypeE struct {
	fx.In
	C TypeC
	D *TypeD
}

func newTypeE(c TypeC, d *TypeD) *TypeE {
	fmt.Println("newTypeE")
	return &TypeE{
		C: c,
		D: d,
	}
}

func (c *TypeC) Start() {
	fmt.Println("typeC start")
}

func (c *TypeC) Stop() {
	fmt.Println("typeC stop")
}

func (c *TypeD) Start() {
	fmt.Println("typeD start")
	time.Sleep(3 * time.Second)
	fmt.Println("typeD start sleep end")
}

func (c *TypeD) Stop() {
	fmt.Println("typeD stop")
}

func (c *TypeE) Start() {
	fmt.Println("typeE start")
}

func (c *TypeE) Stop() {
	fmt.Println("typeE stop")
}

func RunHello() {
	//time.Sleep(3*time.Second)
	fmt.Println("run hello")
}

func Test_fx(t *testing.T) {
	ctx := context.Background()
	var typeE *TypeE
	typeA := TypeA("pointer")
	stop, err := New(ctx,
		Override(new(TypeB), TypeB(1)),

		Override(RunHelloKey, RunHello),
		Override(new(*TypeD), newTypeD),

		Override(new(TypeC), newTypeC),

		func(s *Settings) error {
			s.nodeType = FullNode
			return nil
		},
		ApplyIf(isType(FullNode),
			Override(new(TypeA), TypeA("yc-test")),
			Override(new(*TypeA), &typeA),
		),

		func(s *Settings) error {
			//s.nodeType = FullNode
			e := &TypeE{}
			s.invokes[InitTypeE] = fx.Populate(e)
			typeE = e
			return nil
		},
	)
	assert.NoError(t, err)

	log.Println(typeE)
	typeE.Start()

	defer stop(ctx)
}
