package utils
//
//import (
//	"container/list"
//	"errors"
//)
//
//type cacheNode struct {
//	Key,Value interface{}
//}
//
//func (cnode *cacheNode)NewCacheNode(k,v interface{})*cacheNode{
//	return &cacheNode{k,v}
//}
//
//type LRUCache struct {
//	Capacity int
//	dlist *list.List
//	cacheMap map[interface{}]*list.Element
//}
//
//func NewLRUCache(cap int)(*LRUCache){
//	return &LRUCache{
//		Capacity:cap,
//		dlist: list.New(),
//		cacheMap: make(map[interface{}]*list.Element)}
//}
//
//func (lru *LRUCache)Size()(int){
//	return lru.dlist.Len()
//}
//
//func (lru *LRUCache)Set(k,v interface{})(error){
//
//	if lru.dlist == nil {
//		return errors.New("LRUCache结构体未初始化.")
//	}
//
//	if pElement,ok := lru.cacheMap[k]; ok {
//		lru.dlist.MoveToFront(pElement)
//		pElement.Value.(*cacheNode).Value = v
//		return nil
//	}
//
//	newElement := lru.dlist.PushFront( &cacheNode{k,v} )
//	lru.cacheMap[k] = newElement
//
//	if lru.dlist.Len() > lru.Capacity {
//		//移掉最后一个
//		lastElement := lru.dlist.Back()
//		if lastElement == nil {
//			return nil
//		}
//		cacheNode := lastElement.Value.(*cacheNode)
//		delete(lru.cacheMap,cacheNode.Key)
//		lru.dlist.Remove(lastElement)
//	}
//	return nil
//}
//
//
//func (lru *LRUCache)Get(k interface{})(v interface{},ret bool,err error){
//
//	if lru.cacheMap == nil {
//		return v,false,errors.New("LRUCache结构体未初始化.")
//	}
//
//	if pElement,ok := lru.cacheMap[k]; ok {
//		lru.dlist.MoveToFront(pElement)
//		return pElement.Value.(*cacheNode).Value,true,nil
//	}
//	return v,false,nil
//}
//
//
//func (lru *LRUCache)Remove(k interface{})(bool){
//
//	if lru.cacheMap == nil {
//		return false
//	}
//
//	if pElement,ok := lru.cacheMap[k]; ok {
//		cacheNode := pElement.Value.(*cacheNode)
//		delete(lru.cacheMap,cacheNode.Key)
//		lru.dlist.Remove(pElement)
//		return true
//	}
//	return false
//}
//
///*
// 遍历LRUCache
// such as :
// var current *list.Element = nil
// for v,next,err := Fetch(current); err == nil; v,next,err = Fetch(current){
// 	now := time.Now().Unix()
//	user := v.(User)
//	if now > user.Expire {
//		UserMgr.cacheLRUlist.Remove(user.Id)
//	}
//
//	current = next
//
//	if current == nil {//遍历结束了
//		break
//	}
//
//}
// */
//
//
//func (lru *LRUCache)Fetch(current *list.Element)(v interface{}, next *list.Element, err error){
//
//	if lru.dlist == nil {
//		return v,nil,errors.New("not init.")
//	}
//
//	if current == nil {
//		pElement := lru.dlist.Front()
//		if pElement != nil {
//			return pElement.Value.(*cacheNode).Value, pElement.Next(), nil
//		}
//	}else {
//		return current.Value.(*cacheNode).Value, current.Next(), nil
//	}
//
//	return v, nil, errors.New("end.")
//}