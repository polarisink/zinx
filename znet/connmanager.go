package znet

import (
	"errors"
	"sync"
	"zinx/ziface"
)

// @author lqs
// @date 2023/4/16 4:33 PM

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (m *ConnManager) Add(conn ziface.IConnection) {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	m.connections[conn.GetConnId()] = conn
}

func (m *ConnManager) Remove(connection ziface.IConnection) {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	delete(m.connections, connection.GetConnId())
}

func (m *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	m.connLock.RLock()
	defer m.connLock.RUnlock()
	if connection, ok := m.connections[connId]; ok {
		return connection, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (m *ConnManager) Len() int {
	return len(m.connections)
}

func (m *ConnManager) ClearConn() {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	for connId, conn := range m.connections {
		//停止
		conn.Stop()
		//删除
		delete(m.connections, connId)
	}
}
