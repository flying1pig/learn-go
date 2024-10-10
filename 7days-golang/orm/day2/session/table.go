package session

import "learn-go/7days-golang/orm/day2/schema"

func (s *Session) Model(value interface{}) *Session {
	return nil
}

func (s *Session) RefTable() *schema.Schema {
	return nil
}

func (s *Session) CreateTable() error {
	return nil
}

func (s *Session) DropTable() error {
	return nil
}

func (s *Session) HasTable() bool {
	return false
}
