// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "LoadBalancer":
			resolverName, err := entityResolverNameForLoadBalancer(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancer": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancer": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerAnnotation":
			resolverName, err := entityResolverNameForLoadBalancerAnnotation(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerAnnotation": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerAnnotationByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerAnnotationByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerAnnotationByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerAnnotation": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerOrigin":
			resolverName, err := entityResolverNameForLoadBalancerOrigin(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerOrigin": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerOriginByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerOriginByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerOriginByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerOrigin": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerPool":
			resolverName, err := entityResolverNameForLoadBalancerPool(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerPool": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerPoolByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerPoolByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerPoolByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerPool": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerPort":
			resolverName, err := entityResolverNameForLoadBalancerPort(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerPort": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerPortByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerPortByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerPortByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerPort": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerProvider":
			resolverName, err := entityResolverNameForLoadBalancerProvider(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerProvider": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerProviderByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerProviderByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerProviderByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerProvider": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "LoadBalancerStatus":
			resolverName, err := entityResolverNameForLoadBalancerStatus(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "LoadBalancerStatus": %w`, err)
			}
			switch resolverName {

			case "findLoadBalancerStatusByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLoadBalancerStatusByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLoadBalancerStatusByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "LoadBalancerStatus": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Location":
			resolverName, err := entityResolverNameForLocation(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Location": %w`, err)
			}
			switch resolverName {

			case "findLocationByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLocationByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLocationByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Location": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Tenant":
			resolverName, err := entityResolverNameForTenant(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Tenant": %w`, err)
			}
			switch resolverName {

			case "findTenantByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findTenantByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindTenantByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Tenant": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForLoadBalancer(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancer", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerAnnotation(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerAnnotationByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerAnnotation", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerOrigin(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerOriginByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerOrigin", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerPool(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerPoolByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerPool", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerPort(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerPortByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerPort", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerProvider(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerProviderByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerProvider", ErrTypeNotFound)
}

func entityResolverNameForLoadBalancerStatus(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLoadBalancerStatusByID", nil
	}
	return "", fmt.Errorf("%w for LoadBalancerStatus", ErrTypeNotFound)
}

func entityResolverNameForLocation(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findLocationByID", nil
	}
	return "", fmt.Errorf("%w for Location", ErrTypeNotFound)
}

func entityResolverNameForTenant(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findTenantByID", nil
	}
	return "", fmt.Errorf("%w for Tenant", ErrTypeNotFound)
}