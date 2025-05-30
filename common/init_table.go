/*
Copyright 2021 The DnsJia Authors.
WebSite:  https://github.com/dnsjia/luban
Email:    OpenSource@dnsjia.com

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

package common

import (
	"github.com/dnsjia/luban/models"
	"github.com/dnsjia/luban/models/cmdb"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func MysqlTables(db *gorm.DB) {
	/*
		注册数据库表专用
	*/
	err := db.AutoMigrate(
		models.User{},
		models.Menu{},
		models.Role{},
		models.Dept{},
		models.K8SCluster{},
		models.ClusterVersion{},
		cmdb.CloudPlatform{},
		cmdb.VirtualMachine{},
		cmdb.TreeMenu{},
		cmdb.SSHRecord{},
		cmdb.SSHGlobalConfig{},
		//

	)
	if err != nil {
		LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	LOG.Info("register table success")
}
