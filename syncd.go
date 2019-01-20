// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package syncd

import (
	"github.com/gin-gonic/gin"
)

type Syncd struct {
	Gin		*gin.Engine
}

func NewApp() *Syncd {
	return &Syncd{
		Gin: gin.Default(),
	}
}

func (s *Syncd) Start() error {
	return s.Gin.Run(":8868")
}
