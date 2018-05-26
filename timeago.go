// Copyright 2018 Raül Pérez, repejota@gmail.com. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package timeago

import (
	"errors"
	"strconv"
	"time"
)

// TimeAgo ...
type TimeAgo struct {
	time.Duration
}

// NewTimeAgo ...
func NewTimeAgo(from time.Time) *TimeAgo {
	var timeago TimeAgo
	timeago.Duration = time.Now().Sub(from)
	return &timeago
}

// Render ...
func (t *TimeAgo) Render() (string, error) {

	// right now
	if t.Duration.Seconds() < 1 {
		return "right now", nil
	}

	// a few seconds ago
	if t.Duration.Seconds() < 15 {
		return "a few seconds ago", nil
	}

	// x seconds ago
	if t.Duration.Seconds() < 60 {
		return strconv.Itoa(int(t.Duration.Seconds())) + " seconds ago", nil
	}

	// unknown error
	err := errors.New("Unknown error")

	return err.Error(), err
}
