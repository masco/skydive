/*
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package agents

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"github.com/redhat-cip/skydive/flow"
	"github.com/redhat-cip/skydive/storage"
)

const (
	MAX_DGRAM_SIZE = 1500
)

type SFlowAgent struct {
	Addr    string
	Port    int
	Storage storage.Connection
}

func (agent *SFlowAgent) Start() error {
	var buf [MAX_DGRAM_SIZE]byte

	addr := net.UDPAddr{
		Port: agent.Port,
		IP:   net.ParseIP(agent.Addr),
	}
	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		return err
	}

	for {
		_, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			continue
		}

		p := gopacket.NewPacket(buf[:], layers.LayerTypeSFlow, gopacket.Default)
		sflowLayer := p.Layer(layers.LayerTypeSFlow)
		sflowPacket, ok := sflowLayer.(*layers.SFlowDatagram)
		if !ok {
			continue
		}

		if sflowPacket.SampleCount > 0 {
			for _, sample := range sflowPacket.FlowSamples {
				flows := flow.FLowsFromSFlowSample(sample)
				if agent.Storage != nil {
					agent.Storage.StoreFlows(flows)
				}
				fmt.Println(flows)
			}
		}
	}

	return nil
}

func NewSFlowAgent(addr string, port int, storage storage.Connection) SFlowAgent {
	agent := SFlowAgent{Addr: addr, Port: port, Storage: storage}
	return agent
}
