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

import (
	"bytes"
	"github.com/metaleaf-io/assert"
	"net"
	"sync"
	"testing"
)

func TestSetLevel_WithWarning(t *testing.T) {
	// All subsequent tests should emit only warning and error messages.
	SetLevel(WARN)
}

func TestEmit_WithInfo_ShouldBeEmpty(t *testing.T) {
	var buf bytes.Buffer
	SetWriter(&buf)

	Info("test")
	assert.With(t).That(buf.String()).IsEmpty()
}

func TestEmit_WithWarn_ShouldNotBeEmpty(t *testing.T) {
	var buf bytes.Buffer
	SetWriter(&buf)

	Warn("test")
	assert.With(t).That(buf.String()).IsNotEmpty()
}

func TestSetServer_WithServer_ShouldReceive(t *testing.T) {
	// Create a simple UDP server that exits after one message is received.
	addr, err := net.ResolveUDPAddr("udp", "localhost:9999")
	assert.With(t).That(err).IsOk()

	sock, err := net.ListenUDP("udp", addr)
	assert.With(t).That(err).IsOk()

	defer sock.Close()

	var wait sync.WaitGroup

	go func() {
		buf := make([]byte, 4096)
		len, _, err := sock.ReadFromUDP(buf)
		assert.With(t).That(err).IsOk()
		assert.With(t).That(len).IsGreaterThan(0)
		assert.With(t).That(buf).IsNotNil()
		assert.With(t).That(string(buf)).IsNotEmpty()
		wait.Done()
	}()

	// Prevent output on the console.
	var buf bytes.Buffer
	SetWriter(&buf)

	// Send a message to the server.
	SetServer("localhost:9999")
	wait.Add(1)
	Error("test")
	wait.Wait()
}
