// Copyright (C) 2018 The Nori Authors info@nori.io
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program; if not, see <http://www.gnu.org/licenses/>.
package main

import (
	"context"

	"github.com/nori-io/common/v4/pkg/domain/config"
	em "github.com/nori-io/common/v4/pkg/domain/enum/meta"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	m "github.com/nori-io/common/v4/pkg/meta"
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/dummy/basic/pkg/dummy"
	http2 "github.com/nori-plugins/dummy/installable/internal/handler/http"
	"github.com/nori-plugins/dummy/installable/internal/handler/http/test"

	p "github.com/nori-io/common/v4/pkg/domain/plugin"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	Inst http.Http
}

var Plugin plugin

func (p plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	return nil
}

func (p *plugin) Instance() interface{} {
	return p.Inst
}

func (p plugin) Meta() meta.Meta {
	return m.Meta{
		ID: m.ID{
			ID:      "nori/http/dummy",
			Version: "1.0.0",
		},
		Author: m.Author{
			Name: "Nori",
			URL:  "https://nori.io",
		},
		Dependencies: []meta.Dependency{
			http.HttpInterface,
		},
		Description: nil,
		Interface:   dummy.DummyInterface,
		License:     nil,
		Links:       nil,
		Repository: m.Repository{
			Type: em.Git,
			URL:  "github.com/nori-plugins/dummy",
		},
		Tags: []string{"dummy", "rest", "api"},
	}
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {
	if p.Inst == nil {
		var err error
		p.Inst, err = http.GetHttp(registry)
		if err != nil {
			return err
		}

		http2.New(http2.Params{
			R:           p.Inst,
			TestHandler: &test.TestHandler{},
		})
	}

	return nil
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Install(ctx context.Context, registry registry.Registry) error {
	return nil
}
func (p plugin) UnInstall(ctx context.Context, registry registry.Registry) error {
	return nil
}
