package dummy

import "github.com/nori-io/common/v4/pkg/domain/meta"

const DummyInterface meta.Interface = "nori/Dummy"

type HttpDummy interface {
	Stop()
}
