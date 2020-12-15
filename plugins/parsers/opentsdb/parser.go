package opentsdb

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
)

type OpentsdbParser struct {
	Method      string
	MetricName  string
    Timestamp   string
	DefaultTags map[string]string
}
func (p *OpentsdbParser) setDefaultTags(tags map[string]string) {
	p.DefaultTags = tags
}


func NewOpentsdbParser() (*OpentsdbParser, error) {
	var err error
	return fmt.Errorf("TODO %s", err.Error())
}
