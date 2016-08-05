package filters

import (
	"github.com/valyala/fasthttp"
	r "gopkg.in/dancannon/gorethink.v2"
)

var QUERY_PARAMS []string = []string{
	"status",
	"environment",
	"service",
}

func BuildAlertsFilter(queryArgs *fasthttp.Args) (rowFilter r.Term) {
	for i, queryParam := range QUERY_PARAMS {
		if !queryArgs.Has(queryParam){
			continue
		}

		paramFilter := buildQueryForParam(queryParam, queryArgs)

		if i == 0 {
			rowFilter = paramFilter
		}else {
			rowFilter = rowFilter.And(paramFilter)
		}
	}

	return rowFilter
}

func buildQueryForParam(queryParam string, queryArgs *fasthttp.Args)(r.Term){
	paramFilter := r.Row

	for i, queryValue := range getQueryValues(queryParam, queryArgs) {
		if queryParam == "service"{
			if i == 0 {
				paramFilter = paramFilter.Field(queryParam).Contains(queryValue)
			} else {
				paramFilter = paramFilter.Or(r.Row.Field(queryParam).Contains(queryValue))
			}
		}else {
			if i == 0 {
				paramFilter = paramFilter.Field(queryParam).Eq(queryValue)
			} else {
				paramFilter = paramFilter.Or(r.Row.Field(queryParam).Eq(queryValue))
			}
		}

	}

	return paramFilter
}

func getQueryValues(key string, queryArgs *fasthttp.Args) (values []string) {
	for _, value := range queryArgs.PeekMulti(key) {
		values = append(values, string(value))
	}
	return
}
