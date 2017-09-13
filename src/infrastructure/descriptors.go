package infrastructure

import (
	"github.com/with-hindsight/chronicle/src/domain/object"
	"github.com/satori/go.uuid"
)

var EcommerceDomainId = uuid.FromStringOrNil("b9a185c3-4c80-4c0e-b44a-9e37423beb44")
var ShoppingContextId = uuid.FromStringOrNil("958bbea6-89be-4c56-a54e-34a27095941e")
var CartAggregateId = uuid.FromStringOrNil("f5e5d99a-220a-4946-a0b2-778dac099ff3")

var DescriptorMap = map[string]*object.DescriptorNode{
	"e-commerce": &object.DescriptorNode{
		Descriptor: nil,
		Children: map[string]*object.DescriptorNode{
			"shopping": &object.DescriptorNode{
				Descriptor: nil,
				Children: map[string]*object.DescriptorNode{
					"cart": &object.DescriptorNode{
						Descriptor: nil,
						Children: map[string]*object.DescriptorNode{
							"event": &object.DescriptorNode{
								Descriptor: nil,
							},
							"command": &object.DescriptorNode{
								Descriptor: nil,
								Children: map[string]*object.DescriptorNode{
									"create": &object.DescriptorNode{
										Descriptor: &object.Descriptor{
											TypeId: uuid.FromStringOrNil("08d14870-076e-44ed-a044-6947bdd448da"),
											AggregateTypeId: CartAggregateId,
											ContextTypeId: ShoppingContextId,
											DomainTypeId: EcommerceDomainId,
											Version: 1,
										},
										Children: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
