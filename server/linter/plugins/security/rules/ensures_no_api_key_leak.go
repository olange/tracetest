package rules

import (
	"context"
	"fmt"

	"github.com/kubeshop/tracetest/server/model"
)

type ensuresNoApiKeyLeakRule struct {
	model.BaseRule
}

var (
	httpHeadersFields  = []string{"authorization", "x-api-key"}
	httpResponseHeader = "http.response.header."
	httpRequestHeader  = "http.request.header."
)

func NewEnsuresNoApiKeyLeakRule() model.Rule {
	return &ensuresNoApiKeyLeakRule{
		BaseRule: model.BaseRule{
			Name:             "No API Key Leak",
			Description:      "Ensure no API keys are leaked in http headers",
			ErrorDescription: "The following attributes are exposing API keys:",
			Tips:             []string{},
			Weight:           80,
		},
	}
}

func (r ensuresNoApiKeyLeakRule) Evaluate(ctx context.Context, trace model.Trace) (model.RuleResult, error) {
	passed := true
	results := make([]model.Result, 0)
	for _, span := range trace.Flat {
		if span.Attributes.Get("tracetest.span.type") == "http" {
			result := r.validate(span)
			if !result.Passed {
				passed = false
			}
			results = append(results, result)
		}
	}

	return model.RuleResult{
		BaseRule: r.BaseRule,
		Passed:   passed,
		Results:  results,
	}, nil
}

func (r ensuresNoApiKeyLeakRule) validate(span *model.Span) model.Result {
	leakedFields := make([]model.Error, 0)
	for _, field := range httpHeadersFields {
		requestHeader := fmt.Sprintf("%s%s", httpRequestHeader, field)
		if span.Attributes.Get(requestHeader) != "" {
			leakedFields = append(leakedFields, model.Error{
				Value:       field,
				Description: fmt.Sprintf("Leaked request API Key found for attribute: %s. Value: %s", field, span.Attributes.Get(requestHeader)),
			})
		}

		responseHeader := fmt.Sprintf("%s%s", httpResponseHeader, field)
		if span.Attributes.Get(responseHeader) != "" {
			leakedFields = append(leakedFields, model.Error{
				Value:       field,
				Description: fmt.Sprintf("Leaked response API Key found for attribute: %s. Value: %s", field, span.Attributes.Get(responseHeader)),
			})
		}
	}

	return model.Result{
		Passed: len(leakedFields) == 0,
		SpanID: span.ID.String(),
		Errors: leakedFields,
	}
}
