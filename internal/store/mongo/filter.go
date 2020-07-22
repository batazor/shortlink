package mongo

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/batazor/shortlink/internal/store/query"
)

func getFilter(filter *query.Filter) bson.D {
	filterQuery := bson.D{}
	r := reflect.ValueOf(filter)

	for _, key := range filter.GetKeys() {
		val := reflect.Indirect(r).FieldByName(key).Interface().(*query.StringFilterInput)

		// Eq
		if val != nil && val.Eq != nil {
			filterQuery = append(filterQuery, primitive.E{Key: strings.ToLower(key), Value: *val.Eq})
		}

		// Ne
		if val != nil && val.Ne != nil {
			filterQuery = append(filterQuery, primitive.E{
				Key:   strings.ToLower(key),
				Value: bson.D{{"$ne", *val.Ne}},
			})
		}
	}

	return filterQuery
}
