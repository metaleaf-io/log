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
	"fmt"
	"os"
	"path"
	"time"
)

const (
	ISO8601Micro = "2006-01-02T15:04:05.000000Z0700"
)

var (
	programName string
)

// init creates a default console logger that can be used immediately with no
// further configuration necessary.
func init() {
	programName = path.Base(os.Args[0])
}

// Debug emits a message with the DEBUG level.
func Debug(msg string) {
	emit(DEBUG, msg)
}

// Info emits a message with the INFO level.
func Info(msg string) {
	emit(INFO, msg)
}

// Warn emits a message with the WARN level.
func Warn(msg string) {
	emit(WARN, msg)
}

// Error emits a message with the ERROR level.
func Error(msg string) {
	emit(ERROR, msg)
}

// The actual emitter
func emit(level Level, message string) {
    fmt.Fprintf(os.Stderr, "%s [%s] %s %s\n",
        time.Now().UTC().Format(ISO8601Micro), programName, level.String(), message)
}