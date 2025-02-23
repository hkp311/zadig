/*
Copyright 2022 The KodeRover Authors.

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

package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/koderover/zadig/pkg/microservice/aslan/config"
	"github.com/koderover/zadig/pkg/microservice/user/core/repository/models"
	mongotool "github.com/koderover/zadig/pkg/tool/mongo"
)

type ProjectColl struct {
	*mongo.Collection

	coll string
}

func NewProjectColl() *ProjectColl {
	name := models.Project{}.TableName()
	return &ProjectColl{Collection: mongotool.Database(config.MongoDatabase()).Collection(name), coll: name}
}

func (c *ProjectColl) GetCollectionName() string {
	return c.coll
}

func (c *ProjectColl) List() ([]*models.Project, error) {
	var resp []*models.Project

	cursor, err := c.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	err = cursor.All(context.TODO(), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
