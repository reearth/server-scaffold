// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/reearth/scaffold/server/internal/transport/gql/gqlmodel"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Asset struct {
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Project   func(childComplexity int) int
		ProjectID func(childComplexity int) int
	}

	AssetConnection struct {
		Edges    func(childComplexity int) int
		PageInfo func(childComplexity int) int
	}

	AssetEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Mutation struct {
		CreateAsset func(childComplexity int, input gqlmodel.CreateAssetInput) int
		DeleteAsset func(childComplexity int, assetID gqlmodel.ID) int
		UpdateAsset func(childComplexity int, input gqlmodel.UpdateAssetInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Project struct {
		Assets      func(childComplexity int, filter gqlmodel.AssetFilter) int
		ID          func(childComplexity int) int
		Name        func(childComplexity int) int
		WorkspaceID func(childComplexity int) int
	}

	Query struct {
		Assets func(childComplexity int, filter gqlmodel.AssetFilter) int
		Node   func(childComplexity int, id gqlmodel.ID) int
		Nodes  func(childComplexity int, ids []gqlmodel.ID) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Asset.id":
		if e.complexity.Asset.ID == nil {
			break
		}

		return e.complexity.Asset.ID(childComplexity), true

	case "Asset.name":
		if e.complexity.Asset.Name == nil {
			break
		}

		return e.complexity.Asset.Name(childComplexity), true

	case "Asset.project":
		if e.complexity.Asset.Project == nil {
			break
		}

		return e.complexity.Asset.Project(childComplexity), true

	case "Asset.projectId":
		if e.complexity.Asset.ProjectID == nil {
			break
		}

		return e.complexity.Asset.ProjectID(childComplexity), true

	case "AssetConnection.edges":
		if e.complexity.AssetConnection.Edges == nil {
			break
		}

		return e.complexity.AssetConnection.Edges(childComplexity), true

	case "AssetConnection.pageInfo":
		if e.complexity.AssetConnection.PageInfo == nil {
			break
		}

		return e.complexity.AssetConnection.PageInfo(childComplexity), true

	case "AssetEdge.cursor":
		if e.complexity.AssetEdge.Cursor == nil {
			break
		}

		return e.complexity.AssetEdge.Cursor(childComplexity), true

	case "AssetEdge.node":
		if e.complexity.AssetEdge.Node == nil {
			break
		}

		return e.complexity.AssetEdge.Node(childComplexity), true

	case "Mutation.createAsset":
		if e.complexity.Mutation.CreateAsset == nil {
			break
		}

		args, err := ec.field_Mutation_createAsset_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateAsset(childComplexity, args["input"].(gqlmodel.CreateAssetInput)), true

	case "Mutation.deleteAsset":
		if e.complexity.Mutation.DeleteAsset == nil {
			break
		}

		args, err := ec.field_Mutation_deleteAsset_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteAsset(childComplexity, args["assetId"].(gqlmodel.ID)), true

	case "Mutation.updateAsset":
		if e.complexity.Mutation.UpdateAsset == nil {
			break
		}

		args, err := ec.field_Mutation_updateAsset_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateAsset(childComplexity, args["input"].(gqlmodel.UpdateAssetInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Project.assets":
		if e.complexity.Project.Assets == nil {
			break
		}

		args, err := ec.field_Project_assets_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Project.Assets(childComplexity, args["filter"].(gqlmodel.AssetFilter)), true

	case "Project.id":
		if e.complexity.Project.ID == nil {
			break
		}

		return e.complexity.Project.ID(childComplexity), true

	case "Project.name":
		if e.complexity.Project.Name == nil {
			break
		}

		return e.complexity.Project.Name(childComplexity), true

	case "Project.workspaceID":
		if e.complexity.Project.WorkspaceID == nil {
			break
		}

		return e.complexity.Project.WorkspaceID(childComplexity), true

	case "Query.assets":
		if e.complexity.Query.Assets == nil {
			break
		}

		args, err := ec.field_Query_assets_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Assets(childComplexity, args["filter"].(gqlmodel.AssetFilter)), true

	case "Query.Node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_Node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(gqlmodel.ID)), true

	case "Query.Nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_Nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]gqlmodel.ID)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAssetFilter,
		ec.unmarshalInputCreateAssetInput,
		ec.unmarshalInputUpdateAssetInput,
	)
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, opCtx.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, opCtx.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../schema/asset.graphql", Input: `type Asset implements Node {
  id: ID!
  projectId: ID!
  name: String!

  project: Project
}

type AssetConnection {
  pageInfo: PageInfo!
  edges: [AssetEdge!]
}

type AssetEdge {
  cursor: String!
  node: Asset!
}

input AssetFilter {
  projectId: ID
  first: Int
  last: Int
  after: Cursor
  before: Cursor
}

input CreateAssetInput {
  projectId: ID!
  name: String!
}

input UpdateAssetInput {
  assetId: ID!
  name: String
}

extend type Query {
  assets(filter: AssetFilter!): AssetConnection!
}

extend type Mutation {
  createAsset(input: CreateAssetInput!): Asset!
  updateAsset(input: UpdateAssetInput!): Asset!
  deleteAsset(assetId: ID!): ID!
}
`, BuiltIn: false},
	{Name: "../../../schema/project.graphql", Input: `type Project implements Node {
  id: ID!
  workspaceID: ID!
  name: String!

  assets(filter: AssetFilter!): AssetConnection!
  # workspace: Workspace
}
`, BuiltIn: false},
	{Name: "../../../schema/shared.graphql", Input: `scalar Cursor

interface Node {
  id: ID!
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: Cursor
  endCursor: Cursor
}

type Query {
  Node(id: ID!): Node
  Nodes(ids: [ID!]!): [Node]
}

type Mutation

schema {
  query: Query
  mutation: Mutation
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
