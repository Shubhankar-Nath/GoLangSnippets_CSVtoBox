// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package main

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type csvRecord_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var CsvRecordBinding = csvRecord_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 5165652673268497278,
}

// CsvRecord_ contains type-based Property helpers to facilitate some common operations such as Queries.
var CsvRecord_ = struct {
	Id        *objectbox.PropertyUint64
	serial    *objectbox.PropertyInt
	firstname *objectbox.PropertyString
	lastname  *objectbox.PropertyString
	email     *objectbox.PropertyString
	phone     *objectbox.PropertyString
	message   *objectbox.PropertyString
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	serial: &objectbox.PropertyInt{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	firstname: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	lastname: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	email: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	phone: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &CsvRecordBinding.Entity,
		},
	},
	message: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &CsvRecordBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (csvRecord_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (csvRecord_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("CsvRecord", 1, 5165652673268497278)
	model.Property("Id", 6, 1, 6958219774282140533)
	model.PropertyFlags(1)
	model.Property("serial", 6, 2, 3277826739377797544)
	model.Property("firstname", 9, 3, 7952788190216806884)
	model.Property("lastname", 9, 4, 1992680531661709499)
	model.Property("email", 9, 5, 3739284796375327075)
	model.Property("phone", 9, 6, 1196024410668352971)
	model.Property("message", 9, 7, 8818305894100925341)
	model.EntityLastPropertyId(7, 8818305894100925341)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (csvRecord_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*CsvRecord).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (csvRecord_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*CsvRecord).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (csvRecord_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (csvRecord_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*CsvRecord)
	var offsetfirstname = fbutils.CreateStringOffset(fbb, obj.firstname)
	var offsetlastname = fbutils.CreateStringOffset(fbb, obj.lastname)
	var offsetemail = fbutils.CreateStringOffset(fbb, obj.email)
	var offsetphone = fbutils.CreateStringOffset(fbb, obj.phone)
	var offsetmessage = fbutils.CreateStringOffset(fbb, obj.message)

	// build the FlatBuffers object
	fbb.StartObject(7)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, int64(obj.serial))
	fbutils.SetUOffsetTSlot(fbb, 2, offsetfirstname)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetlastname)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetemail)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetphone)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetmessage)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (csvRecord_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'CsvRecord' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	return &CsvRecord{
		Id:        propId,
		serial:    fbutils.GetIntSlot(table, 6),
		firstname: fbutils.GetStringSlot(table, 8),
		lastname:  fbutils.GetStringSlot(table, 10),
		email:     fbutils.GetStringSlot(table, 12),
		phone:     fbutils.GetStringSlot(table, 14),
		message:   fbutils.GetStringSlot(table, 16),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (csvRecord_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*CsvRecord, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (csvRecord_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*CsvRecord), nil)
	}
	return append(slice.([]*CsvRecord), object.(*CsvRecord))
}

// Box provides CRUD access to CsvRecord objects
type CsvRecordBox struct {
	*objectbox.Box
}

// BoxForCsvRecord opens a box of CsvRecord objects
func BoxForCsvRecord(ob *objectbox.ObjectBox) *CsvRecordBox {
	return &CsvRecordBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the CsvRecord.Id property on the passed object will be assigned the new ID as well.
func (box *CsvRecordBox) Put(object *CsvRecord) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the CsvRecord.Id property on the passed object will be assigned the new ID as well.
func (box *CsvRecordBox) Insert(object *CsvRecord) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *CsvRecordBox) Update(object *CsvRecord) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *CsvRecordBox) PutAsync(object *CsvRecord) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the CsvRecord.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the CsvRecord.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *CsvRecordBox) PutMany(objects []*CsvRecord) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *CsvRecordBox) Get(id uint64) (*CsvRecord, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*CsvRecord), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *CsvRecordBox) GetMany(ids ...uint64) ([]*CsvRecord, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*CsvRecord), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *CsvRecordBox) GetManyExisting(ids ...uint64) ([]*CsvRecord, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*CsvRecord), nil
}

// GetAll reads all stored objects
func (box *CsvRecordBox) GetAll() ([]*CsvRecord, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*CsvRecord), nil
}

// Remove deletes a single object
func (box *CsvRecordBox) Remove(object *CsvRecord) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *CsvRecordBox) RemoveMany(objects ...*CsvRecord) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the CsvRecord_ struct to create conditions.
// Keep the *CsvRecordQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *CsvRecordBox) Query(conditions ...objectbox.Condition) *CsvRecordQuery {
	return &CsvRecordQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the CsvRecord_ struct to create conditions.
// Keep the *CsvRecordQuery if you intend to execute the query multiple times.
func (box *CsvRecordBox) QueryOrError(conditions ...objectbox.Condition) (*CsvRecordQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &CsvRecordQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See CsvRecordAsyncBox for more information.
func (box *CsvRecordBox) Async() *CsvRecordAsyncBox {
	return &CsvRecordAsyncBox{AsyncBox: box.Box.Async()}
}

// CsvRecordAsyncBox provides asynchronous operations on CsvRecord objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type CsvRecordAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForCsvRecord creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use CsvRecordBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForCsvRecord(ob *objectbox.ObjectBox, timeoutMs uint64) *CsvRecordAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &CsvRecordAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *CsvRecordAsyncBox) Put(object *CsvRecord) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *CsvRecordAsyncBox) Insert(object *CsvRecord) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *CsvRecordAsyncBox) Update(object *CsvRecord) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *CsvRecordAsyncBox) Remove(object *CsvRecord) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all CsvRecord which Id is either 42 or 47:
// 		box.Query(CsvRecord_.Id.In(42, 47)).Find()
type CsvRecordQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *CsvRecordQuery) Find() ([]*CsvRecord, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*CsvRecord), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *CsvRecordQuery) Offset(offset uint64) *CsvRecordQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *CsvRecordQuery) Limit(limit uint64) *CsvRecordQuery {
	query.Query.Limit(limit)
	return query
}
