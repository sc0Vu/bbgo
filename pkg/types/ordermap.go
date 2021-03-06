package types

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// OrderMap is used for storing orders by their order id
type OrderMap map[uint64]Order

func (m OrderMap) Backup() (orderForms []SubmitOrder) {
	for _, order := range m {
		orderForms = append(orderForms, order.Backup())
	}

	return orderForms
}

func (m OrderMap) Add(o Order) {
	m[o.OrderID] = o
}

// Update only updates the order when the order exists in the map
func (m OrderMap) Update(o Order) {
	if _, ok := m[o.OrderID]; ok {
		m[o.OrderID] = o
	}
}

func (m OrderMap) Remove(orderID uint64) {
	delete(m, orderID)
}

func (m OrderMap) IDs() (ids []uint64) {
	for id := range m {
		ids = append(ids, id)
	}

	return ids
}

func (m OrderMap) Exists(orderID uint64) bool {
	_, ok := m[orderID]
	return ok
}

func (m OrderMap) FindByStatus(status OrderStatus) (orders OrderSlice) {
	for _, o := range m {
		if o.Status == status {
			orders = append(orders, o)
		}
	}

	return orders
}

func (m OrderMap) Filled() OrderSlice {
	return m.FindByStatus(OrderStatusFilled)
}

func (m OrderMap) Canceled() OrderSlice {
	return m.FindByStatus(OrderStatusCanceled)
}

func (m OrderMap) Orders() (orders OrderSlice) {
	for _, o := range m {
		orders = append(orders, o)
	}
	return orders
}

type SyncOrderMap struct {
	orders         OrderMap
	pendingRemoval map[uint64]time.Time

	sync.RWMutex
}

func NewSyncOrderMap() *SyncOrderMap {
	return &SyncOrderMap{
		orders:         make(OrderMap),
		pendingRemoval: make(map[uint64]time.Time, 10),
	}
}

func (m *SyncOrderMap) Backup() []SubmitOrder {
	m.Lock()
	defer m.Unlock()
	return m.orders.Backup()
}

func (m *SyncOrderMap) Remove(orderID uint64) (exists bool) {

	exists = m.Exists(orderID)
	if exists {
		m.Lock()
		m.orders.Remove(orderID)
		m.Unlock()
	} else {
		m.pendingRemoval[orderID] = time.Now()
	}

	return exists
}

func (m *SyncOrderMap) Add(o Order) {
	m.Lock()

	match := false
	if len(m.pendingRemoval) > 0 {
		expireTime := time.Now().Add(-10 * time.Second)
		newPendingRemoval := make(map[uint64]time.Time, 10)
		for orderID, creationTime := range m.pendingRemoval {
			if o.OrderID == orderID {
				log.Warnf("found pending removal orderID = %d, removing order %+v from the store", orderID, o)
				match = true
				continue
			}

			if creationTime.Before(expireTime) {
				continue
			}

			newPendingRemoval[orderID] = creationTime
		}
		m.pendingRemoval = newPendingRemoval
	}

	if !match {
		m.orders.Add(o)
	}

	m.Unlock()
}

func (m *SyncOrderMap) Update(o Order) {
	m.Lock()
	m.orders.Update(o)
	m.Unlock()
}

func (m *SyncOrderMap) Iterate(it func(id uint64, order Order) bool) {
	m.Lock()
	for id := range m.orders {
		if it(id, m.orders[id]) {
			break
		}
	}
	m.Unlock()
}

func (m *SyncOrderMap) Exists(orderID uint64) (exists bool) {
	m.Lock()
	exists = m.orders.Exists(orderID)
	m.Unlock()
	return exists
}

func (m *SyncOrderMap) Len() int {
	m.Lock()
	defer m.Unlock()
	return len(m.orders)
}

func (m *SyncOrderMap) IDs() (ids []uint64) {
	m.Lock()
	ids = m.orders.IDs()
	m.Unlock()
	return ids
}

func (m *SyncOrderMap) FindByStatus(status OrderStatus) OrderSlice {
	m.Lock()
	defer m.Unlock()

	return m.orders.FindByStatus(status)
}

func (m *SyncOrderMap) Filled() OrderSlice {
	return m.FindByStatus(OrderStatusFilled)
}

// AnyFilled find any order is filled and stop iterating the order map
func (m *SyncOrderMap) AnyFilled() (order Order, ok bool) {
	m.Lock()
	defer m.Unlock()

	for _, o := range m.orders {
		if o.Status == OrderStatusFilled {
			ok = true
			order = o
			return order, ok
		}
	}

	return
}

func (m *SyncOrderMap) Canceled() OrderSlice {
	return m.FindByStatus(OrderStatusCanceled)
}

func (m *SyncOrderMap) Orders() (slice OrderSlice) {
	m.RLock()
	slice = m.orders.Orders()
	m.RUnlock()
	return slice
}

type OrderSlice []Order

func (s OrderSlice) IDs() (ids []uint64) {
	for _, o := range s {
		ids = append(ids, o.OrderID)
	}
	return ids
}
