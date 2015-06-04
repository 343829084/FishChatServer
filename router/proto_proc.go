//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"github.com/oikomi/FishChatServer/log"
	"github.com/oikomi/FishChatServer/libnet"
	"github.com/oikomi/FishChatServer/protocol"
	"github.com/oikomi/FishChatServer/common"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ProtoProc struct {
	Router   *Router
}

func NewProtoProc(r *Router) *ProtoProc {
	return &ProtoProc {
		Router : r,
	}
}

func (self *ProtoProc)procSendMsgP2P(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procSendMsgP2P")
	var err error
	send2ID := cmd.GetArgs()[0]
	send2Msg := cmd.GetArgs()[1]
	log.Info(send2Msg)
	self.Router.readMutex.Lock()
	defer self.Router.readMutex.Unlock()
	store_session, err := common.GetSessionFromCID(self.Router.sessionStore, send2ID)
	if err != nil {
		log.Warningf("no ID : %s", send2ID)
		
		return err
	}
	log.Info(store_session.MsgServerAddr)
	
	cmd.ChangeCmdName(protocol.ROUTE_MESSAGE_P2P_CMD)
	
	err = self.Router.msgServerClientMap[store_session.MsgServerAddr].Send(libnet.Json(cmd))
	if err != nil {
		log.Error("error:", err)
		return err
	}
	
	return nil
}

func (self *ProtoProc)procCreateTopic(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procCreateTopic")
	topicName := cmd.GetArgs()[0]
	serverAddr := cmd.GetAnyData().(string)
	self.Router.topicServerMap[topicName] = serverAddr
	
	return nil
}

//Note: router do not process topic
func (self *ProtoProc)procJoinTopic(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procJoinTopic")
	
	return nil
}

//Note: router do not process topic
func (self *ProtoProc)procSendMsgTopic(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procSendMsgTopic")
	//var err error
	//topicName := string(cmd.Args[0])
	//send2Msg := string(cmd.Args[1])

	
	return nil
}


