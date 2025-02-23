/*
Copyright 2023 The KodeRover Authors.

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

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/koderover/zadig/pkg/microservice/aslan/config"
)

type DBInstance struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty"         json:"id,omitempty"`
	Type      config.DBInstanceType `bson:"type"                  json:"type"`
	Name      string                `bson:"name"                  json:"name"`
	Projects  []string              `bson:"projects"              json:"projects"`
	Host      string                `bson:"host"                  json:"host"`
	Port      string                `bson:"port"                  json:"port"`
	Username  string                `bson:"username"              json:"username"`
	Password  string                `bson:"password"              json:"password,omitempty"`
	UpdateBy  string                `bson:"update_by"             json:"update_by"`
	CreatedAt int64                 `bson:"created_at"            json:"created_at"`
	UpdatedAt int64                 `bson:"updated_at"            json:"updated_at"`
}

func (h DBInstance) TableName() string {
	return "db_instance"
}
