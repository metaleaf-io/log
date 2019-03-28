/*
 * Copyright 2019 Metaleaf.io
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import "testing"

func TestLevel_String(t *testing.T) {
	var actual, expected string

	actual = DEBUG.String()
	expected = "DEBUG"
	if actual != expected {
		t.Error("TestLevel_String expected:", expected, "actual:", actual)
	}

	actual = INFO.String()
	expected = "INFO"
	if actual != expected {
		t.Error("TestLevel_String expected:", expected, "actual:", actual)
	}

	actual = WARN.String()
	expected = "WARN"
	if actual != expected {
		t.Error("TestLevel_String expected:", expected, "actual:", actual)
	}

	actual = ERROR.String()
	expected = "ERROR"
	if actual != expected {
		t.Error("TestLevel_String expected:", expected, "actual:", actual)
	}

	// Test with a bad level.
	defer func() {
		if r := recover(); r == nil {
			t.Error("TestLevel_String: did not panic")
		}
	}()

	Level(-1).String()
}

func TestParse(t *testing.T) {
	var actual, expected Level

	actual = Parse("DEBUG")
	expected = DEBUG
	if actual != expected {
		t.Error("TestParse expected:", expected, "actual:", actual)
	}

	actual = Parse("INFO")
	expected = INFO
	if actual != expected {
		t.Error("TestParse expected:", expected, "actual:", actual)
	}

	actual = Parse("WARN")
	expected = WARN
	if actual != expected {
		t.Error("TestParse expected:", expected, "actual:", actual)
	}

	actual = Parse("ERROR")
	expected = ERROR
	if actual != expected {
		t.Error("TestParse expected:", expected, "actual:", actual)
	}

	// Test with a bad string.
	defer func() {
		if r := recover(); r == nil {
			t.Error("TestParse: did not panic")
		}
	}()

	Parse("foobar")
}