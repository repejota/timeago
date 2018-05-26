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
	"time"
)

// TimeAgo type represents the amount of time passed since an specific moment
// in time
type TimeAgo struct {
	time.Duration
}

// From sets the initial time of the TimeAgo type
func (t *TimeAgo) From(from *time.Time) *TimeAgo {
	t.Duration = time.Now().Sub(*from)
	return t
}

// String implements the Stringer interface for this type
func (t *TimeAgo) String() string {

	/*
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
	*/

	// unknown error
	return t.Duration.String()
}
