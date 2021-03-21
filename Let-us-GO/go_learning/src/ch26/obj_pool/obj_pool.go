package objectpool

import (
	"errors"
	"time"
)

// ReusableObj 要池化的對象
type ReusableObj struct {
}

// ObjPool pool
type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

// NewObjPool create pool
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

// GetObj get obj from pool
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}

}

// ReleaseObj release
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}