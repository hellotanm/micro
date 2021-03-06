package server

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"time"

	gostore "github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/internal/namespace"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/store"
	pb "github.com/micro/micro/v3/service/store/proto"
)

const (
	defaultDatabase = namespace.DefaultNamespace
	defaultTable    = namespace.DefaultNamespace
	internalTable   = "store"
)

type handler struct {
	// local stores cache
	sync.RWMutex
	stores map[string]bool
}

// List all the keys in a table
func (h *handler) List(ctx context.Context, req *pb.ListRequest, stream pb.Store_ListStream) error {
	// set defaults
	if req.Options == nil {
		req.Options = &pb.ListOptions{}
	}
	if len(req.Options.Database) == 0 {
		req.Options.Database = defaultDatabase
	}
	if len(req.Options.Table) == 0 {
		req.Options.Table = defaultTable
	}

	// authorize the request
	if err := namespace.Authorize(ctx, req.Options.Database); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.List", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.List", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.List", err.Error())
	}

	// setup the store
	if err := h.setupTable(req.Options.Database, req.Options.Table); err != nil {
		return errors.InternalServerError("store.Store.List", err.Error())
	}

	// setup the options
	opts := []gostore.ListOption{
		gostore.ListFrom(req.Options.Database, req.Options.Table),
	}
	if len(req.Options.Prefix) > 0 {
		opts = append(opts, gostore.ListPrefix(req.Options.Prefix))
	}
	if req.Options.Offset > 0 {
		opts = append(opts, gostore.ListOffset(uint(req.Options.Offset)))
	}
	if req.Options.Limit > 0 {
		opts = append(opts, gostore.ListLimit(uint(req.Options.Limit)))
	}

	// list from the store
	vals, err := store.List(opts...)
	if err != nil && err == gostore.ErrNotFound {
		return errors.NotFound("store.Store.List", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.List", err.Error())
	}

	// serialize the response
	// TODO: batch sync
	rsp := new(pb.ListResponse)
	for _, val := range vals {
		rsp.Keys = append(rsp.Keys, val)
	}

	err = stream.Send(rsp)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return errors.InternalServerError("store.Store.List", err.Error())
	}
	return nil
}

// Read records from the store
func (h *handler) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	// set defaults
	if req.Options == nil {
		req.Options = &pb.ReadOptions{}
	}
	if len(req.Options.Database) == 0 {
		req.Options.Database = defaultDatabase
	}
	if len(req.Options.Table) == 0 {
		req.Options.Table = defaultTable
	}

	// authorize the request
	if err := namespace.Authorize(ctx, req.Options.Database); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.Read", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.Read", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Read", err.Error())
	}

	// setup the store
	if err := h.setupTable(req.Options.Database, req.Options.Table); err != nil {
		return errors.InternalServerError("store.Store.Read", err.Error())
	}

	// setup the options
	opts := []gostore.ReadOption{
		gostore.ReadFrom(req.Options.Database, req.Options.Table),
	}
	if req.Options.Prefix {
		opts = append(opts, gostore.ReadPrefix())
	}
	if req.Options.Limit > 0 {
		opts = append(opts, gostore.ReadLimit(uint(req.Options.Limit)))
	}
	if req.Options.Offset > 0 {
		opts = append(opts, gostore.ReadOffset(uint(req.Options.Offset)))
	}

	// read from the database
	vals, err := store.Read(req.Key, opts...)
	if err != nil && err == gostore.ErrNotFound {
		return errors.NotFound("store.Store.Read", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Read", err.Error())
	}

	// serialize the result
	for _, val := range vals {
		metadata := make(map[string]*pb.Field)
		for k, v := range val.Metadata {
			metadata[k] = &pb.Field{
				Type:  reflect.TypeOf(v).String(),
				Value: fmt.Sprintf("%v", v),
			}
		}
		rsp.Records = append(rsp.Records, &pb.Record{
			Key:      val.Key,
			Value:    val.Value,
			Expiry:   int64(val.Expiry.Seconds()),
			Metadata: metadata,
		})
	}
	return nil
}

