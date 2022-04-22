package config

type DelegatingResolver struct {
	resolvers []Resolver
	Resolver
}

func (delegatingResolver DelegatingResolver) Resolve(object interface{}, path string) (interface{}, error) {
	resolvedObject := object
	var err error
	for i := 0; i < len(delegatingResolver.resolvers); i++ {
		resolvedObject, err = delegatingResolver.resolvers[i].Resolve(resolvedObject, path)
		if err != nil {
			return nil, err
		}
	}
	return resolvedObject, nil
}
