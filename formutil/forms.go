package formutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type FormHandler interface {
	Handle(c *gin.Context) (interface{}, error)
}

func Handle(form FormHandler, c *gin.Context) {
	if err := c.ShouldBind(form); err != nil {
		logrus.Error(errors.Wrap(err, "bind data to form failed"))
		ResponseErr(c, http.StatusBadRequest, err)
		return
	}

	data, err := form.Handle(c)
	if err == nil {
		ResponseOK(c, data)
		return
	}

	// If error has status code, return with e.Status
	// Else, return status 200 and error msg in body.
	if e, ok := err.(FormError); ok {
		ResponseErr(c, e.Status, err)
	} else {
		ResponseErr(c, http.StatusOK, err)
	}
}