// Write to the store
func (h *handler) Write(ctx context.Context, req *pb.WriteRequest, rsp *pb.WriteResponse) error {
	// validate the request
	if req.Record == nil {
		return errors.BadRequest("store.Store.Write", "no record specified")
	}

	// set defaults
	if req.Options == nil {
		req.Options = &pb.WriteOptions{}
	}
	if len(req.Options.Database) == 0 {
		req.Options.Database = defaultDatabase
	}
	if len(req.Options.Table) == 0 {
		req.Options.Table = defaultTable
	}

	// authorize the request
	if err := namespace.Authorize(ctx, req.Options.Database); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.Write", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.Write", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Write", err.Error())
	}

	// setup the store
	if err := h.setupTable(req.Options.Database, req.Options.Table); err != nil {
		return errors.InternalServerError("store.Store.Write", err.Error())
	}

	// setup the options
	opts := []gostore.WriteOption{
		gostore.WriteTo(req.Options.Database, req.Options.Table),
	}

	// construct the record
	metadata := make(map[string]interface{})
	for k, v := range req.Record.Metadata {
		metadata[k] = v.Value
	}
	record := &gostore.Record{
		Key:      req.Record.Key,
		Value:    req.Record.Value,
		Expiry:   time.Duration(req.Record.Expiry) * time.Second,
		Metadata: metadata,
	}

	// write to the store
	err := store.Write(record, opts...)
	if err != nil && err == gostore.ErrNotFound {
		return errors.NotFound("store.Store.Write", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Write", err.Error())
	}

	return nil
}

func (h *handler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	// set defaults
	if req.Options == nil {
		req.Options = &pb.DeleteOptions{}
	}
	if len(req.Options.Database) == 0 {
		req.Options.Database = defaultDatabase
	}
	if len(req.Options.Table) == 0 {
		req.Options.Table = defaultTable
	}

	// authorize the request
	if err := namespace.Authorize(ctx, req.Options.Database); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.Delete", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.Delete", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Delete", err.Error())
	}

	// setup the store
	if err := h.setupTable(req.Options.Database, req.Options.Table); err != nil {
		return errors.InternalServerError("store.Store.Delete", err.Error())
	}

	// setup the options
	opts := []gostore.DeleteOption{
		gostore.DeleteFrom(req.Options.Database, req.Options.Table),
	}

	// delete from the store
	if err := store.Delete(req.Key, opts...); err == gostore.ErrNotFound {
		return errors.NotFound("store.Store.Delete", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Delete", err.Error())
	}

	return nil
}

// Databases lists all the databases
func (h *handler) Databases(ctx context.Context, req *pb.DatabasesRequest, rsp *pb.DatabasesResponse) error {
	// authorize the request
	if err := namespace.Authorize(ctx, defaultDatabase); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.Databases", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.Databases", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Databases", err.Error())
	}

	// read the databases from the store
	opts := []gostore.ReadOption{
		gostore.ReadPrefix(),
		gostore.ReadFrom(defaultDatabase, internalTable),
	}
	recs, err := store.Read("databases/", opts...)
	if err != nil {
		return errors.InternalServerError("store.Store.Databases", err.Error())
	}

	// serialize the response
	rsp.Databases = make([]string, len(recs))
	for i, r := range recs {
		rsp.Databases[i] = strings.TrimPrefix(r.Key, "databases/")
	}
	return nil
}

// Tables returns all the tables in a database
func (h *handler) Tables(ctx context.Context, req *pb.TablesRequest, rsp *pb.TablesResponse) error {
	// set defaults
	if len(req.Database) == 0 {
		req.Database = defaultDatabase
	}

	// authorize the request
	if err := namespace.Authorize(ctx, req.Database); err == namespace.ErrForbidden {
		return errors.Forbidden("store.Store.Tables", err.Error())
	} else if err == namespace.ErrUnauthorized {
		return errors.Unauthorized("store.Store.Tables", err.Error())
	} else if err != nil {
		return errors.InternalServerError("store.Store.Tables", err.Error())
	}

	// construct the options
	opts := []gostore.ReadOption{
		gostore.ReadPrefix(),
		gostore.ReadFrom(defaultDatabase, internalTable),
	}

	// perform the query
	query := fmt.Sprintf("tables/%v/", req.Database)
	recs, err := store.Read(query, opts...)
	if err != nil {
		return errors.InternalServerError("store.Store.Tables", err.Error())
	}

	// serialize the response
	rsp.Tables = make([]string, len(recs))
	for i, r := range recs {
		rsp.Tables[i] = strings.TrimPrefix(r.Key, "tables/"+req.Database+"/")
	}
	return nil
}

func (h *handler) setupTable(database, table string) error {
	// lock (might be a race)
	h.Lock()
	defer h.Unlock()

	// attempt to get the database
	if _, ok := h.stores[database+":"+table]; ok {
		return nil
	}

	// record the new database in the internal store
	opt := gostore.WriteTo(defaultDatabase, internalTable)
	dbRecord := &gostore.Record{Key: "databases/" + database, Value: []byte{}}
	if err := store.Write(dbRecord, opt); err != nil {
		return fmt.Errorf("Error writing new database to internal table: %v", err)
	}

	// record the new table in the internal store
	tableRecord := &gostore.Record{Key: "tables/" + database + "/" + table, Value: []byte{}}
	if err := store.Write(tableRecord, opt); err != nil {
		return fmt.Errorf("Error writing new table to internal table: %v", err)
	}

	// write to the cache
	h.stores[database+":"+table] = true
	return nil
}
