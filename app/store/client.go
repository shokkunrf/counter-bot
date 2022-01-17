package store

type Counter struct {
	UserID string
	Count  int
}

type Client struct {
	counters []Counter
}

var client, clientErr = newClient()

func newClient() (*Client, error) {
	return &Client{
		counters: []Counter{},
	}, nil
}

func GetClient() (*Client, error) {
	if clientErr != nil {
		return &Client{}, clientErr
	}

	return client, nil
}

func (c *Client) GetCounters() ([]Counter, error) {
	return c.counters, nil
}

func (c *Client) IncrementCount(userID string) error {
	for i, counter := range c.counters {
		if counter.UserID == userID {
			c.counters[i].Count++
			return nil
		}
	}

	c.counters = append(c.counters, Counter{
		UserID: userID,
		Count:  1,
	})
	return nil
}

func (c *Client) DecrementCount(userID string) error {
	for i, counter := range c.counters {
		if counter.UserID == userID {
			c.counters[i].Count--
			return nil
		}
	}

	c.counters = append(c.counters, Counter{
		UserID: userID,
		Count:  -1,
	})
	return nil
}

func (c *Client) ResetCount(userID string) error {
	for i, counter := range c.counters {
		if counter.UserID == userID {
			c.counters[i].Count = 0
			return nil
		}
	}

	c.counters = append(c.counters, Counter{
		UserID: userID,
		Count:  0,
	})
	return nil
}
