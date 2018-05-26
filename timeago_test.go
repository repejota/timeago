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

package timeago_test

import (
	"testing"

	"github.com/repejota/timeago"
)

func TestDummy(t *testing.T) {

}

func TestInstance(t *testing.T) {
	var _ timeago.TimeAgo
}

/*
func TestNow(t *testing.T) {
	tm := time.Now()
	s, err := TimeAgo(tm)
	if err != nil {
		t.Error(err)
	}
	if s != "right now" {
		t.Error("Expected: 'right now' but got ", s)
	}
}

func TestFewSeconds(t *testing.T) {
	ago := time.Now().Add(-time.Second * 4)
	s, err := TimeAgo(ago)
	if err != nil {
		t.Error(err)
	}
	if s != "a few seconds ago" {
		t.Error("Expected: 'a few seconds ago' but got ", s)
	}
}

func TestSeconds(t *testing.T) {
	ago := time.Now().Add(-time.Second * 15)
	s, err := TimeAgo(ago)
	if err != nil {
		t.Error(err)
	}
	if s != "15 seconds ago" {
		t.Error("Expected: '15 seconds ago' but got ", s)
	}
}
*/
