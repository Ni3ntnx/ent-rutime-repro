package rule

import (
    "fmt"
	"context"
    "errors"

	"entgo.io/ent/entql"
	"entrepro/ent/privacy"
)

func FilterTenantRule() privacy.QueryMutationRule {
	type TenantsFilter interface {
		WhereTenantID(entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		tf, ok := f.(TenantsFilter)
		if !ok {
			//return privacy.Denyf("unexpected filter type %T", f)
			return errors.New(fmt.Sprintf("unexpected filter type %T", f))
		}

		//tf.WhereTenantID(entql.StringEQ(GetTenantID(ctx)))
		tf.WhereTenantID(entql.StringEQ("asdasdas"))

		return privacy.Skip
	})
}