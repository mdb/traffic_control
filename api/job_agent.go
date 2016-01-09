// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	"fmt"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	null "gopkg.in/guregu/null.v3"
	"time"
)

type JobAgent struct {
	Id          int64       `db:"id" json:"id"`
	Name        null.String `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	Active      int64       `db:"active" json:"active"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleJobAgent(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getJobAgent(id)
	} else if method == "POST" {
		return postJobAgent(payload)
	} else if method == "PUT" {
		return putJobAgent(id, payload)
	} else if method == "DELETE" {
		return delJobAgent(id)
	}
	return nil, nil
}

func getJobAgent(id int) (interface{}, error) {
	if id >= 0 {
		return getJobAgentById(id)
	} else {
		return getJobAgents()
	}
}

// @Title getJobAgentById
// @Description retrieves the job_agent information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobAgent
// @Resource /api/2.0
// @Router /api/2.0/job_agent/{id} [get]
func getJobAgentById(id int) (interface{}, error) {
	ret := []JobAgent{}
	arg := JobAgent{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from job_agent where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getJobAgents
// @Description retrieves the job_agent information for a certain id
// @Accept  application/json
// @Success 200 {array}    JobAgent
// @Resource /api/2.0
// @Router /api/2.0/job_agent [get]
func getJobAgents() (interface{}, error) {
	ret := []JobAgent{}
	queryStr := "select * from job_agent"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postJobAgent
// @Description enter a new job_agent
// @Accept  application/json
// @Param                 Body body     JobAgent   true "JobAgent object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_agent [post]
func postJobAgent(payload []byte) (interface{}, error) {
	var v JobAgent
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO job_agent("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",active"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:active"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putJobAgent
// @Description modify an existing job_agententry
// @Accept  application/json
// @Param                 Body body     JobAgent   true "JobAgent object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/job_agent [put]
func putJobAgent(id int, payload []byte) (interface{}, error) {
	var v JobAgent
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE job_agent SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",active = :active"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delJobAgentById
// @Description deletes job_agent information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    JobAgent
// @Resource /api/2.0
// @Router /api/2.0/job_agent/{id} [delete]
func delJobAgent(id int) (interface{}, error) {
	arg := JobAgent{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM job_agent WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
