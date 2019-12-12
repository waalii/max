// Copyright 2018 MaiCoin Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package max

import (
	"context"
	"net/http"
	"time"

	"github.com/waalii/max/api"
)

func NewClient(opts ...ClientOption) *client {
	c := &client{
		requestTimeout: 10 * time.Second,
		cfg:            api.NewConfiguration(),
		middlewares:    make([]middleware, 0),
		stopCh:         make(chan struct{}),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.c = api.NewAPIClient(c.config())

	go timeCalibrater(c, 1*time.Hour)

	return c
}

func (c *client) Close() {
	close(c.stopCh)
}

// Interface check
var _ PublicAPI = &publicClient{}
var _ PrivateAPI = &privateClient{}

type client struct {
	c              *api.APIClient
	requestTimeout time.Duration
	middlewares    []middleware
	stopCh         chan struct{}

	timeDiff       time.Duration
	cfg            *api.Configuration
	timeCalibrater *time.Ticker
}

type middleware func(http.RoundTripper) http.RoundTripper

func (c *client) config() *api.Configuration {
	s := http.DefaultTransport

	for _, m := range c.middlewares {
		s = m(s)
	}

	c.cfg.HTTPClient = http.DefaultClient
	c.cfg.HTTPClient.Transport = s
	c.cfg.HTTPClient.Timeout = c.requestTimeout

	return c.cfg
}

func timeCalibrater(c *client, period time.Duration) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t, err := c.Time(context.Background())
			if err != nil {
				continue
			}

			c.timeDiff = t.Sub(time.Now())
		case <-c.stopCh:
			return
		}
	}
}
