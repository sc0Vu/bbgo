package ftx

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type orderRequest struct {
	*restRequest
}

/*
{
  "market": "XRP-PERP",
  "side": "sell",
  "price": 0.306525,
  "type": "limit",
  "size": 31431.0,
  "reduceOnly": false,
  "ioc": false,
  "postOnly": false,
  "clientId": null
}
*/
type PlaceOrderPayload struct {
	Market     string
	Side       string
	Price      float64
	Type       string
	Size       float64
	ReduceOnly bool
	IOC        bool
	PostOnly   bool
	ClientID   string
}

func (r *orderRequest) PlaceOrder(ctx context.Context, p PlaceOrderPayload) (orderResponse, error) {
	resp, err := r.
		Method("POST").
		ReferenceURL("api/orders").
		Payloads(map[string]interface{}{
			"market":     p.Market,
			"side":       p.Side,
			"price":      p.Price,
			"type":       p.Type,
			"size":       p.Size,
			"reduceOnly": p.ReduceOnly,
			"ioc":        p.IOC,
			"postOnly":   p.PostOnly,
			"clientId":   p.ClientID,
		}).
		DoAuthenticatedRequest(ctx)

	if err != nil {
		return orderResponse{}, err
	}
	var o orderResponse
	if err := json.Unmarshal(resp.Body, &o); err != nil {
		return orderResponse{}, fmt.Errorf("failed to unmarshal order response body to json: %w", err)
	}

	return o, nil
}
func (r *orderRequest) CancelOrderByOrderID(ctx context.Context, orderID uint64) (cancelOrderResponse, error) {
	resp, err := r.
		Method("DELETE").
		ReferenceURL("api/orders").
		ID(strconv.FormatUint(orderID, 10)).
		DoAuthenticatedRequest(ctx)
	if err != nil {
		return cancelOrderResponse{}, err
	}

	var co cancelOrderResponse
	if err := json.Unmarshal(resp.Body, &r); err != nil {
		return cancelOrderResponse{}, err
	}
	return co, nil
}

func (r *orderRequest) CancelOrderByClientID(ctx context.Context, clientID string) (cancelOrderResponse, error) {
	resp, err := r.
		Method("DELETE").
		ReferenceURL("api/orders/by_client_id").
		ID(clientID).
		DoAuthenticatedRequest(ctx)
	if err != nil {
		return cancelOrderResponse{}, err
	}

	var co cancelOrderResponse
	if err := json.Unmarshal(resp.Body, &r); err != nil {
		return cancelOrderResponse{}, err
	}
	return co, nil
}

func (r *orderRequest) OpenOrders(ctx context.Context, market string) (ordersResponse, error) {
	resp, err := r.
		Method("GET").
		ReferenceURL("api/orders").
		Query(map[string]string{"market": market}).
		DoAuthenticatedRequest(ctx)

	if err != nil {
		return ordersResponse{}, err
	}

	var o ordersResponse
	if err := json.Unmarshal(resp.Body, &o); err != nil {
		return ordersResponse{}, fmt.Errorf("failed to unmarshal open orders response body to json: %w", err)
	}

	return o, nil
}

func (r *orderRequest) OrdersHistory(ctx context.Context, market string, start, end time.Time, limit int64) (ordersHistoryResponse, error) {
	q := make(map[string]string)

	if limit > 0 {
		q["limit"] = strconv.FormatInt(limit, 10)
	}
	if len(market) > 0 {
		q["market"] = market
	}
	if start != (time.Time{}) {
		q["start_time"] = strconv.FormatInt(start.Unix(), 10)
	}
	if end != (time.Time{}) {
		q["end_time"] = strconv.FormatInt(end.Unix(), 10)
	}

	resp, err := r.
		Method("GET").
		ReferenceURL("api/orders/history").
		Query(q).
		DoAuthenticatedRequest(ctx)

	if err != nil {
		return ordersHistoryResponse{}, err
	}

	var o ordersHistoryResponse
	if err := json.Unmarshal(resp.Body, &o); err != nil {
		return ordersHistoryResponse{}, fmt.Errorf("failed to unmarshal orders history response body to json: %w", err)
	}

	return o, nil
}
