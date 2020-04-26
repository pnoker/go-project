/*
 * Copyright Pnoker. All Rights Reserved.
 */

package mongo

import (
	"emulator/internal/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (mc MongoClient) getNodeInfo(q bson.M) (models.NodeInfo, error) {
	s := mc.session.Copy()
	defer s.Close()

	var res models.NodeInfo
	err := s.DB(mc.database.Name).C(NodeInfo).Find(q).One(&res)
	if err != nil {
		return models.NodeInfo{}, errorMap(err)
	}

	return res, nil
}

func (mc MongoClient) getNodeInfos(q bson.M) ([]models.NodeInfo, error) {
	s := mc.session.Copy()
	defer s.Close()

	var res []models.NodeInfo
	err := s.DB(mc.database.Name).C(NodeInfo).Find(q).All(&res)
	if err != nil {
		return []models.NodeInfo{}, errorMap(err)
	}

	return res, nil
}

func (mc MongoClient) GetNodeInfoByKey(key string) (models.NodeInfo, error) {
	return mc.getNodeInfo(bson.M{"key": key})
}

func (mc MongoClient) GetAllNodeInfos() ([]models.NodeInfo, error) {
	return mc.getNodeInfos(nil)
}

func (mc MongoClient) AddNodeInfo(nodeInfo models.NodeInfo) error {
	s := mc.session.Copy()
	defer s.Close()

	indexNodeName := mgo.Index{
		Key:    []string{"key"},
		Unique: true,
	}
	indexNodeExpireTime := mgo.Index{
		Key:         []string{"expire_time"},
		Background:  true,
		ExpireAfter: time.Second * time.Duration(13),
	}

	if err := s.DB(mc.database.Name).C(NodeInfo).EnsureIndex(indexNodeName); err != nil {
		return err
	}
	if err := s.DB(mc.database.Name).C(NodeInfo).EnsureIndex(indexNodeExpireTime); err != nil {
		return err
	}
	if err := s.DB(mc.database.Name).C(NodeInfo).Insert(nodeInfo); err != nil {
		return err
	}

	return nil
}

func (mc MongoClient) UpdateNodeInfo(nodeInfo models.NodeInfo) error {
	s := mc.session.Copy()
	defer s.Close()

	if err := s.DB(mc.database.Name).C(NodeInfo).Update(bson.M{"key": nodeInfo.Key}, nodeInfo); err != nil {
		return err
	}

	return nil
}
