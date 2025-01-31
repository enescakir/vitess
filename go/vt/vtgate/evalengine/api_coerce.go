/*
Copyright 2023 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package evalengine

import (
	"vitess.io/vitess/go/mysql/collations"
	"vitess.io/vitess/go/sqltypes"
)

func CoerceTo(value sqltypes.Value, typ sqltypes.Type) (sqltypes.Value, error) {
	cast, err := valueToEvalCast(value, value.Type(), collations.Unknown)
	if err != nil {
		return sqltypes.Value{}, err
	}
	return evalToSQLValueWithType(cast, typ), nil
}
