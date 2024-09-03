/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package global

import (
	"sync"

	"github.com/go-ldap/ldap/v3"
)

// 缓存相关

type StoreInterface interface {
}

type system struct {
	LdapConfig *ldap.Conn
}

type store struct {
	rmx       sync.RWMutex
	systemMap map[string]*system
}

var (
	instance *store
	once     sync.Once
)

func NewStore() StoreInterface {
	once.Do(func() {
		instance = &store{
			systemMap: make(map[string]*system),
		}
	})
	return instance
}

func (s *store) SetCache(name string, system *system) {
	s.rmx.Lock()
	defer s.rmx.Unlock()
	s.systemMap[name] = system
}

func (s *store) DelCache(name string) bool {
	s.rmx.Lock()
	defer s.rmx.Unlock()
	if _, ok := s.systemMap[name]; !ok {
		return false
	}
	delete(s.systemMap, name)
	return true
}

func (s *store) UpdateCache(name string, system *system) bool {
	s.rmx.Lock()
	defer s.rmx.Unlock()
	if _, ok := s.systemMap[name]; !ok {
		s.SetCache(name, system)
		return true
	}
	s.systemMap[name] = system
	return true
}

func (s *store) GetLdapConfigCache(name string) (*ldap.Conn, bool) {
	s.rmx.RLock()
	defer s.rmx.RUnlock()
	if config, ok := s.systemMap[name]; ok {
		return config.LdapConfig, true
	}
	return nil, false
}
