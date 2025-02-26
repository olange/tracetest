package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/trace"
)

const (
	TracetestMetadataFieldStartTime string = "tracetest.span.start_time"
	TracetestMetadataFieldEndTime   string = "tracetest.span.end_time"
	TracetestMetadataFieldDuration  string = "tracetest.span.duration"
	TracetestMetadataFieldType      string = "tracetest.span.type"
	TracetestMetadataFieldName      string = "tracetest.span.name"
	TracetestMetadataFieldParentID  string = "tracetest.span.parent_id"
	TracetestMetadataFieldKind      string = "tracetest.span.kind"
)

type Attributes map[string]string

func (a Attributes) Get(key string) string {
	if v, ok := a[key]; ok {
		return v
	}

	return ""
}

func (a Attributes) SetPointerValue(key string, value *string) {
	if value != nil {
		a[key] = *value
	}
}

type Spans []Span

func (s Spans) ForEach(fn func(ix int, _ Span) bool) Spans {
	for i, span := range s {
		doNext := fn(i, span)
		if !doNext {
			break
		}
	}
	return s
}

func (s Spans) OrEmpty(fn func()) Spans {
	if len(s) == 0 {
		fn()
	}
	return s
}

type SpanKind string

var (
	SpanKindClient       SpanKind = "client"
	SpanKindServer       SpanKind = "server"
	SpanKindConsumer     SpanKind = "consumer"
	SpanKindProducer     SpanKind = "producer"
	SpanKindInternal     SpanKind = "internal"
	SpanKindUnespecified SpanKind = "unespecified"
)

type Span struct {
	ID         trace.SpanID
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	Attributes Attributes
	Kind       SpanKind

	Parent   *Span   `json:"-"`
	Children []*Span `json:"-"`
}

type encodedSpan struct {
	ID         string
	Name       string
	StartTime  string
	EndTime    string
	Attributes Attributes
	Children   []encodedSpan
}

func (s Span) IsZero() bool {
	return !s.ID.IsValid()
}

func (s Span) MarshalJSON() ([]byte, error) {
	enc := encodeSpan(s)
	return json.Marshal(&enc)
}

func encodeSpan(s Span) encodedSpan {
	return encodedSpan{
		ID:         s.ID.String(),
		Name:       s.Name,
		StartTime:  fmt.Sprintf("%d", s.StartTime.UnixMilli()),
		EndTime:    fmt.Sprintf("%d", s.EndTime.UnixMilli()),
		Attributes: s.Attributes,
		Children:   encodeChildren(s.Children),
	}
}

func encodeChildren(children []*Span) []encodedSpan {
	res := make([]encodedSpan, len(children))
	for i, c := range children {
		res[i] = encodeSpan(*c)
	}
	return res
}

func (s *Span) UnmarshalJSON(data []byte) error {
	aux := encodedSpan{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	return s.decodeSpan(aux)
}

func (s *Span) decodeSpan(aux encodedSpan) error {
	sid, err := trace.SpanIDFromHex(aux.ID)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	children, err := decodeChildren(s, aux.Children)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	startTime, err := getTimeFromString(aux.StartTime)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	endTime, err := getTimeFromString(aux.EndTime)
	if err != nil {
		return fmt.Errorf("unmarshal span: %w", err)
	}

	s.ID = sid
	s.Name = aux.Name
	s.StartTime = startTime.UTC()
	s.EndTime = endTime.UTC()
	s.Attributes = aux.Attributes
	s.Children = children

	return nil
}

func getTimeFromString(value string) (time.Time, error) {
	milliseconds, err := strconv.Atoi(value)
	if err != nil {
		// Maybe it is in RFC3339 format. Convert it for compatibility sake
		output, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return time.Time{}, fmt.Errorf("could not convert string (%s) to time: %w", value, err)
		}

		return output, nil
	}

	return time.UnixMilli(int64(milliseconds)), nil
}

func decodeChildren(parent *Span, children []encodedSpan) ([]*Span, error) {
	if len(children) == 0 {
		return nil, nil
	}
	res := make([]*Span, len(children))
	for i, c := range children {
		span := &Span{
			Parent: parent,
		}
		if err := span.decodeSpan(c); err != nil {
			return nil, fmt.Errorf("unmarshal children: %w", err)
		}

		children, err := decodeChildren(span, c.Children)
		if err != nil {
			return nil, fmt.Errorf("unmarshal children: %w", err)
		}

		span.Children = children
		res[i] = span
	}
	return res, nil
}

func (span Span) setMetadataAttributes() Span {
	span.Attributes[TracetestMetadataFieldName] = span.Name
	span.Attributes[TracetestMetadataFieldType] = spanType(span.Attributes)
	span.Attributes[TracetestMetadataFieldDuration] = spanDuration(span)
	span.Attributes[TracetestMetadataFieldStartTime] = fmt.Sprintf("%d", span.StartTime.UnixNano())
	span.Attributes[TracetestMetadataFieldEndTime] = fmt.Sprintf("%d", span.EndTime.UnixNano())

	return span
}

func (span Span) setTriggerResultAttributes(result TriggerResult) Span {
	switch result.Type {
	case TriggerTypeHTTP:
		resp := result.HTTP
		jsonheaders, _ := json.Marshal(resp.Headers)
		span.Attributes["tracetest.response.status"] = fmt.Sprintf("%d", resp.StatusCode)
		span.Attributes["tracetest.response.body"] = resp.Body
		span.Attributes["tracetest.response.headers"] = string(jsonheaders)
	case TriggerTypeGRPC:
		resp := result.GRPC
		jsonheaders, _ := json.Marshal(resp.Metadata)
		span.Attributes["tracetest.response.status"] = fmt.Sprintf("%d", resp.StatusCode)
		span.Attributes["tracetest.response.body"] = resp.Body
		span.Attributes["tracetest.response.headers"] = string(jsonheaders)
	}

	return span
}

func AugmentRootSpan(span Span, result TriggerResult) Span {
	return span.
		setMetadataAttributes().
		setTriggerResultAttributes(result)
}

func NewTracetestRootSpan(run Run) Span {
	return AugmentRootSpan(Span{
		ID:         IDGen.SpanID(),
		Name:       TriggerSpanName,
		StartTime:  run.ServiceTriggeredAt,
		EndTime:    run.ServiceTriggerCompletedAt,
		Attributes: Attributes{},
		Children:   []*Span{},
	}, run.TriggerResult)
}
