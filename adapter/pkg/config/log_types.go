/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package config

type pkg struct {
	Name     string
	LogLevel string
}

type accessLog struct {
	Enable  bool
	LogFile string
	LogType string
	// ReservedLogFormat is reserved for Choreo Gateway Access Logs Observability feature. Changes to this may
	// break the functionality in the observability feature.
	ReservedLogFormat string
	// SecondaryLogFormat can be used by dev to log properties for debug purposes
	SecondaryLogFormat string
	JSONFormat         map[string]string
}

// LogConfig represents the configurations related to adapter logs and envoy access logs.
type LogConfig struct {
	Logfile  string
	LogLevel string
	// log rotation parameters.
	Rotation struct {
		IP         string
		Port       string
		MaxSize    int // megabytes
		MaxBackups int
		MaxAge     int //days
		Compress   bool
	}

	Pkg        []pkg
	AccessLogs *accessLog
}

func getDefaultLogConfig() *LogConfig {
	adapterLogConfig = &LogConfig{
		Logfile:  "/dev/null",
		LogLevel: "INFO",
		AccessLogs: &accessLog{
			Enable:  false,
			LogFile: "/dev/stdout",
			LogType: "text",
			// Following default value of "ReservedLogFormat" is document in log_config.toml for references.
			// Update log_config.toml if any changes are done here.
			ReservedLogFormat: "[%START_TIME%]' '%DYNAMIC_METADATA(envoy.filters.http.ext_authz:originalHost)%' " +
				"'%REQ(:AUTHORITY)%' '%REQ(:METHOD)%' '%DYNAMIC_METADATA(envoy.filters.http.ext_authz:originalPath)%' " +
				"'%REQ(:PATH)%' '%PROTOCOL%' '%RESPONSE_CODE%' '%RESPONSE_CODE_DETAILS%' '%RESPONSE_FLAGS%' '%REQ(USER-AGENT)%' " +
				"'%REQ(X-REQUEST-ID)%' '%REQ(X-FORWARDED-FOR)%' '%UPSTREAM_HOST%' '%BYTES_RECEIVED%' '%BYTES_SENT%' '%DURATION%' " +
				"'%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%' '%REQUEST_TX_DURATION%' '%RESPONSE_TX_DURATION%' '%REQUEST_DURATION%' " +
				"'%RESPONSE_DURATION%' '%ROUTE_NAME%' '",
			SecondaryLogFormat: "",
			JSONFormat: map[string]string{
				"time":          "%START_TIME%",
				"gwHost":        "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:originalHost)%",
				"host":          "%REQ(:AUTHORITY)%",
				"method":        "%REQ(:METHOD)%",
				"apiPath":       "%DYNAMIC_METADATA(envoy.filters.http.ext_authz:originalPath)%",
				"upstrmPath":    "%REQ(:PATH)%",
				"rName":         "%ROUTE_NAME%",
				"prot":          "%PROTOCOL%",
				"respCode":      "%RESPONSE_CODE%",
				"respCodeDtls":  "%RESPONSE_CODE_DETAILS%",
				"respFlag":      "%RESPONSE_FLAGS%",
				"ua":            "%REQ(USER-AGENT)%",
				"reqId":         "%REQ(X-REQUEST-ID)%",
				"xff":           "%REQ(X-FORWARDED-FOR)%",
				"upstrmHost":    "%UPSTREAM_HOST%",
				"bytesRecv":     "%BYTES_RECEIVED%",
				"bytesSent":     "%BYTES_SENT%",
				"dur":           "%DURATION%",
				"upstrmSvcTime": "%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%",
				"reqTxDur":      "%REQUEST_TX_DURATION%",
				"respTxDur":     "%RESPONSE_TX_DURATION%",
				"reqDur":        "%REQUEST_DURATION%",
				"respDur":       "%RESPONSE_DURATION%",
			},
		},
	}
	adapterLogConfig.Rotation.MaxSize = 10
	adapterLogConfig.Rotation.MaxAge = 2
	adapterLogConfig.Rotation.MaxBackups = 3
	adapterLogConfig.Rotation.Compress = true
	return adapterLogConfig
}
