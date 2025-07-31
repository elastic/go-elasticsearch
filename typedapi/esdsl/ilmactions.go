// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ilmActions struct {
	v *types.IlmActions
}

func NewIlmActions() *_ilmActions {

	return &_ilmActions{v: types.NewIlmActions()}

}

func (s *_ilmActions) Allocate(allocate types.AllocateActionVariant) *_ilmActions {

	s.v.Allocate = allocate.AllocateActionCaster()

	return s
}

func (s *_ilmActions) Delete(delete types.DeleteActionVariant) *_ilmActions {

	s.v.Delete = delete.DeleteActionCaster()

	return s
}

func (s *_ilmActions) Downsample(downsample types.DownsampleActionVariant) *_ilmActions {

	s.v.Downsample = downsample.DownsampleActionCaster()

	return s
}

func (s *_ilmActions) Forcemerge(forcemerge types.ForceMergeActionVariant) *_ilmActions {

	s.v.Forcemerge = forcemerge.ForceMergeActionCaster()

	return s
}

func (s *_ilmActions) Freeze(freeze types.EmptyObjectVariant) *_ilmActions {

	s.v.Freeze = freeze.EmptyObjectCaster()

	return s
}

func (s *_ilmActions) Migrate(migrate types.MigrateActionVariant) *_ilmActions {

	s.v.Migrate = migrate.MigrateActionCaster()

	return s
}

func (s *_ilmActions) Readonly(readonly types.EmptyObjectVariant) *_ilmActions {

	s.v.Readonly = readonly.EmptyObjectCaster()

	return s
}

func (s *_ilmActions) Rollover(rollover types.RolloverActionVariant) *_ilmActions {

	s.v.Rollover = rollover.RolloverActionCaster()

	return s
}

func (s *_ilmActions) SearchableSnapshot(searchablesnapshot types.SearchableSnapshotActionVariant) *_ilmActions {

	s.v.SearchableSnapshot = searchablesnapshot.SearchableSnapshotActionCaster()

	return s
}

func (s *_ilmActions) SetPriority(setpriority types.SetPriorityActionVariant) *_ilmActions {

	s.v.SetPriority = setpriority.SetPriorityActionCaster()

	return s
}

func (s *_ilmActions) Shrink(shrink types.ShrinkActionVariant) *_ilmActions {

	s.v.Shrink = shrink.ShrinkActionCaster()

	return s
}

func (s *_ilmActions) Unfollow(unfollow types.EmptyObjectVariant) *_ilmActions {

	s.v.Unfollow = unfollow.EmptyObjectCaster()

	return s
}

func (s *_ilmActions) WaitForSnapshot(waitforsnapshot types.WaitForSnapshotActionVariant) *_ilmActions {

	s.v.WaitForSnapshot = waitforsnapshot.WaitForSnapshotActionCaster()

	return s
}

func (s *_ilmActions) IlmActionsCaster() *types.IlmActions {
	return s.v
}
