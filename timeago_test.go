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
	"time"

	"github.com/repejota/timeago"
)

func TestDummy(t *testing.T) {

}

func TestInstance(t *testing.T) {
	var _ timeago.TimeAgo
}

func TestString(t *testing.T) {
	var tago timeago.TimeAgo
	expected := "right now"
	if tago.String() != expected {
		t.Errorf("Expected %q but got %q", tago.String(), expected)
	}
}

func TestSeconds(t *testing.T) {
	var tago timeago.TimeAgo
	expected := 0.0
	if tago.Seconds() != expected {
		t.Errorf("Expected %f but got %f", tago.Seconds(), expected)
	}
}

func TestMinutes(t *testing.T) {
	var tago timeago.TimeAgo
	expected := 0.0
	if tago.Minutes() != expected {
		t.Errorf("Expected %f but got %f", tago.Seconds(), expected)
	}
}

func TestHours(t *testing.T) {
	var tago timeago.TimeAgo
	expected := 0.0
	if tago.Hours() != expected {
		t.Errorf("Expected %f but got %f", tago.Seconds(), expected)
	}
}

func TestFewSecondsAgo(t *testing.T) {
	// 4 seconds ago
	ago := time.Now().Add(-time.Second * 4)

	var tago timeago.TimeAgo
	tago.Since(ago)
	expected := "a few seconds ago"
	if tago.String() != expected {
		t.Errorf("Expected %q but got %q", tago.String(), expected)
	}
}
