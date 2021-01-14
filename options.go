/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package log

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
)

const (
	flagLevel            = "log.level"
	flagFormat           = "log.format"
	flagEnableColor      = "log.enable-color"
	flagEnableCaller     = "log.enable-caller"
	flagOutputPaths      = "log.output-paths"
	flagErrorOutputPaths = "log.error-output-paths"

	consoleFormat = "console"
	jsonFormat    = "json"
)

// Options contains configuration items related to log.
type Options struct {
	Level            string   `json:"level" mapstructure:"level"`
	Format           string   `json:"format" mapstructure:"format"`
	EnableColor      bool     `json:"enable-color" mapstructure:"enable-color"`
	EnableCaller     bool     `json:"enable-caller" mapstructure:"enable-caller"`
	OutputPaths      []string `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths []string `json:"error-output-paths" mapstructure:"error-output-paths"`
}

// NewOptions creates a Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		Level:            zapcore.InfoLevel.String(),
		Format:           consoleFormat,
		EnableColor:      false,
		EnableCaller:     false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Validate validate the options fields.
func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Format)
	if format != consoleFormat && format != jsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}

	return errs
}

// AddFlags adds flags for log to the specified FlagSet object.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Level, flagLevel, o.Level, "Minimum log output `LEVEL`.")
	fs.StringVar(&o.Format, flagFormat, o.Format, "Log output `FORMAT`, support plain or json format.")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "Enable output ansi colors in plain format logs.")
	fs.BoolVar(&o.EnableCaller, flagEnableCaller, o.EnableCaller, "Enable output of caller information in the log.")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "Output paths of log.")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths, "Error output paths of log.")
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
